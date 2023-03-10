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
)

// ToString output data with string type.
func (s *CryptoS) ToString() (string, error) {
	return string(s.OutputData), s.Errors
}

// ToBytes output data with byte type.
func (s *CryptoS) ToBytes() ([]byte, error) {
	return s.OutputData, s.Errors
}

// ToBase64String output data with base64 string.
func (s *CryptoS) ToBase64String() (string, error) {
	return base64.StdEncoding.EncodeToString(s.OutputData), s.Errors
}

// ToHexString output data with hex string.
func (s *CryptoS) ToHexString() (string, error) {
	return hex.EncodeToString(s.OutputData), s.Errors
}
