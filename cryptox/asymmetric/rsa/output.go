package rsa

import (
	"encoding/base64"
	"encoding/hex"
)

// ToString output data with string type.
func (s *RSA) ToString() (string, error) {
	return string(s.OutputData), s.Errors
}

// ToBytes output data with byte type.
func (s *RSA) ToBytes() ([]byte, error) {
	return s.OutputData, s.Errors
}

// ToBase64String output data with base64 string.
func (s *RSA) ToBase64String() (string, error) {
	return base64.StdEncoding.EncodeToString(s.OutputData), s.Errors
}

// ToHexString output data with hex string.
func (s *RSA) ToHexString() (string, error) {
	return hex.EncodeToString(s.OutputData), s.Errors
}
