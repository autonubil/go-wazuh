package zlib

// Pure-Go RFC 1950 (zlib) / RFC 1951 (deflate) decoder. Output is uniquely
// determined by the compressed stream, so byte-identical-to-libz isn't a
// meaningful goal here — we just have to decode correctly. The conformance
// tests verify by round-tripping libz output through Decompress and
// comparing to the original input.

import (
	"errors"
	"fmt"
	"io"
)

// Length / distance base values + extra-bit counts for the inflate side
// (RFC 1951 §3.2.5). The first entry is for code 257.

var lengthBase = [29]int{
	3, 4, 5, 6, 7, 8, 9, 10, 11, 13, 15, 17, 19, 23, 27, 31,
	35, 43, 51, 59, 67, 83, 99, 115, 131, 163, 195, 227, 258,
}

var lengthExtra = [29]int{
	0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 2, 2, 2, 2,
	3, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 0,
}

var distBase = [30]int{
	1, 2, 3, 4, 5, 7, 9, 13, 17, 25, 33, 49, 65, 97, 129, 193,
	257, 385, 513, 769, 1025, 1537, 2049, 3073, 4097, 6145, 8193,
	12289, 16385, 24577,
}

var distExtra = [30]int{
	0, 0, 0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6,
	7, 7, 8, 8, 9, 9, 10, 10, 11, 11, 12, 12, 13, 13,
}

// codeLengthOrder is the order code-length-code lengths are written in
// dynamic blocks (RFC 1951 §3.2.7). Same as blOrder in the encoder.
var codeLengthOrder = [19]int{
	16, 17, 18, 0, 8, 7, 9, 6, 10, 5, 11, 4, 12, 3, 13, 2, 14, 1, 15,
}

// huffTable is a canonical Huffman decoding table. We use a simple
// length-then-symbol lookup: for each bit-length we store the firstCode and
// the firstSymbol; decoding reads bits LSB-first within bytes and
// MSB-first within a code. Adequate for code lengths up to 15.
type huffTable struct {
	// counts[i] = number of codes of length i, for i in 1..15.
	counts [16]int
	// symbol[i] = the i'th symbol in canonical order.
	symbol []int
}

func buildHuffman(lens []int) (*huffTable, error) {
	t := &huffTable{}
	t.symbol = make([]int, len(lens))
	var offs [16]int
	for _, l := range lens {
		if l < 0 || l > 15 {
			return nil, errors.New("zlib: huffman: code length out of range")
		}
		t.counts[l]++
	}
	if t.counts[0] == len(lens) {
		// All-zero code lengths — no codes. Valid in some corner cases
		// (e.g. empty distance tree).
		return t, nil
	}
	// Verify the Kraft inequality holds (left = 1 means complete code).
	left := 1
	for l := 1; l <= 15; l++ {
		left <<= 1
		left -= t.counts[l]
		if left < 0 {
			return nil, errors.New("zlib: huffman: over-subscribed code")
		}
	}
	// A single-symbol code is allowed (left > 0 with one count) only for the
	// distance tree in certain encodings; we accept it permissively.
	// Build the symbol array in canonical order.
	offs[1] = 0
	for l := 1; l < 15; l++ {
		offs[l+1] = offs[l] + t.counts[l]
	}
	for sym, l := range lens {
		if l != 0 {
			t.symbol[offs[l]] = sym
			offs[l]++
		}
	}
	return t, nil
}

// bitReader consumes bits LSB-first from an io.Reader-backed buffer.
type bitReader struct {
	src  []byte
	pos  int
	buf  uint64 // accumulated bits, low-order = oldest
	cnt  uint   // number of valid bits in buf
	done bool
}

func (br *bitReader) need(n uint) error {
	for br.cnt < n {
		if br.pos >= len(br.src) {
			return io.ErrUnexpectedEOF
		}
		br.buf |= uint64(br.src[br.pos]) << br.cnt
		br.pos++
		br.cnt += 8
	}
	return nil
}

func (br *bitReader) readBits(n uint) (uint32, error) {
	if err := br.need(n); err != nil {
		return 0, err
	}
	v := uint32(br.buf & ((1 << n) - 1))
	br.buf >>= n
	br.cnt -= n
	return v, nil
}

func (br *bitReader) align() {
	// Discard bits up to next byte boundary.
	br.buf >>= br.cnt & 7
	br.cnt -= br.cnt & 7
	// Drop any buffered bytes that came after the alignment point — for
	// stored blocks we read straight from src.
	br.buf = 0
	br.cnt = 0
}

func (br *bitReader) readByte() (byte, error) {
	if br.pos >= len(br.src) {
		return 0, io.ErrUnexpectedEOF
	}
	b := br.src[br.pos]
	br.pos++
	return b, nil
}

func (br *bitReader) readBytes(n int) ([]byte, error) {
	if br.pos+n > len(br.src) {
		return nil, io.ErrUnexpectedEOF
	}
	out := br.src[br.pos : br.pos+n]
	br.pos += n
	return out, nil
}

// decodeSymbol pulls bits MSB-first within a code (LSB-first in bytes) and
// walks the canonical Huffman table. Returns the symbol or an error.
func (br *bitReader) decodeSymbol(t *huffTable) (int, error) {
	code := 0
	first := 0
	index := 0
	for l := 1; l <= 15; l++ {
		bit, err := br.readBits(1)
		if err != nil {
			return 0, err
		}
		code = (code << 1) | int(bit)
		count := t.counts[l]
		if code-count < first {
			return t.symbol[index+(code-first)], nil
		}
		index += count
		first = (first + count) << 1
	}
	return 0, errors.New("zlib: invalid huffman code")
}

// Decompress decodes a zlib stream and returns the uncompressed bytes.
func Decompress(src []byte) ([]byte, error) {
	if len(src) < 2 {
		return nil, errors.New("zlib: stream too short")
	}
	cmf := src[0]
	flg := src[1]
	if cmf&0x0f != 8 {
		return nil, fmt.Errorf("zlib: unsupported compression method %d", cmf&0x0f)
	}
	if (uint16(cmf)*256+uint16(flg))%31 != 0 {
		return nil, errors.New("zlib: header check failed")
	}
	if flg&0x20 != 0 {
		return nil, errors.New("zlib: preset dictionary not supported")
	}
	body := src[2 : len(src)-4]
	out, err := inflate(body)
	if err != nil {
		return nil, err
	}
	wantA := uint32(src[len(src)-4])<<24 | uint32(src[len(src)-3])<<16 |
		uint32(src[len(src)-2])<<8 | uint32(src[len(src)-1])
	if got := adler32(1, out); got != wantA {
		return nil, fmt.Errorf("zlib: adler32 mismatch: got=%08x want=%08x", got, wantA)
	}
	return out, nil
}

// DecompressRaw decodes a raw deflate stream (no zlib header/trailer).
func DecompressRaw(src []byte) ([]byte, error) {
	return inflate(src)
}

func inflate(src []byte) ([]byte, error) {
	br := &bitReader{src: src}
	// Lazily-built static trees (only first dynamic block needs them).
	staticLit, staticDist, err := buildStaticTrees()
	if err != nil {
		return nil, err
	}

	var out []byte
	for {
		bfinal, err := br.readBits(1)
		if err != nil {
			return nil, err
		}
		btype, err := br.readBits(2)
		if err != nil {
			return nil, err
		}
		switch btype {
		case 0: // stored
			br.align()
			if br.pos+4 > len(br.src) {
				return nil, io.ErrUnexpectedEOF
			}
			length := int(br.src[br.pos]) | int(br.src[br.pos+1])<<8
			nlen := int(br.src[br.pos+2]) | int(br.src[br.pos+3])<<8
			br.pos += 4
			if length^0xffff != nlen {
				return nil, errors.New("zlib: stored block length mismatch")
			}
			data, err := br.readBytes(length)
			if err != nil {
				return nil, err
			}
			out = append(out, data...)

		case 1: // static
			if err := inflateBlock(br, staticLit, staticDist, &out); err != nil {
				return nil, err
			}

		case 2: // dynamic
			lit, dist, err := readDynamicTrees(br)
			if err != nil {
				return nil, err
			}
			if err := inflateBlock(br, lit, dist, &out); err != nil {
				return nil, err
			}

		default:
			return nil, errors.New("zlib: invalid block type")
		}
		if bfinal != 0 {
			break
		}
	}
	return out, nil
}

func inflateBlock(br *bitReader, lit, dist *huffTable, out *[]byte) error {
	for {
		sym, err := br.decodeSymbol(lit)
		if err != nil {
			return err
		}
		if sym < 256 {
			*out = append(*out, byte(sym))
			continue
		}
		if sym == 256 {
			return nil
		}
		if sym > 285 {
			return errors.New("zlib: invalid literal/length symbol")
		}
		lcode := sym - 257
		length := lengthBase[lcode]
		if eb := lengthExtra[lcode]; eb > 0 {
			x, err := br.readBits(uint(eb))
			if err != nil {
				return err
			}
			length += int(x)
		}
		dsym, err := br.decodeSymbol(dist)
		if err != nil {
			return err
		}
		if dsym > 29 {
			return errors.New("zlib: invalid distance symbol")
		}
		d := distBase[dsym]
		if eb := distExtra[dsym]; eb > 0 {
			x, err := br.readBits(uint(eb))
			if err != nil {
				return err
			}
			d += int(x)
		}
		if d > len(*out) {
			return errors.New("zlib: distance back-reference before start of output")
		}
		// Copy length bytes from d back. Note: copy must be byte-by-byte
		// because src and dst may overlap (RLE).
		start := len(*out) - d
		for i := 0; i < length; i++ {
			*out = append(*out, (*out)[start+i])
		}
	}
}

func readDynamicTrees(br *bitReader) (*huffTable, *huffTable, error) {
	hlit, err := br.readBits(5)
	if err != nil {
		return nil, nil, err
	}
	hdist, err := br.readBits(5)
	if err != nil {
		return nil, nil, err
	}
	hclen, err := br.readBits(4)
	if err != nil {
		return nil, nil, err
	}
	nLit := int(hlit) + 257
	nDist := int(hdist) + 1
	nClen := int(hclen) + 4

	var clLens [19]int
	for i := 0; i < nClen; i++ {
		x, err := br.readBits(3)
		if err != nil {
			return nil, nil, err
		}
		clLens[codeLengthOrder[i]] = int(x)
	}
	clTree, err := buildHuffman(clLens[:])
	if err != nil {
		return nil, nil, err
	}

	all := make([]int, nLit+nDist)
	for i := 0; i < len(all); {
		sym, err := br.decodeSymbol(clTree)
		if err != nil {
			return nil, nil, err
		}
		switch {
		case sym < 16:
			all[i] = sym
			i++
		case sym == 16:
			if i == 0 {
				return nil, nil, errors.New("zlib: rep with no previous length")
			}
			x, err := br.readBits(2)
			if err != nil {
				return nil, nil, err
			}
			n := int(x) + 3
			for k := 0; k < n; k++ {
				all[i] = all[i-1]
				i++
			}
		case sym == 17:
			x, err := br.readBits(3)
			if err != nil {
				return nil, nil, err
			}
			i += int(x) + 3
		case sym == 18:
			x, err := br.readBits(7)
			if err != nil {
				return nil, nil, err
			}
			i += int(x) + 11
		default:
			return nil, nil, errors.New("zlib: invalid code-length symbol")
		}
	}
	lit, err := buildHuffman(all[:nLit])
	if err != nil {
		return nil, nil, err
	}
	dist, err := buildHuffman(all[nLit:])
	if err != nil {
		return nil, nil, err
	}
	return lit, dist, nil
}

func buildStaticTrees() (*huffTable, *huffTable, error) {
	litLens := make([]int, 288)
	for i := 0; i <= 143; i++ {
		litLens[i] = 8
	}
	for i := 144; i <= 255; i++ {
		litLens[i] = 9
	}
	for i := 256; i <= 279; i++ {
		litLens[i] = 7
	}
	for i := 280; i <= 287; i++ {
		litLens[i] = 8
	}
	distLens := make([]int, 30)
	for i := range distLens {
		distLens[i] = 5
	}
	lit, err := buildHuffman(litLens)
	if err != nil {
		return nil, nil, err
	}
	dist, err := buildHuffman(distLens)
	if err != nil {
		return nil, nil, err
	}
	return lit, dist, nil
}

// Reader is the io.ReadCloser-style decoder counterpart to Writer. Like
// Writer it is currently one-shot: ReadAll consumes the entire stream
// up-front and the decoded bytes are served from a buffer.
type Reader struct {
	buf []byte
	pos int
	err error
}

// NewReader reads the entire zlib stream from r and decodes it.
func NewReader(r io.Reader) (*Reader, error) {
	raw, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	out, err := Decompress(raw)
	if err != nil {
		return nil, err
	}
	return &Reader{buf: out}, nil
}

func (z *Reader) Read(p []byte) (int, error) {
	if z.err != nil {
		return 0, z.err
	}
	if z.pos >= len(z.buf) {
		z.err = io.EOF
		return 0, io.EOF
	}
	n := copy(p, z.buf[z.pos:])
	z.pos += n
	return n, nil
}

func (z *Reader) Close() error { return nil }
