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
	"fmt"
)

// Base32Decode decodes base32 string.
func Base32Decode(data string) (string, error) {
	result, err := base32.StdEncoding.DecodeString(data)
	if err != nil {
		return "", fmt.Errorf("failed to decode base32 string, err: %v", err)
	}
	return string(result), nil
}

// Base64Decode decodes base64 string.
func Base64Decode(data string) (string, error) {
	decodeString, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", fmt.Errorf("failed to decode base64 string, err: %v", err)
	}
	return string(decodeString), nil
}

// HexDecode decodes hex string.
func HexDecode(data string) (string, error) {
	decodeString, err := hex.DecodeString(data)
	if err != nil {
		return "", fmt.Errorf("failed to decode hex string, err: %v", err)
	}
	return string(decodeString), nil
}
