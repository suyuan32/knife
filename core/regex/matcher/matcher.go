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
	"regexp"
)

type Pattern string

const (
	// DIGIT

	// DigitPure matches string which only contains digits.
	DigitPure Pattern = "^[0-9]*$"

	// DigitPositive matches string which is a positive value.
	DigitPositive Pattern = "^(\\+)?\\d+(\\.\\d+)?$"

	// DigitNegative matches string which is a negative value.
	DigitNegative Pattern = "^\\-\\d+(\\.\\d+)?$"

	// DigitPositiveInteger matches string which is a positive integer.
	DigitPositiveInteger Pattern = "^(\\+)?\\d+$"

	// DigitNegativeInteger matches string which is a negative integer.
	DigitNegativeInteger Pattern = "^\\-\\d+$"

	// DigitFloat matches string which is a float value.
	DigitFloat Pattern = "^(-?\\d+)(\\.\\d+)$"

	// DigitPositiveFloat matches string which is a positive float value.
	DigitPositiveFloat Pattern = "^[1-9]\\d+\\.\\d+|0\\.\\d+[1-9]\\d+$"

	// DigitNegativeFloat matches string which is a negative float value.
	DigitNegativeFloat Pattern = "^-([1-9]\\d+\\.\\d+|0\\.\\d+[1-9]\\d+)$"

	// CHINESE

	// ChinesePure matches string which is a pure Chinese string.
	ChinesePure = "^\\p{Han}*$"

	// ChineseAndDigit string which contains Chinese string and digits.
	ChineseAndDigit = "^[\\p{Han}0-9]*$"

	// ChineseLetterAndDigit string which contains Chinese string, letters and digits.
	ChineseLetterAndDigit = "^[\\p{Han}0-9a-zA-z]*$"

	// ENGLISH

	// LetterAndDigit matches string which contains letters and digits.
	LetterAndDigit = "^[A-Za-z0-9]+$"

	// LetterPure matches string which only contains letters.
	LetterPure = "^[A-Za-z]+$"

	// LetterPureUpper matches string which only contains uppercase letters.
	LetterPureUpper = "^[A-Z]+$"

	// LetterPureLower matches string which only contains lowercase letters.
	LetterPureLower = "^[a-z]+$"

	// LetterDigitAndUnderline matches string which contains letters, digits and underline.
	LetterDigitAndUnderline = "^\\w+$"
)

// NewMatcher returns a matcher to match strings from a pattern.
//
// Usage:
//
//	result := NewMatcher(PureDigit).MatchString("123456")
func NewMatcher(pattern Pattern) *regexp.Regexp {
	return regexp.MustCompile(string(pattern))
}
