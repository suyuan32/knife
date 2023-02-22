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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBase32Decode(t *testing.T) {
	decodeData, err := Base32Decode("ONXW2ZJAMRQXIYJAO5UXI2BAAAQGC3TEEDX3XPY=")
	assert.Nil(t, err)
	assert.Equal(t, "some data with \x00 and \ufeff", decodeData)

	decodeData, err = Base32Decode("x")
	assert.NotNil(t, err)
}

func TestBase64Decode(t *testing.T) {
	decodeData, err := Base64Decode("SGVsbG8sIOS4lueVjA==")
	assert.Nil(t, err)
	assert.Equal(t, "Hello, 世界", decodeData)

	decodeData, err = Base64Decode("x")
	assert.NotNil(t, err)
}

func TestHexDecode(t *testing.T) {
	decodeData, err := HexDecode("48656c6c6f20476f7068657221")
	assert.Nil(t, err)
	assert.Equal(t, "Hello Gopher!", decodeData)

	decodeData, err = HexDecode(".")
	assert.NotNil(t, err)
}
