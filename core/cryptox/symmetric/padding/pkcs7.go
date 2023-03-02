package padding

import (
	"bytes"
)

// PaddingPKCS7 add padding at the end of byte slice with PKCS7 bytes.
// PKCS7 padding is a generalization of PKCS5 padding (also known as standard padding). PKCS7 padding works by appending N bytes with the value of chr(N) , where N is the number of bytes required to make the final block of data the same size as the block size.
func PaddingPKCS7(data []byte, blockSize int) []byte {
	dataLen := len(data)
	if dataLen == 0 || blockSize < 1 {
		return data
	}
	paddingSize := blockSize - (dataLen % blockSize)
	return append(data, bytes.Repeat([]byte{byte(paddingSize)}, paddingSize)...)
}

// DePaddingPKCS7 remove PKCS7 padding at the end of byte slice.
func DePaddingPKCS7(data []byte) []byte {
	if dataLen := len(data); dataLen == 0 {
		return data
	} else {
		if dataLeft := dataLen - int(data[dataLen-1]); dataLeft > 0 {
			return data[:dataLeft]
		} else {
			return data
		}
	}
}

// PaddingPKCS5 is similar with PKCS7, its block size is 8.
func PaddingPKCS5(data []byte) []byte {
	return PaddingPKCS7(data, 8)
}

// DePaddingPKCS5 is similar with PKCS7, its block size is 8.
func DePaddingPKCS5(data []byte) []byte {
	return DePaddingPKCS7(data)
}
