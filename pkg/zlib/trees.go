package zlib

// Port of trees.c: bit-level writer, Huffman tree construction, scan/send
// of tree representations, stored/static/dynamic block emission, _tr_tally.

// ---------------------------------------------------------------------------
// Bit writer.
// ---------------------------------------------------------------------------

func (s *state) putByte(b byte) { s.out = append(s.out, b) }

func (s *state) putShortLSB(w uint16) {
	s.out = append(s.out, byte(w&0xff), byte(w>>8))
}

// sendBits writes the low `length` bits of value to the bit stream (LSB first
// within bytes). Matches the C send_bits macro exactly so that two-byte
// flushes happen at the same points.
func (s *state) sendBits(value int, length int) {
	if s.biValid > bufSize-length {
		val := value
		s.biBuf |= uint16(val) << s.biValid
		s.putShortLSB(s.biBuf)
		s.biBuf = uint16(val) >> uint(bufSize-s.biValid)
		s.biValid += length - bufSize
	} else {
		s.biBuf |= uint16(value) << s.biValid
		s.biValid += length
	}
}

func (s *state) sendCode(c int, tree []ctData) {
	s.sendBits(int(tree[c].fc), int(tree[c].dl))
}

// biFlush flushes whole bytes from the bit buffer, keeping at most 7 bits.
func (s *state) biFlush() {
	if s.biValid == 16 {
		s.putShortLSB(s.biBuf)
		s.biBuf = 0
		s.biValid = 0
	} else if s.biValid >= 8 {
		s.putByte(byte(s.biBuf))
		s.biBuf >>= 8
		s.biValid -= 8
	}
}

// biWindup byte-aligns the bit stream.
func (s *state) biWindup() {
	if s.biValid > 8 {
		s.putShortLSB(s.biBuf)
	} else if s.biValid > 0 {
		s.putByte(byte(s.biBuf))
	}
	s.biBuf = 0
	s.biValid = 0
}

// biReverse: bit-reverse the low `len` bits of code.
func biReverse(code uint, length int) uint {
	res := uint(0)
	for {
		res |= code & 1
		code >>= 1
		length--
		if length <= 0 {
			break
		}
		res <<= 1
	}
	return res
}

// ---------------------------------------------------------------------------
// Tree construction.
// ---------------------------------------------------------------------------

func (s *state) initBlock() {
	for n := 0; n < lCodes; n++ {
		s.dynLTree[n].fc = 0
	}
	for n := 0; n < dCodes; n++ {
		s.dynDTree[n].fc = 0
	}
	for n := 0; n < blCodes; n++ {
		s.blTree[n].fc = 0
	}
	s.dynLTree[endBlock].fc = 1
	s.optLen = 0
	s.staticLen = 0
	s.symNext = 0
	s.matches = 0
}

func (s *state) trInit() {
	// l_desc, d_desc, bl_desc already point to dyn arrays + static descs (init()).
	s.biBuf = 0
	s.biValid = 0
	s.initBlock()
}

const smallest = 1

func (s *state) smaller(tree []ctData, n, m int) bool {
	if tree[n].fc < tree[m].fc {
		return true
	}
	return tree[n].fc == tree[m].fc && s.depth[n] <= s.depth[m]
}

func (s *state) pqdownheap(tree []ctData, k int) {
	v := s.heap[k]
	j := k << 1
	for j <= s.heapLen {
		if j < s.heapLen && s.smaller(tree, s.heap[j+1], s.heap[j]) {
			j++
		}
		if s.smaller(tree, v, s.heap[j]) {
			break
		}
		s.heap[k] = s.heap[j]
		k = j
		j <<= 1
	}
	s.heap[k] = v
}

// genBitlen: compute optimal bit lengths for tree, fill bl_count, update opt_len.
func (s *state) genBitlen(desc *treeDesc) {
	tree := desc.dynTree
	maxCode := desc.maxCode
	stree := desc.statDesc.staticTree
	extra := desc.statDesc.extraBits
	base := desc.statDesc.extraBase
	maxLength := desc.statDesc.maxLength

	for bits := 0; bits <= maxBits; bits++ {
		s.blCount[bits] = 0
	}

	tree[s.heap[s.heapMax]].dl = 0 // root of the heap

	overflow := 0
	for h := s.heapMax + 1; h < heapSize; h++ {
		n := s.heap[h]
		bits := int(tree[tree[n].dl].dl) + 1
		if bits > maxLength {
			bits = maxLength
			overflow++
		}
		tree[n].dl = uint16(bits)

		if n > maxCode {
			continue
		}
		s.blCount[bits]++
		xbits := 0
		if n >= base {
			xbits = extra[n-base]
		}
		f := uint32(tree[n].fc)
		s.optLen += f * uint32(bits+xbits)
		if stree != nil {
			s.staticLen += f * uint32(int(stree[n].dl)+xbits)
		}
	}
	if overflow == 0 {
		return
	}

	// Adjust bit counts to fit the maximum length constraint.
	for {
		bits := maxLength - 1
		for s.blCount[bits] == 0 {
			bits--
		}
		s.blCount[bits]--
		s.blCount[bits+1] += 2
		s.blCount[maxLength]--
		overflow -= 2
		if overflow <= 0 {
			break
		}
	}

	h := heapSize
	for bits := maxLength; bits != 0; bits-- {
		n := int(s.blCount[bits])
		for n != 0 {
			h--
			m := s.heap[h]
			if m > maxCode {
				continue
			}
			if int(tree[m].dl) != bits {
				s.optLen += (uint32(bits) - uint32(tree[m].dl)) * uint32(tree[m].fc)
				tree[m].dl = uint16(bits)
			}
			n--
		}
	}
}

// genCodes assigns canonical codes for the given tree and bit lengths.
func genCodes(tree []ctData, maxCode int, blCount *[maxBits + 1]uint16) {
	var nextCode [maxBits + 1]uint16
	code := uint16(0)
	for bits := 1; bits <= maxBits; bits++ {
		code = (code + blCount[bits-1]) << 1
		nextCode[bits] = code
	}
	for n := 0; n <= maxCode; n++ {
		length := int(tree[n].dl)
		if length == 0 {
			continue
		}
		tree[n].fc = uint16(biReverse(uint(nextCode[length]), length))
		nextCode[length]++
	}
}

// buildTree: full Huffman tree construction (heap, combine, gen_bitlen, gen_codes).
func (s *state) buildTree(desc *treeDesc) {
	tree := desc.dynTree
	stree := desc.statDesc.staticTree
	elems := desc.statDesc.elems

	maxCode := -1
	s.heapLen = 0
	s.heapMax = heapSize

	for n := 0; n < elems; n++ {
		if tree[n].fc != 0 {
			s.heapLen++
			s.heap[s.heapLen] = n
			maxCode = n
			s.depth[n] = 0
		} else {
			tree[n].dl = 0
		}
	}

	// Force at least two codes of non-zero frequency.
	for s.heapLen < 2 {
		s.heapLen++
		var node int
		if maxCode < 2 {
			maxCode++
			node = maxCode
		} else {
			node = 0
		}
		s.heap[s.heapLen] = node
		tree[node].fc = 1
		s.depth[node] = 0
		s.optLen--
		if stree != nil {
			s.staticLen -= uint32(stree[node].dl)
		}
	}
	desc.maxCode = maxCode

	for n := s.heapLen / 2; n >= 1; n-- {
		s.pqdownheap(tree, n)
	}

	node := elems
	for s.heapLen >= 2 {
		// pqremove
		n := s.heap[smallest]
		s.heap[smallest] = s.heap[s.heapLen]
		s.heapLen--
		s.pqdownheap(tree, smallest)
		m := s.heap[smallest]

		s.heapMax--
		s.heap[s.heapMax] = n
		s.heapMax--
		s.heap[s.heapMax] = m

		tree[node].fc = tree[n].fc + tree[m].fc
		if s.depth[n] >= s.depth[m] {
			s.depth[node] = s.depth[n] + 1
		} else {
			s.depth[node] = s.depth[m] + 1
		}
		tree[n].dl = uint16(node)
		tree[m].dl = uint16(node)

		s.heap[smallest] = node
		node++
		s.pqdownheap(tree, smallest)
	}

	s.heapMax--
	s.heap[s.heapMax] = s.heap[smallest]

	s.genBitlen(desc)
	genCodes(tree, maxCode, &s.blCount)
}

// scanTree counts the frequencies of bit-length codes that would be needed
// to encode tree's bit-length sequence.
func (s *state) scanTree(tree []ctData, maxCode int) {
	prevlen := -1
	nextlen := int(tree[0].dl)
	count := 0
	maxCount := 7
	minCount := 4
	if nextlen == 0 {
		maxCount = 138
		minCount = 3
	}
	// Guard at tree[maxCode+1].
	tree[maxCode+1].dl = 0xffff

	for n := 0; n <= maxCode; n++ {
		curlen := nextlen
		nextlen = int(tree[n+1].dl)
		count++
		if count < maxCount && curlen == nextlen {
			continue
		} else if count < minCount {
			s.blTree[curlen].fc += uint16(count)
		} else if curlen != 0 {
			if curlen != prevlen {
				s.blTree[curlen].fc++
			}
			s.blTree[repTreeS].fc++
		} else if count <= 10 {
			s.blTree[repZ3_10].fc++
		} else {
			s.blTree[repZ11_138].fc++
		}
		count = 0
		prevlen = curlen
		if nextlen == 0 {
			maxCount = 138
			minCount = 3
		} else if curlen == nextlen {
			maxCount = 6
			minCount = 3
		} else {
			maxCount = 7
			minCount = 4
		}
	}
}

// sendTree emits a literal/distance tree's bit-length sequence using bl_tree.
func (s *state) sendTree(tree []ctData, maxCode int) {
	prevlen := -1
	nextlen := int(tree[0].dl)
	count := 0
	maxCount := 7
	minCount := 4
	if nextlen == 0 {
		maxCount = 138
		minCount = 3
	}

	for n := 0; n <= maxCode; n++ {
		curlen := nextlen
		nextlen = int(tree[n+1].dl)
		count++
		if count < maxCount && curlen == nextlen {
			continue
		} else if count < minCount {
			for {
				s.sendCode(curlen, s.blTree[:])
				count--
				if count == 0 {
					break
				}
			}
		} else if curlen != 0 {
			if curlen != prevlen {
				s.sendCode(curlen, s.blTree[:])
				count--
			}
			s.sendCode(repTreeS, s.blTree[:])
			s.sendBits(count-3, 2)
		} else if count <= 10 {
			s.sendCode(repZ3_10, s.blTree[:])
			s.sendBits(count-3, 3)
		} else {
			s.sendCode(repZ11_138, s.blTree[:])
			s.sendBits(count-11, 7)
		}
		count = 0
		prevlen = curlen
		if nextlen == 0 {
			maxCount = 138
			minCount = 3
		} else if curlen == nextlen {
			maxCount = 6
			minCount = 3
		} else {
			maxCount = 7
			minCount = 4
		}
	}
}

// buildBLTree builds the Huffman tree for bit lengths and returns the index
// in bl_order of the last bit-length code to send (>= 3).
func (s *state) buildBLTree() int {
	s.scanTree(s.dynLTree[:], s.lDesc.maxCode)
	s.scanTree(s.dynDTree[:], s.dDesc.maxCode)
	s.buildTree(&s.blDesc)

	maxBlIndex := blCodes - 1
	for maxBlIndex >= 3 {
		if s.blTree[blOrder[maxBlIndex]].dl != 0 {
			break
		}
		maxBlIndex--
	}
	s.optLen += 3*uint32(maxBlIndex+1) + 5 + 5 + 4
	return maxBlIndex
}

// sendAllTrees: header for a dynamic-trees block.
func (s *state) sendAllTrees(lcodes, dcodes, blcodes int) {
	s.sendBits(lcodes-257, 5)
	s.sendBits(dcodes-1, 5)
	s.sendBits(blcodes-4, 4)
	for rank := 0; rank < blcodes; rank++ {
		s.sendBits(int(s.blTree[blOrder[rank]].dl), 3)
	}
	s.sendTree(s.dynLTree[:], lcodes-1)
	s.sendTree(s.dynDTree[:], dcodes-1)
}

// trStoredBlock emits a STORED (uncompressed) block.
func (s *state) trStoredBlock(buf []byte, last bool) {
	bit := 0
	if last {
		bit = 1
	}
	s.sendBits((storedBlock<<1)+bit, 3)
	s.biWindup()
	storedLen := uint16(len(buf))
	s.putShortLSB(storedLen)
	s.putShortLSB(^storedLen)
	s.out = append(s.out, buf...)
}

// trAlign emits one empty static block (10 bits). Used by Z_PARTIAL_FLUSH.
func (s *state) trAlign() {
	s.sendBits(staticTrees<<1, 3)
	s.sendCode(endBlock, staticLTree[:])
	s.biFlush()
}

// compressBlock emits a block's symbols using the supplied trees.
func (s *state) compressBlock(ltree, dtree []ctData) {
	sx := uint(0)
	if s.symNext != 0 {
		for {
			dist := uint(s.symBuf[sx]) & 0xff
			sx++
			dist += uint(s.symBuf[sx]) << 8
			sx++
			lc := int(s.symBuf[sx])
			sx++
			if dist == 0 {
				s.sendCode(lc, ltree)
			} else {
				code := int(lengthCode[lc])
				s.sendCode(code+literals+1, ltree)
				extra := extraLBits[code]
				if extra != 0 {
					lc -= baseLength[code]
					s.sendBits(lc, extra)
				}
				dist--
				dc := int(dCode(dist))
				s.sendCode(dc, dtree)
				extra = extraDBits[dc]
				if extra != 0 {
					dist -= uint(baseDist[dc])
					s.sendBits(int(dist), extra)
				}
			}
			if sx >= s.symNext {
				break
			}
		}
	}
	s.sendCode(endBlock, ltree)
}

// detectDataType: zlib's text/binary heuristic, used to set strm->data_type
// (does not affect output bytes — kept for parity).
func (s *state) detectDataType() int {
	blockMask := uint32(0xf3ffc07f)
	for n := 0; n <= 31; n++ {
		if blockMask&1 != 0 && s.dynLTree[n].fc != 0 {
			return 0 // binary
		}
		blockMask >>= 1
	}
	if s.dynLTree[9].fc != 0 || s.dynLTree[10].fc != 0 || s.dynLTree[13].fc != 0 {
		return 1 // text
	}
	for n := 32; n < literals; n++ {
		if s.dynLTree[n].fc != 0 {
			return 1
		}
	}
	return 0
}

// trFlushBlock chooses dynamic / static / stored encoding for the current
// block and writes it out.
func (s *state) trFlushBlock(buf []byte, storedLen uint32, last bool) {
	var optLenb, staticLenb uint32
	maxBlIndex := 0

	if s.level > 0 {
		s.buildTree(&s.lDesc)
		s.buildTree(&s.dDesc)
		maxBlIndex = s.buildBLTree()

		optLenb = (s.optLen + 3 + 7) >> 3
		staticLenb = (s.staticLen + 3 + 7) >> 3
		if staticLenb <= optLenb || s.strategy == StrategyFixed {
			optLenb = staticLenb
		}
	} else {
		// Force stored block.
		optLenb = storedLen + 5
		staticLenb = optLenb
		_ = staticLenb
	}

	bit := 0
	if last {
		bit = 1
	}

	if storedLen+4 <= optLenb && buf != nil {
		s.trStoredBlock(buf, last)
	} else if staticLenb == optLenb {
		s.sendBits((staticTrees<<1)+bit, 3)
		s.compressBlock(staticLTree[:], staticDTree[:])
	} else {
		s.sendBits((dynTrees<<1)+bit, 3)
		s.sendAllTrees(s.lDesc.maxCode+1, s.dDesc.maxCode+1, maxBlIndex+1)
		s.compressBlock(s.dynLTree[:], s.dynDTree[:])
	}

	s.initBlock()
	if last {
		s.biWindup()
	}
}

// ---------------------------------------------------------------------------
// _tr_tally helpers. Match the (dist, lc) encoding in deflate.h's
// _tr_tally_lit / _tr_tally_dist macros exactly.
// ---------------------------------------------------------------------------

func (s *state) trTallyLit(c byte) bool {
	s.symBuf[s.symNext] = 0
	s.symBuf[s.symNext+1] = 0
	s.symBuf[s.symNext+2] = c
	s.symNext += 3
	s.dynLTree[c].fc++
	return s.symNext == s.symEnd
}

func (s *state) trTallyDist(dist, lc uint) bool {
	s.symBuf[s.symNext] = byte(dist)
	s.symBuf[s.symNext+1] = byte(dist >> 8)
	s.symBuf[s.symNext+2] = byte(lc)
	s.symNext += 3
	s.matches++
	dist--
	s.dynLTree[int(lengthCode[lc])+literals+1].fc++
	s.dynDTree[dCode(dist)].fc++
	return s.symNext == s.symEnd
}
