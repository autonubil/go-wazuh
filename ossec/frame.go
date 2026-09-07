package ossec

import (
	"encoding/binary"
	"fmt"
	"io"
)

// WriteFrame writes payload to w prefixed with its length as a 4-byte
// little-endian header, matching the Wazuh TCP wire framing produced by the
// agent transport (see writeMessage). Header and payload are written together.
func WriteFrame(w io.Writer, payload []byte) error {
	if uint64(len(payload)) > 0xFFFFFFFF {
		return fmt.Errorf("ossec: frame payload too large: %d bytes", len(payload))
	}
	buf := make([]byte, 4+len(payload))
	binary.LittleEndian.PutUint32(buf[:4], uint32(len(payload)))
	copy(buf[4:], payload)
	_, err := w.Write(buf)
	return err
}

// ReadFrame reads one length-prefixed frame from r. The 4-byte little-endian
// length header is validated against maxLen before any payload buffer is
// allocated, so a hostile peer cannot force a large allocation. maxLen must be
// positive; a frame claiming more than maxLen bytes is rejected.
func ReadFrame(r io.Reader, maxLen uint32) ([]byte, error) {
	var hdr [4]byte
	if _, err := io.ReadFull(r, hdr[:]); err != nil {
		return nil, err
	}
	n := binary.LittleEndian.Uint32(hdr[:])
	if maxLen == 0 || n > maxLen {
		return nil, fmt.Errorf("ossec: frame length %d exceeds max %d", n, maxLen)
	}
	buf := make([]byte, n)
	if _, err := io.ReadFull(r, buf); err != nil {
		return nil, err
	}
	return buf, nil
}
