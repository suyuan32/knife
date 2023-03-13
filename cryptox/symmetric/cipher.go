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
	"crypto/aes"
	"crypto/cipher"
	"errors"

	"golang.org/x/crypto/cast5"
	"golang.org/x/crypto/tea"
	"golang.org/x/crypto/twofish"
	"golang.org/x/crypto/xtea"

	"github.com/suyuan32/knife/cryptox/symmetric/method"
	"github.com/suyuan32/knife/cryptox/symmetric/method/sm4"
)

// NewCipher returns a cipher block from the cryptos.
func (s *CryptoS) NewCipher() (cipher.Block, error) {
	switch s.Method {
	case method.AES:
		return aes.NewCipher(s.Key)
	case method.SM4:
		return sm4.NewCipher(s.Key)
	case method.CAST5:
		return cast5.NewCipher(s.Key)
	case method.Twofish:
		return twofish.NewCipher(s.Key)
	case method.TEA:
		return tea.NewCipher(s.Key)
	case method.XTEA:
		return xtea.NewCipher(s.Key)
	}
	return nil, errors.New("the method is not supported")
}
