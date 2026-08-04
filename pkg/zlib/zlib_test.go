//go:build zlibcompat

package zlib

// Byte-identical conformance tests against the C zlib library via the
// cgo helper in libz_cgo.go (both files gated by the `zlibcompat` build tag).
//
// Run with:
//   go test -tags=zlibcompat ./pkg/zlib/
//
// For each random input we run libz's deflate at the chosen level (default
// windowBits=15, memLevel=8, strategy=Z_DEFAULT_STRATEGY) and compare bytes
// to the pure-Go encoder. Any mismatch fails with a minimal reproducer so we
// can find the divergence.

import (
	"bytes"
	"fmt"
	"math/rand"
	"testing"
)

func libzCompress(t *testing.T, src []byte, level, windowBits, memLevel, strategy int) []byte {
	t.Helper()
	out, rc := czCompress(src, level, windowBits, memLevel, strategy)
	if rc != 0 {
		t.Fatalf("libz deflate failed: rc=%d", rc)
	}
	return out
}

func TestAdler32MatchesLibz(t *testing.T) {
	r := rand.New(rand.NewSource(1))
	for trial := 0; trial < 50; trial++ {
		n := 5 + r.Intn(2044)
		src := make([]byte, n)
		r.Read(src)
		want := libzCompress(t, src, 6, 15, 8, 0)
		if len(want) < 4 {
			t.Fatalf("trial %d: short libz output", trial)
		}
		got := adler32(1, src)
		expA := uint32(want[len(want)-4])<<24 | uint32(want[len(want)-3])<<16 |
			uint32(want[len(want)-2])<<8 | uint32(want[len(want)-1])
		if got != expA {
			t.Fatalf("trial %d: adler32 mismatch: got=%08x want=%08x", trial, got, expA)
		}
	}
}

// TestRandomMatchesLibz is the main conformance test: random strings of
// random length, every compression level 0..9, compared byte-for-byte to libz.
func TestRandomMatchesLibz(t *testing.T) {
	levels := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	if testing.Short() {
		levels = []int{0, 6, 9}
	}

	trials := 500
	if testing.Short() {
		trials = 50
	}

	for _, level := range levels {
		level := level
		t.Run(fmt.Sprintf("level=%d", level), func(t *testing.T) {
			r := rand.New(rand.NewSource(int64(level)*100 + 1))
			for i := 0; i < trials; i++ {
				n := 5 + r.Intn(2044) // 5..2048 inclusive
				src := make([]byte, n)
				r.Read(src)
				assertSameBytes(t, src, level, i)
			}
		})
	}
}

// TestAdversarialMatchesLibz exercises inputs that hit non-random regions of
// the deflate state machine: long runs (RLE friendly), repeating substrings
// (heavy hash hits), structured text, and all-zero. Each is tested at every
// level. The interesting paths are: longest_match early-exits, lazy match
// switching, dynamic-vs-static-vs-stored decisions, and the bit-length tree.
func TestAdversarialMatchesLibz(t *testing.T) {
	levels := []int{0, 1, 6, 9}
	if !testing.Short() {
		levels = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	}

	cases := map[string][]byte{
		"all-zero":     bytes.Repeat([]byte{0}, 2048),
		"all-ff":       bytes.Repeat([]byte{0xff}, 2048),
		"ababab":       bytes.Repeat([]byte("ab"), 1024),
		"abcdefgh-rep": bytes.Repeat([]byte("abcdefgh"), 256),
		"single-byte":  []byte{0x42},
		"short-5":      []byte("hello"),
		"runs-256":     runs(256, 8),
		"english-ish":  bytes.Repeat([]byte("the quick brown fox jumps over the lazy dog. "), 45),
		"sparse-runs":  sparseRuns(2000),
		"window-edge":  bytes.Repeat([]byte("X"), 2047),
	}

	for _, level := range levels {
		level := level
		t.Run(fmt.Sprintf("level=%d", level), func(t *testing.T) {
			for name, src := range cases {
				name, src := name, src
				t.Run(name, func(t *testing.T) {
					assertSameBytes(t, src, level, 0)
				})
			}
		})
	}
}

func assertSameBytes(t *testing.T, src []byte, level, trial int) {
	t.Helper()
	want := libzCompress(t, src, level, 15, 8, 0)
	got, err := Compress(src, level)
	if err != nil {
		t.Fatalf("trial %d level=%d: Compress err: %v", trial, level, err)
	}
	if !bytes.Equal(got, want) {
		div := firstDiff(got, want)
		// keep the dump small for big inputs
		dump := func(b []byte) string {
			if len(b) > 200 {
				return fmt.Sprintf("%x...(+%d more)", b[:200], len(b)-200)
			}
			return fmt.Sprintf("%x", b)
		}
		t.Fatalf("level=%d trial=%d: output differs at byte %d\n  src(%d)=%s\n  got (%d)=%s\n  want(%d)=%s",
			level, trial, div, len(src), dump(src), len(got), dump(got), len(want), dump(want))
	}
}

func runs(n, runLen int) []byte {
	out := make([]byte, 0, n)
	for i := 0; len(out) < n; i++ {
		b := byte(i & 0xff)
		for j := 0; j < runLen && len(out) < n; j++ {
			out = append(out, b)
		}
	}
	return out
}

func sparseRuns(n int) []byte {
	out := make([]byte, n)
	for i := range out {
		// Mostly zeros with occasional non-zero bytes.
		if i%37 == 0 {
			out[i] = byte(i)
		}
	}
	return out
}

// TestDecoderRoundtrip verifies that bytes compressed by libz can be decoded
// by our pure-Go decoder. We also test the decoder on our own encoder's
// output (i.e. self-roundtrip) as a sanity check.
func TestDecoderRoundtrip(t *testing.T) {
	levels := []int{0, 1, 6, 9}
	r := rand.New(rand.NewSource(123))
	for _, level := range levels {
		level := level
		t.Run(fmt.Sprintf("level=%d", level), func(t *testing.T) {
			for i := 0; i < 100; i++ {
				n := 5 + r.Intn(2044)
				src := make([]byte, n)
				r.Read(src)

				// 1) Decode libz's output.
				cz := libzCompress(t, src, level, 15, 8, 0)
				out, err := Decompress(cz)
				if err != nil {
					t.Fatalf("Decompress(libz output) failed: %v", err)
				}
				if !bytes.Equal(out, src) {
					t.Fatalf("decode of libz output != src (level=%d trial=%d len=%d)",
						level, i, n)
				}

				// 2) Self-roundtrip.
				gz, err := Compress(src, level)
				if err != nil {
					t.Fatalf("Compress err: %v", err)
				}
				out, err = Decompress(gz)
				if err != nil {
					t.Fatalf("Decompress(self) failed: %v", err)
				}
				if !bytes.Equal(out, src) {
					t.Fatalf("self-roundtrip != src (level=%d trial=%d len=%d)",
						level, i, n)
				}
			}
		})
	}
}

// TestLargeInputsMatchLibz exercises inputs larger than the 32K window so we
// hit window-slide / hash-table-slide code paths. Sizes are chosen to land
// just before, exactly at, and after the window boundary.
func TestLargeInputsMatchLibz(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping large-input test in -short")
	}
	levels := []int{1, 6, 9}
	sizes := []int{16 * 1024, 32*1024 - 1, 32 * 1024, 32*1024 + 1, 64 * 1024, 100 * 1024}

	r := rand.New(rand.NewSource(7))
	for _, size := range sizes {
		size := size
		src := make([]byte, size)
		r.Read(src)
		for _, level := range levels {
			level := level
			t.Run(fmt.Sprintf("size=%d/level=%d", size, level), func(t *testing.T) {
				assertSameBytes(t, src, level, 0)
			})
		}
	}
}

func firstDiff(a, b []byte) int {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	for i := 0; i < n; i++ {
		if a[i] != b[i] {
			return i
		}
	}
	return n
}
