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

package matcher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDigitPure(t *testing.T) {
	matcher := NewMatcher(DigitPure)
	assert.Equal(t, true, matcher.MatchString("123456"))
	assert.Equal(t, false, matcher.MatchString("123456sad"))
}

func TestDigitPositive(t *testing.T) {
	matcher := NewMatcher(DigitPositive)
	assert.Equal(t, true, matcher.MatchString("10"))
	assert.Equal(t, true, matcher.MatchString("10.1"))
	assert.Equal(t, false, matcher.MatchString("-10"))
	assert.Equal(t, false, matcher.MatchString("-10.1"))
	assert.Equal(t, false, matcher.MatchString("10x"))
}

func TestDigitNegative(t *testing.T) {
	matcher := NewMatcher(DigitNegative)
	assert.Equal(t, false, matcher.MatchString("10"))
	assert.Equal(t, false, matcher.MatchString("10.1"))
	assert.Equal(t, true, matcher.MatchString("-10"))
	assert.Equal(t, true, matcher.MatchString("-10.1"))
	assert.Equal(t, false, matcher.MatchString("10x"))
}

func TestDigitPositiveInteger(t *testing.T) {
	matcher := NewMatcher(DigitPositiveInteger)
	assert.Equal(t, true, matcher.MatchString("10"))
	assert.Equal(t, false, matcher.MatchString("10.1"))
	assert.Equal(t, false, matcher.MatchString("-10"))
	assert.Equal(t, false, matcher.MatchString("-10.1"))
	assert.Equal(t, false, matcher.MatchString("10x"))
}

func TestDigitNegativeInteger(t *testing.T) {
	matcher := NewMatcher(DigitNegativeInteger)
	assert.Equal(t, false, matcher.MatchString("10"))
	assert.Equal(t, false, matcher.MatchString("10.1"))
	assert.Equal(t, true, matcher.MatchString("-10"))
	assert.Equal(t, false, matcher.MatchString("-10.1"))
	assert.Equal(t, false, matcher.MatchString("10x"))
}

func TestDigitFloat(t *testing.T) {
	matcher := NewMatcher(DigitFloat)
	assert.Equal(t, false, matcher.MatchString("10"))
	assert.Equal(t, true, matcher.MatchString("10.1"))
	assert.Equal(t, false, matcher.MatchString("-10"))
	assert.Equal(t, true, matcher.MatchString("-10.1"))
	assert.Equal(t, false, matcher.MatchString("10x"))
}

func TestDigitPositiveFloat(t *testing.T) {
	matcher := NewMatcher(DigitPositiveFloat)
	assert.Equal(t, false, matcher.MatchString("10"))
	assert.Equal(t, true, matcher.MatchString("10.1"))
	assert.Equal(t, false, matcher.MatchString("-10"))
	assert.Equal(t, false, matcher.MatchString("-10.1"))
	assert.Equal(t, false, matcher.MatchString("10x"))
}

func TestDigitNegativeFloat(t *testing.T) {
	matcher := NewMatcher(DigitNegativeFloat)
	assert.Equal(t, false, matcher.MatchString("10"))
	assert.Equal(t, false, matcher.MatchString("10.1"))
	assert.Equal(t, false, matcher.MatchString("-10"))
	assert.Equal(t, true, matcher.MatchString("-10.1"))
	assert.Equal(t, false, matcher.MatchString("10x"))
}

func TestChinesePure(t *testing.T) {
	matcher := NewMatcher(ChinesePure)
	assert.Equal(t, true, matcher.MatchString("你好"))
	assert.Equal(t, false, matcher.MatchString("Hi 你好"))
}

func TestChineseDigit(t *testing.T) {
	matcher := NewMatcher(ChineseAndDigit)
	assert.Equal(t, true, matcher.MatchString("你好"))
	assert.Equal(t, true, matcher.MatchString("你好1"))
	assert.Equal(t, false, matcher.MatchString("你好a"))
}

func TestChineseLetterAndDigit(t *testing.T) {
	matcher := NewMatcher(ChineseLetterAndDigit)
	assert.Equal(t, true, matcher.MatchString("你好"))
	assert.Equal(t, true, matcher.MatchString("你好1"))
	assert.Equal(t, true, matcher.MatchString("你好a"))
	assert.Equal(t, true, matcher.MatchString("你好a1"))
	assert.Equal(t, false, matcher.MatchString("你好a."))
}

func TestLetterAndDigit(t *testing.T) {
	matcher := NewMatcher(LetterAndDigit)
	assert.Equal(t, true, matcher.MatchString("AB10C"))
	assert.Equal(t, false, matcher.MatchString("AB10C."))
	assert.Equal(t, false, matcher.MatchString("AB10C你好"))
}

func TestLetterPure(t *testing.T) {
	matcher := NewMatcher(LetterPure)
	assert.Equal(t, true, matcher.MatchString("ABC"))
	assert.Equal(t, false, matcher.MatchString("AB10C."))
	assert.Equal(t, false, matcher.MatchString("AB10C"))
	assert.Equal(t, false, matcher.MatchString("AB10C你好"))
}

func TestLetterPureUpper(t *testing.T) {
	matcher := NewMatcher(LetterPureUpper)
	assert.Equal(t, true, matcher.MatchString("ABC"))
	assert.Equal(t, false, matcher.MatchString("AB10C."))
	assert.Equal(t, false, matcher.MatchString("AB10C"))
	assert.Equal(t, false, matcher.MatchString("AB10C你好"))
	assert.Equal(t, false, matcher.MatchString("aBc"))
}

func TestLetterPureLower(t *testing.T) {
	matcher := NewMatcher(LetterPureLower)
	assert.Equal(t, true, matcher.MatchString("abc"))
	assert.Equal(t, false, matcher.MatchString("ABC"))
	assert.Equal(t, false, matcher.MatchString("AB10C."))
	assert.Equal(t, false, matcher.MatchString("AB10C"))
	assert.Equal(t, false, matcher.MatchString("AB10C你好"))
	assert.Equal(t, false, matcher.MatchString("aBc"))
}

func TestLetterDigitAndUnderline(t *testing.T) {
	matcher := NewMatcher(LetterDigitAndUnderline)
	assert.Equal(t, true, matcher.MatchString("abc"))
	assert.Equal(t, true, matcher.MatchString("ABC"))
	assert.Equal(t, true, matcher.MatchString("ABC_abc"))
	assert.Equal(t, false, matcher.MatchString("ABC."))
}
