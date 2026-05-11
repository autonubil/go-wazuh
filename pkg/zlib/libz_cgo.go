//go:build zlibcompat

package zlib

/*
#cgo LDFLAGS: -lz
#include <stdlib.h>
#include <string.h>
#include <zlib.h>

static long czlib_compress(const unsigned char *src, unsigned long srclen,
                           int level, int window_bits, int mem_level, int strategy,
                           unsigned char **out) {
    z_stream zs;
    memset(&zs, 0, sizeof(zs));
    int rc = deflateInit2(&zs, level, Z_DEFLATED, window_bits, mem_level, strategy);
    if (rc != Z_OK) return rc;

    unsigned long cap = deflateBound(&zs, srclen) + 16;
    unsigned char *buf = (unsigned char *)malloc(cap);
    if (!buf) { deflateEnd(&zs); return Z_MEM_ERROR; }

    zs.next_in   = (Bytef *)src;
    zs.avail_in  = srclen;
    zs.next_out  = buf;
    zs.avail_out = cap;

    rc = deflate(&zs, Z_FINISH);
    if (rc != Z_STREAM_END) {
        free(buf);
        deflateEnd(&zs);
        return rc;
    }
    long n = (long)zs.total_out;
    deflateEnd(&zs);
    *out = buf;
    return n;
}
*/
import "C"

import "unsafe"

// czCompress calls libz's deflateInit2 + deflate(Z_FINISH) and returns the
// compressed bytes. Only available when built with the `zlibcompat` tag,
// which is what the conformance tests use.
func czCompress(src []byte, level, windowBits, memLevel, strategy int) ([]byte, int) {
	var cOut *C.uchar
	var srcPtr *C.uchar
	if len(src) > 0 {
		srcPtr = (*C.uchar)(unsafe.Pointer(&src[0]))
	}
	n := C.czlib_compress(srcPtr, C.ulong(len(src)),
		C.int(level), C.int(windowBits), C.int(memLevel), C.int(strategy), &cOut)
	if n < 0 {
		return nil, int(n)
	}
	defer C.free(unsafe.Pointer(cOut))
	return C.GoBytes(unsafe.Pointer(cOut), C.int(n)), 0
}
