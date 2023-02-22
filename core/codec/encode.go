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

package codec

import (
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
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

// HexEncode encodes the string to hex format.
func HexEncode(data string) string {
	return hex.EncodeToString([]byte(data))
}
