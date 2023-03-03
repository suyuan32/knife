package padding

import (
	"bytes"
)

// PaddingZero add padding at the end of byte slice with byte 0.
func PaddingZero(data []byte, blockSize int) []byte {
	dataLen := len(data)
	if dataLen == 0 || blockSize < 1 {
		return data
	}
	return append(data, bytes.Repeat([]byte{byte(0)}, blockSize-(dataLen%blockSize))...)
}

// DePaddingZero remove zero padding at the end of byte slice.
func DePaddingZero(data []byte) []byte {
	return bytes.TrimRight(data, string([]byte{0}))
}
