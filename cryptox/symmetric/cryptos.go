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

package symmetric

import (
	"github.com/suyuan32/knife/cryptox/symmetric/method"
	"github.com/suyuan32/knife/cryptox/symmetric/mode"
	"github.com/suyuan32/knife/cryptox/symmetric/padding"
)

// CryptoS is the struct for symmetric-key algorithm.
type CryptoS struct {
	// InputData is the data to be encrypted.
	InputData []byte

	// OutputData is the data encrypted.
	OutputData []byte

	// Key is the secret used for encrypted
	Key []byte

	// An initialization vector (IV) is an arbitrary number that can be used with a secret key for data encryption
	// to foil cyberattacks. This number, also called a nonce (number used once), is employed only one time in
	// any session to prevent unauthorized decryption of the message by a suspicious or malicious actor.
	IV []byte

	// Method is the encryption method such as AES.
	Method method.MethodType

	// Mode is the encrypted mode such as ECB.
	Mode mode.ModeType

	// Padding is the padding method such as PKCS7.
	Padding padding.PaddingType

	// Errors is the errors
	Errors error
}

func NewCryptoS() CryptoS {
	return CryptoS{
		InputData:  nil,
		OutputData: nil,
		Key:        nil,
		IV:         nil,
		Method:     method.AES,
		Mode:       mode.CBC,
	}
}

// WithMethod set method for CryptoS.
func (s *CryptoS) WithMethod(method method.MethodType) *CryptoS {
	s.Method = method
	return s
}

// WithMode set mode for CryptoS.
func (s *CryptoS) WithMode(mode mode.ModeType) *CryptoS {
	s.Mode = mode
	return s
}

// WithPadding set padding for CryptoS.
func (s *CryptoS) WithPadding(padding padding.PaddingType) *CryptoS {
	s.Padding = padding
	return s
}

// WithIV set IV for CryptoS.
func (s *CryptoS) WithIV(data []byte) *CryptoS {
	s.IV = data
	return s
}

// WithKey set Key for CryptoS.
func (s *CryptoS) WithKey(data []byte) *CryptoS {
	s.Key = data
	return s
}

// Reset set all data to default for CryptoS.
func (s *CryptoS) Reset() {
	s.InputData = nil
	s.Errors = nil
	s.Key = nil
	s.IV = nil
	s.Method = method.AES
	s.Mode = mode.CBC
}
