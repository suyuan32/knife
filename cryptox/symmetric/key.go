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
