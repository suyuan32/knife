package codec

import (
	"encoding/base32"
	"encoding/base64"
)

// Base32Encode encodes the string to base32 format.
func Base32Encode(data string) string {
	dst := make([]byte, base32.StdEncoding.EncodedLen(len(data)))
	base32.StdEncoding.Encode(dst, []byte(data))
	return string(dst)
}

// Base64Encode encodes the string to base64 format.
func Base64Encode(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}
