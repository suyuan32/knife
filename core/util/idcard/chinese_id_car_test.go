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

package idcard

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChineseID_IsValidCard(t *testing.T) {
	card := &ChineseID{Id: "420101198808240047"}
	assert.Equal(t, true, card.IsValidCard())

	data := []struct {
		Id   string
		Want bool
	}{
		{
			"420101198808240048",
			false,
		},
		{
			"420101198808420048",
			false,
		},
		{
			"420101198808310048",
			false,
		},
		{
			"000101198808310048",
			false,
		},
		{
			"110101199003076931",
			true,
		},
		{
			"11010119900307X931",
			false,
		},
		{
			"",
			false,
		},
		{
			"A158774414",
			true,
		},
		{
			"A158774415",
			false,
		},
		{
			"C668668(E)",
			true,
		},
		{
			"CAA8668(E)",
			false,
		},
		{
			"1000248(3)",
			true,
		},
		{
			"123456",
			false,
		},
	}

	for _, v := range data {
		card.Id = v.Id
		assert.Equal(t, v.Want, card.IsValidCard())
	}

	for i := 0; i < 9; i++ {
		card.Id = fmt.Sprintf("4201011%d8808240047", i)
		assert.Equal(t, false, card.IsValidCard())
	}

	for key, _ := range ProvinceCode {
		if key == "42" {
			continue
		}
		card.Id = fmt.Sprintf("%s010119880824004A", key)
		assert.Equal(t, false, card.IsValidCard())
	}
}
