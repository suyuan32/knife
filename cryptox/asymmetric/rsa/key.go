// Copyright 2023 The Ryan SU Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
)

var (
	errorNotValidPEMKey     = errors.New("the key must be a PEM string encoded by PKCS1 or PKCS8")
	errorNotValidPrivateKey = errors.New("the key is not a valid RSA private key")
	errorNotValidPublicKey  = errors.New(" the key is not a valid RSA public key")
)

// GenerateKeyPair set public key and private key for RSA struct.
func (s *RSA) GenerateKeyPair(bits int) *RSA {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		s.Errors = errors.Join(s.Errors, fmt.Errorf("rsa: generate key pair failed, err : %v", err))
		return s
	}
	s.PrivateKey = privateKey
	s.PublicKey = &privateKey.PublicKey
	return s
}

// PrivateKeyFromPEM gets private key from a PEM byte slice.
func (s *RSA) PrivateKeyFromPEM(data []byte) *RSA {
	block, _ := pem.Decode(data)
	if block == nil {
		s.Errors = errors.Join(s.Errors, errorNotValidPEMKey)
		return s
	}

	switch s.Standard {
	case PKCS1:
		if parse, err := x509.ParsePKCS1PrivateKey(block.Bytes); err != nil {
			s.Errors = errors.Join(s.Errors, errorNotValidPrivateKey)
			return s
		} else {
			s.PrivateKey = parse
			return s
		}
	case PKCS8:
		if parse, err := x509.ParsePKCS8PrivateKey(block.Bytes); err != nil {
			s.Errors = errors.Join(s.Errors, errorNotValidPrivateKey)
			return s
		} else {
			s.PrivateKey = parse.(*rsa.PrivateKey)
			return s
		}
	default:
		s.Errors = errors.Join(s.Errors, errorNotValidPrivateKey)
		return s
	}
}
