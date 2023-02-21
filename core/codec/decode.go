package codec

import (
	"encoding/base32"
	"encoding/base64"
	"fmt"
)

// Base32Decode decodes base32 string.
func Base32Decode(data string) (string, error) {
	result, err := base32.StdEncoding.DecodeString(data)
	if err != nil {
		return "", fmt.Errorf("failed to decode string to base32, err: %v", err)
	}
	return string(result), nil
}

// Base64Decode decodes base64 string.
func Base64Decode(data string) (string, error) {
	decodeString, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", fmt.Errorf("failed to decode string to base64, err: %v", err)
	}
	return string(decodeString), nil
}
