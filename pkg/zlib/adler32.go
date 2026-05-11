package zlib

// Adler-32 checksum, RFC 1950. Ported from adler32.c. The constants and
// loop unrolling factor (NMAX) match libz so that intermediate `(a, b)` state
// is identical to libz at every chunk boundary.

const (
	adlerBase = 65521 // largest prime smaller than 65536
	adlerNMax = 5552  // largest n such that 255*n*(n+1)/2 + (n+1)*(BASE-1) <= 2^32-1
)

// adler32 returns the Adler-32 checksum of buf, with initial state adler
// (use 1 for the initial value used by deflate's zlib wrapper).
func adler32(adler uint32, buf []byte) uint32 {
	a := adler & 0xffff
	b := (adler >> 16) & 0xffff
	n := len(buf)
	for n > 0 {
		k := n
		if k > adlerNMax {
			k = adlerNMax
		}
		n -= k
		// Unroll loop by 16, matching libz. Equivalence is bit-exact at
		// every NMAX boundary because mod is taken at the same points.
		for k >= 16 {
			a += uint32(buf[0])
			b += a
			a += uint32(buf[1])
			b += a
			a += uint32(buf[2])
			b += a
			a += uint32(buf[3])
			b += a
			a += uint32(buf[4])
			b += a
			a += uint32(buf[5])
			b += a
			a += uint32(buf[6])
			b += a
			a += uint32(buf[7])
			b += a
			a += uint32(buf[8])
			b += a
			a += uint32(buf[9])
			b += a
			a += uint32(buf[10])
			b += a
			a += uint32(buf[11])
			b += a
			a += uint32(buf[12])
			b += a
			a += uint32(buf[13])
			b += a
			a += uint32(buf[14])
			b += a
			a += uint32(buf[15])
			b += a
			buf = buf[16:]
			k -= 16
		}
		for ; k > 0; k-- {
			a += uint32(buf[0])
			b += a
			buf = buf[1:]
		}
		a %= adlerBase
		b %= adlerBase
	}
	return (b << 16) | a
}
