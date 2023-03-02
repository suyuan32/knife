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
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var data = NewCryptoS()

func TestNewCryptoS(t *testing.T) {
	c := NewCryptoS()
	assert.Equal(t, AES, c.Method)
	assert.Equal(t, CBC, c.Mode)
}

func TestCryptoS_AES_CBC(t *testing.T) {
	result := data.WithMethod(AES).WithMode(CBC).WithPadding(Zero).Encrypt().ToString()
	fmt.Println(result)
}
