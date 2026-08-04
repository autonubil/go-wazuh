package zlib

// Port of deflate.c: state, fill_window, longest_match, deflate_slow,
// deflate_fast, deflate_stored, and the top-level Compress entry point.
// Identifiers follow the C source so future cross-checks are easy.

import (
	"errors"
	"io"
)

// state is the deflate_state from deflate.h. Field order follows the C struct
// where it matters for code locality (it does not affect output).
type state struct {
	// Output buffer accumulated by trees.go and put_byte/put_short helpers.
	out []byte

	// Wrap: 0 = raw deflate, 1 = zlib. Gzip (2) is not implemented (yet).
	wrap int

	level    int
	strategy int

	// Window parameters.
	wBits      uint
	wSize      uint
	wMask      uint
	windowSize uint
	window     []byte

	// LZ77 hash table.
	prev      []uint16
	head      []uint16
	insH      uint
	hashSize  uint
	hashBits  uint
	hashMask  uint
	hashShift uint

	blockStart int // can be negative when window slides
	strstart   uint
	lookahead  uint
	matchStart uint

	matchLength    uint
	prevLength     uint
	prevMatch      uint
	matchAvailable int

	maxChainLength uint
	maxLazyMatch   uint
	goodMatch      uint
	niceMatch      int

	insert uint

	// Trees (literal/length, distance, bit-length).
	dynLTree [heapSize]ctData
	dynDTree [2*dCodes + 1]ctData
	blTree   [2*blCodes + 1]ctData

	lDesc  treeDesc
	dDesc  treeDesc
	blDesc treeDesc

	blCount [maxBits + 1]uint16

	heap    [2*lCodes + 1]int
	heapLen int
	heapMax int

	depth [2*lCodes + 1]byte

	// Symbol buffer accumulated by _tr_tally.
	symBuf  []byte
	symNext uint
	symEnd  uint

	optLen    uint32
	staticLen uint32
	matches   uint
	dataType  int

	// Bit buffer for output.
	biBuf   uint16
	biValid int

	// Adler32 of uncompressed input (when wrap == 1).
	adler uint32

	highWater uint
}

type treeDesc struct {
	dynTree  []ctData
	maxCode  int
	statDesc *staticTreeDesc
}

type staticTreeDesc struct {
	staticTree []ctData
	extraBits  []int
	extraBase  int
	elems      int
	maxLength  int
}

var (
	staticLDescV = staticTreeDesc{
		staticTree: staticLTree[:],
		extraBits:  extraLBits[:],
		extraBase:  literals + 1,
		elems:      lCodes,
		maxLength:  maxBits,
	}
	staticDDescV = staticTreeDesc{
		staticTree: staticDTree[:],
		extraBits:  extraDBits[:],
		extraBase:  0,
		elems:      dCodes,
		maxLength:  maxBits,
	}
	staticBLDescV = staticTreeDesc{
		staticTree: nil,
		extraBits:  extraBLBits[:],
		extraBase:  0,
		elems:      blCodes,
		maxLength:  maxBLBits,
	}
)

// maxDist matches the C macro: (w_size - MIN_LOOKAHEAD).
func (s *state) maxDist() uint { return s.wSize - minLookahead }

// init initializes a fresh deflate stream for the given parameters.
func (s *state) init(level, windowBits, memLevel, strategy int, wrap int) error {
	if level == -1 {
		level = 6
	}
	if level < 0 || level > 9 {
		return errors.New("zlib: invalid level")
	}
	if windowBits < 8 || windowBits > 15 {
		return errors.New("zlib: invalid windowBits")
	}
	if memLevel < 1 || memLevel > 9 {
		return errors.New("zlib: invalid memLevel")
	}
	if strategy < 0 || strategy > StrategyFixed {
		return errors.New("zlib: invalid strategy")
	}
	if windowBits == 8 {
		windowBits = 9 // libz quirk: "until 256-byte window bug fixed"
	}

	s.wrap = wrap
	s.level = level
	s.strategy = strategy

	s.wBits = uint(windowBits)
	s.wSize = 1 << s.wBits
	s.wMask = s.wSize - 1

	s.hashBits = uint(memLevel) + 7
	s.hashSize = 1 << s.hashBits
	s.hashMask = s.hashSize - 1
	s.hashShift = (s.hashBits + minMatch - 1) / minMatch

	s.window = make([]byte, 2*s.wSize)
	s.prev = make([]uint16, s.wSize)
	s.head = make([]uint16, s.hashSize)

	// lit_bufsize = 1 << (memLevel + 6). symBuf has 3*lit_bufsize bytes.
	litBufsize := uint(1) << (uint(memLevel) + 6)
	s.symBuf = make([]byte, 3*litBufsize)
	s.symEnd = (litBufsize - 1) * 3

	s.adler = 1

	s.lDesc.dynTree = s.dynLTree[:]
	s.lDesc.statDesc = &staticLDescV
	s.dDesc.dynTree = s.dynDTree[:]
	s.dDesc.statDesc = &staticDDescV
	s.blDesc.dynTree = s.blTree[:]
	s.blDesc.statDesc = &staticBLDescV

	s.biBuf = 0
	s.biValid = 0

	s.initBlock()
	s.lmInit()
	return nil
}

func (s *state) lmInit() {
	s.windowSize = 2 * s.wSize
	for i := range s.head {
		s.head[i] = 0
	}
	cfg := configTable[s.level]
	s.maxLazyMatch = uint(cfg.maxLazy)
	s.goodMatch = uint(cfg.goodLength)
	s.niceMatch = int(cfg.niceLength)
	s.maxChainLength = uint(cfg.maxChain)

	s.strstart = 0
	s.blockStart = 0
	s.lookahead = 0
	s.insert = 0
	s.matchLength = minMatch - 1
	s.prevLength = minMatch - 1
	s.matchAvailable = 0
	s.insH = 0
}

// updateHash mirrors zlib's UPDATE_HASH macro.
func (s *state) updateHash(h uint, c byte) uint {
	return ((h << s.hashShift) ^ uint(c)) & s.hashMask
}

// insertString inserts strstart into the hash table and returns the previous
// chain head (NIL == 0). Mirrors zlib's INSERT_STRING macro.
func (s *state) insertString(str uint) uint16 {
	s.insH = s.updateHash(s.insH, s.window[str+minMatch-1])
	matchHead := s.head[s.insH]
	s.prev[str&s.wMask] = matchHead
	s.head[s.insH] = uint16(str)
	return matchHead
}

// slideHash slides the head/prev arrays when the window is shifted.
func (s *state) slideHash() {
	wsize := uint16(s.wSize)
	for i := range s.head {
		m := s.head[i]
		if m >= wsize {
			s.head[i] = m - wsize
		} else {
			s.head[i] = 0
		}
	}
	for i := 0; i < int(s.wSize); i++ {
		m := s.prev[i]
		if m >= wsize {
			s.prev[i] = m - wsize
		} else {
			s.prev[i] = 0
		}
	}
}

// fillWindow ports fill_window. It pulls bytes from `in` into the window and
// updates the hash for any inserted strings. fillWindow is only called when
// lookahead < MIN_LOOKAHEAD.
//
// On entry: *inPos is the next byte to consume from in.
// On exit: lookahead has been refilled (as much as in allows).
func (s *state) fillWindow(in []byte, inPos *int) {
	wsize := s.wSize
	for {
		more := s.windowSize - s.lookahead - s.strstart

		// If the window is almost full and there is insufficient lookahead,
		// move the upper half to the lower one to make room in the upper half.
		if s.strstart >= wsize+s.maxDist() {
			copy(s.window[:wsize], s.window[wsize:2*wsize])
			s.matchStart -= wsize
			s.strstart -= wsize
			s.blockStart -= int(wsize)
			if s.insert > s.strstart {
				s.insert = s.strstart
			}
			s.slideHash()
			more += wsize
		}
		if *inPos >= len(in) {
			break
		}

		n := uint(len(in) - *inPos)
		if n > more {
			n = more
		}
		// Update adler32 over the bytes we're about to read.
		if s.wrap == 1 {
			s.adler = adler32(s.adler, in[*inPos:*inPos+int(n)])
		}
		copy(s.window[s.strstart+s.lookahead:], in[*inPos:*inPos+int(n)])
		*inPos += int(n)
		s.lookahead += n

		// Initialize the hash value now that we have some input.
		if s.lookahead+s.insert >= minMatch {
			str := s.strstart - s.insert
			s.insH = uint(s.window[str])
			s.insH = s.updateHash(s.insH, s.window[str+1])
			for s.insert != 0 {
				s.insH = s.updateHash(s.insH, s.window[str+minMatch-1])
				s.prev[str&s.wMask] = s.head[s.insH]
				s.head[s.insH] = uint16(str)
				str++
				s.insert--
				if s.lookahead+s.insert < minMatch {
					break
				}
			}
		}

		if s.lookahead >= minLookahead || *inPos >= len(in) {
			break
		}
	}
	// We do not need WIN_INIT zero-fill: Go's make() already zero-initializes
	// the window, and longest_match in this port does not read past the end
	// of valid data.
}

// longestMatch ports longest_match from deflate.c (non-FASTEST path).
// It uses the C semantics: starts from best_len = prev_length, and discards
// any match <= prev_length. Returns the new best length (clamped to lookahead).
func (s *state) longestMatch(curMatch uint) uint {
	chainLength := s.maxChainLength
	scan := s.strstart
	bestLen := int(s.prevLength)
	niceMatchU := s.niceMatch

	limit := uint(0)
	if s.strstart > s.maxDist() {
		limit = s.strstart - s.maxDist()
	}

	wmask := s.wMask
	strend := scan + maxMatch
	scanEnd1 := s.window[scan+uint(bestLen)-1]
	scanEnd := s.window[scan+uint(bestLen)]

	if s.prevLength >= s.goodMatch {
		chainLength >>= 2
	}
	if uint(niceMatchU) > s.lookahead {
		niceMatchU = int(s.lookahead)
	}

	for {
		match := curMatch
		// Skip to next match if the match length cannot increase or if the
		// match length is less than 2.
		if s.window[match+uint(bestLen)] != scanEnd ||
			s.window[match+uint(bestLen)-1] != scanEnd1 ||
			s.window[match] != s.window[scan] ||
			s.window[match+1] != s.window[scan+1] {
			goto next
		}

		// scan += 2, match++; scan[2] == match[2] guaranteed by hash equality.
		{
			scanP := scan + 2
			matchP := match + 2
			// Unrolled by 8, matching the C inner loop.
			for {
				if s.window[scanP] != s.window[matchP] {
					break
				}
				scanP++
				matchP++
				if s.window[scanP] != s.window[matchP] {
					break
				}
				scanP++
				matchP++
				if s.window[scanP] != s.window[matchP] {
					break
				}
				scanP++
				matchP++
				if s.window[scanP] != s.window[matchP] {
					break
				}
				scanP++
				matchP++
				if s.window[scanP] != s.window[matchP] {
					break
				}
				scanP++
				matchP++
				if s.window[scanP] != s.window[matchP] {
					break
				}
				scanP++
				matchP++
				if s.window[scanP] != s.window[matchP] {
					break
				}
				scanP++
				matchP++
				if s.window[scanP] != s.window[matchP] {
					break
				}
				scanP++
				matchP++
				if scanP >= strend {
					break
				}
			}
			length := maxMatch - int(strend-scanP)
			if length > bestLen {
				s.matchStart = curMatch
				bestLen = length
				if length >= niceMatchU {
					break
				}
				scanEnd1 = s.window[scan+uint(bestLen)-1]
				scanEnd = s.window[scan+uint(bestLen)]
			}
		}

	next:
		curMatch = uint(s.prev[curMatch&wmask])
		if curMatch <= limit {
			break
		}
		chainLength--
		if chainLength == 0 {
			break
		}
	}

	if uint(bestLen) <= s.lookahead {
		return uint(bestLen)
	}
	return s.lookahead
}

// blockState mirrors the C enum.
type blockState int

const (
	needMore blockState = iota
	blockDone
	finishStarted
	finishDone
)

// flushBlockOnly emits the current block to s.out. last==true marks the last
// block. The block contents come from window[blockStart..strstart].
func (s *state) flushBlockOnly(last bool) {
	var buf []byte
	if s.blockStart >= 0 {
		buf = s.window[uint(s.blockStart):s.strstart]
	}
	storedLen := uint32(int(s.strstart) - s.blockStart)
	s.trFlushBlock(buf, storedLen, last)
	s.blockStart = int(s.strstart)
}

// deflateSlow ports deflate_slow (levels 4..9, lazy matching). The Z_FINISH
// path is the only one we exercise (Compress is one-shot).
func (s *state) deflateSlow(in []byte, inPos *int) blockState {
	for {
		// Make sure we have enough lookahead.
		if s.lookahead < minLookahead {
			s.fillWindow(in, inPos)
			if s.lookahead < minLookahead && *inPos < len(in) {
				return needMore
			}
			if s.lookahead == 0 {
				break
			}
		}

		// INSERT_STRING for the current string.
		var hashHead uint16
		if s.lookahead >= minMatch {
			hashHead = s.insertString(s.strstart)
		}

		// Save the previous match info.
		s.prevLength = s.matchLength
		s.prevMatch = s.matchStart
		s.matchLength = minMatch - 1

		if hashHead != 0 && s.prevLength < s.maxLazyMatch &&
			s.strstart-uint(hashHead) <= s.maxDist() {
			s.matchLength = s.longestMatch(uint(hashHead))

			if s.matchLength <= 5 && (s.strategy == StrategyFiltered ||
				(s.matchLength == minMatch && s.strstart-s.matchStart > tooFar)) {
				s.matchLength = minMatch - 1
			}
		}

		// If there was a match at the previous step and the current match is
		// not better, output the previous match.
		if s.prevLength >= minMatch && s.matchLength <= s.prevLength {
			maxInsert := s.strstart + s.lookahead - minMatch

			bflush := s.trTallyDist(s.strstart-1-s.prevMatch, s.prevLength-minMatch)

			s.lookahead -= s.prevLength - 1
			s.prevLength -= 2
			for s.prevLength != 0 {
				s.strstart++
				if s.strstart <= maxInsert {
					_ = s.insertString(s.strstart)
				}
				s.prevLength--
			}
			s.matchAvailable = 0
			s.matchLength = minMatch - 1
			s.strstart++

			if bflush {
				s.flushBlockOnly(false)
			}
		} else if s.matchAvailable != 0 {
			// Truncate the previous match to a single literal.
			bflush := s.trTallyLit(s.window[s.strstart-1])
			if bflush {
				s.flushBlockOnly(false)
			}
			s.strstart++
			s.lookahead--
		} else {
			s.matchAvailable = 1
			s.strstart++
			s.lookahead--
		}
	}

	if s.matchAvailable != 0 {
		_ = s.trTallyLit(s.window[s.strstart-1])
		s.matchAvailable = 0
	}
	if s.strstart < minMatch-1 {
		s.insert = s.strstart
	} else {
		s.insert = minMatch - 1
	}
	s.flushBlockOnly(true)
	return finishDone
}

// deflateFast ports deflate_fast (levels 1..3, no lazy matching).
func (s *state) deflateFast(in []byte, inPos *int) blockState {
	for {
		if s.lookahead < minLookahead {
			s.fillWindow(in, inPos)
			if s.lookahead < minLookahead && *inPos < len(in) {
				return needMore
			}
			if s.lookahead == 0 {
				break
			}
		}

		var hashHead uint16
		if s.lookahead >= minMatch {
			hashHead = s.insertString(s.strstart)
		}

		if hashHead != 0 && s.strstart-uint(hashHead) <= s.maxDist() {
			s.matchLength = s.longestMatch(uint(hashHead))
		}

		var bflush bool
		if s.matchLength >= minMatch {
			bflush = s.trTallyDist(s.strstart-s.matchStart, s.matchLength-minMatch)
			s.lookahead -= s.matchLength

			// Insert strings up to the end of the match. For levels <=3,
			// max_lazy is used as max_insert_length.
			if s.matchLength <= s.maxLazyMatch && s.lookahead >= minMatch {
				s.matchLength--
				for {
					s.strstart++
					_ = s.insertString(s.strstart)
					s.matchLength--
					if s.matchLength == 0 {
						break
					}
				}
				s.strstart++
			} else {
				s.strstart += s.matchLength
				s.matchLength = 0
				s.insH = uint(s.window[s.strstart])
				s.insH = s.updateHash(s.insH, s.window[s.strstart+1])
			}
		} else {
			bflush = s.trTallyLit(s.window[s.strstart])
			s.lookahead--
			s.strstart++
		}
		if bflush {
			s.flushBlockOnly(false)
		}
	}
	if s.strstart < minMatch-1 {
		s.insert = s.strstart
	} else {
		s.insert = minMatch - 1
	}
	s.flushBlockOnly(true)
	return finishDone
}

// deflateStored emits stored (uncompressed) blocks for level 0. We update the
// adler32 over the input as a single pass, then emit chunks of at most 65535
// bytes each — matching libz's behavior of one final block per call for
// small inputs and a sequence of max-sized blocks for larger ones.
func (s *state) deflateStored(in []byte, inPos *int) blockState {
	const maxStored = 65535

	if s.wrap == 1 && *inPos < len(in) {
		s.adler = adler32(s.adler, in[*inPos:])
	}

	rem := len(in) - *inPos
	for {
		n := rem
		if n > maxStored {
			n = maxStored
		}
		rem -= n
		last := rem == 0
		buf := in[*inPos : *inPos+n]
		*inPos += n
		s.trStoredBlock(buf, last)
		if last {
			break
		}
	}
	return finishDone
}

// ---------------------------------------------------------------------------
// Public API.
// ---------------------------------------------------------------------------

// Compress compresses src using deflate at the given level (0..9, or -1 for
// the default level 6). The output is a zlib stream (RFC 1950): two-byte
// header, deflate body, four-byte Adler-32 trailer. The output is byte-
// identical to libz's `compress2()` / `deflateInit + deflate(Z_FINISH)` for
// the same level, default windowBits (15), default memLevel (8), and default
// strategy (Z_DEFAULT_STRATEGY).
func Compress(src []byte, level int) ([]byte, error) {
	return CompressWith(src, level, defaultWindowBits, defaultMemLevel, StrategyDefault)
}

// CompressWith is the full-control variant. Negative windowBits selects raw
// deflate (no zlib header/trailer), matching libz semantics.
func CompressWith(src []byte, level, windowBits, memLevel, strategy int) ([]byte, error) {
	wrap := 1
	if windowBits < 0 {
		wrap = 0
		windowBits = -windowBits
	}
	var s state
	if err := s.init(level, windowBits, memLevel, strategy, wrap); err != nil {
		return nil, err
	}

	// Pre-grow output to a reasonable upper bound to avoid repeated reallocs.
	s.out = make([]byte, 0, deflateBound(len(src), wrap == 1))

	// zlib header (RFC 1950), matching the bit layout that libz writes.
	if wrap == 1 {
		s.writeZlibHeader()
	}

	// Initialize trees state (only once).
	s.trInit()

	pos := 0
	if s.level == 0 || strategy == StrategyHuffman || strategy == StrategyRLE {
		// For levels >= 1, strategy overrides choose deflate_huff/deflate_rle.
		switch {
		case s.level == 0:
			s.deflateStored(src, &pos)
		case strategy == StrategyHuffman:
			s.deflateHuff(src, &pos)
		case strategy == StrategyRLE:
			s.deflateRLE(src, &pos)
		}
	} else if configTable[s.level].useSlow {
		s.deflateSlow(src, &pos)
	} else {
		s.deflateFast(src, &pos)
	}

	// Adler-32 trailer for zlib wrap.
	if wrap == 1 {
		a := s.adler
		s.out = append(s.out, byte(a>>24), byte(a>>16), byte(a>>8), byte(a))
	}
	return s.out, nil
}

func (s *state) writeZlibHeader() {
	header := uint16((8 + ((s.wBits - 8) << 4)) << 8)
	var levelFlags uint16
	if s.strategy >= StrategyHuffman || s.level < 2 {
		levelFlags = 0
	} else if s.level < 6 {
		levelFlags = 1
	} else if s.level == 6 {
		levelFlags = 2
	} else {
		levelFlags = 3
	}
	header |= levelFlags << 6
	header += 31 - (header % 31)
	s.out = append(s.out, byte(header>>8), byte(header&0xff))
}

func deflateBound(srcLen int, wrap bool) int {
	// Tight bound for default settings: srcLen + srcLen/4096 + srcLen/16384 +
	// srcLen/33554432 + 13 - 6 + wraplen. We are generous here.
	n := srcLen + srcLen>>12 + srcLen>>14 + srcLen>>25 + 13
	if wrap {
		n += 6
	}
	return n
}

// deflateHuff: strategy Z_HUFFMAN_ONLY. Just emits literals, no matching.
func (s *state) deflateHuff(in []byte, inPos *int) blockState {
	for {
		if s.lookahead == 0 {
			s.fillWindow(in, inPos)
			if s.lookahead == 0 {
				break
			}
		}
		s.matchLength = 0
		bflush := s.trTallyLit(s.window[s.strstart])
		s.lookahead--
		s.strstart++
		if bflush {
			s.flushBlockOnly(false)
		}
	}
	s.insert = 0
	s.flushBlockOnly(true)
	return finishDone
}

// deflateRLE: strategy Z_RLE. Look for runs of bytes (matches of distance 1
// only). Does not maintain the hash table.
func (s *state) deflateRLE(in []byte, inPos *int) blockState {
	for {
		if s.lookahead <= maxMatch {
			s.fillWindow(in, inPos)
			if s.lookahead <= maxMatch && *inPos < len(in) {
				return needMore
			}
			if s.lookahead == 0 {
				break
			}
		}
		s.matchLength = 0
		if s.lookahead >= minMatch && s.strstart > 0 {
			scan := s.strstart - 1
			prev := s.window[scan]
			if prev == s.window[scan+1] && prev == s.window[scan+2] && prev == s.window[scan+3] {
				strend := s.strstart + maxMatch
				p := scan + 1
				for p < strend && s.window[p] == prev {
					p++
				}
				s.matchLength = maxMatch - (strend - p)
				if s.matchLength > s.lookahead {
					s.matchLength = s.lookahead
				}
			}
		}
		var bflush bool
		if s.matchLength >= minMatch {
			bflush = s.trTallyDist(1, s.matchLength-minMatch)
			s.lookahead -= s.matchLength
			s.strstart += s.matchLength
			s.matchLength = 0
		} else {
			bflush = s.trTallyLit(s.window[s.strstart])
			s.lookahead--
			s.strstart++
		}
		if bflush {
			s.flushBlockOnly(false)
		}
	}
	s.insert = 0
	s.flushBlockOnly(true)
	return finishDone
}

// ---------------------------------------------------------------------------
// io.Writer-style streaming API (one-shot for now: Write buffers, Close emits).
// Streaming with intermediate flushes is not yet implemented because it requires
// chunked deflate() returns; the one-shot path is sufficient for byte-identical
// matching against compress2().
// ---------------------------------------------------------------------------

// Writer is a streaming-style zlib encoder. Currently buffers all writes and
// emits on Close. NewWriter / NewWriterLevel / NewWriterLevelDict mirror the
// stdlib API to ease drop-in replacement.
type Writer struct {
	w     io.Writer
	level int
	buf   []byte
	err   error
}

// NewWriter creates a writer using default compression (level 6).
func NewWriter(w io.Writer) *Writer { return &Writer{w: w, level: 6} }

// NewWriterLevel creates a writer using the given compression level.
func NewWriterLevel(w io.Writer, level int) (*Writer, error) {
	if level == -1 {
		level = 6
	}
	if level < 0 || level > 9 {
		return nil, errors.New("zlib: invalid level")
	}
	return &Writer{w: w, level: level}, nil
}

func (z *Writer) Write(p []byte) (int, error) {
	if z.err != nil {
		return 0, z.err
	}
	z.buf = append(z.buf, p...)
	return len(p), nil
}

func (z *Writer) Close() error {
	if z.err != nil {
		return z.err
	}
	out, err := Compress(z.buf, z.level)
	if err != nil {
		z.err = err
		return err
	}
	_, err = z.w.Write(out)
	z.err = err
	return err
}
