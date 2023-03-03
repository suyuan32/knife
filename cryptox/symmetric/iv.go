package symmetric

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
)

// IVFromBytes set IV data from byte slice.
func (s *CryptoS) IVFromBytes(data []byte) *CryptoS {
	s.IV = data
	return s
}

// IVFromString set IV data from string.
func (s *CryptoS) IVFromString(data string) *CryptoS {
	s.IV = []byte(data)
	return s
}

// IVFromBase64String set IV data from base64 string.
func (s *CryptoS) IVFromBase64String(data string) *CryptoS {
	result, err := base64.StdEncoding.DecodeString(data)
	s.Errors = errors.Join(s.Errors, err)
	s.IV = result
	return s
}

// IVFromHexString set IV data from hex string.
func (s *CryptoS) IVFromHexString(data string) *CryptoS {
	result, err := hex.DecodeString(data)
	s.Errors = errors.Join(s.Errors, err)
	s.IV = result
	return s
}
