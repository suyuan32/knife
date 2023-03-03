package symmetric

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
)

// KeyFromBytes set key data from byte slice.
func (s *CryptoS) KeyFromBytes(data []byte) *CryptoS {
	s.Key = data
	return s
}

// KeyFromString set key data from string.
func (s *CryptoS) KeyFromString(data string) *CryptoS {
	s.Key = []byte(data)
	return s
}

// KeyFromBase64String set key data from base64 string.
func (s *CryptoS) KeyFromBase64String(data string) *CryptoS {
	result, err := base64.StdEncoding.DecodeString(data)
	s.Errors = errors.Join(s.Errors, err)
	s.Key = result
	return s
}

// KeyFromHexString set key data from hex string.
func (s *CryptoS) KeyFromHexString(data string) *CryptoS {
	result, err := hex.DecodeString(data)
	s.Errors = errors.Join(s.Errors, err)
	s.Key = result
	return s
}
