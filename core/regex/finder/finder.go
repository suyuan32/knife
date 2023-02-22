package finder

import (
	"regexp"
)

type Pattern string

const (
	// DigitPure finds all digit in a string.
	DigitPure Pattern = "\\d+\\.?\\d*"

	// ChinesePure finds all Chinese word in a string.
	ChinesePure = "\\p{Han}*"

	// LetterPure finds all letters in a string.
	LetterPure = "[A-Za-z]*"

	// LetterAndDigit finds all letters and digit in a string.
	LetterAndDigit = "[A-Za-z0-9]*"
)

// NewFinder returns a finder to find the strings from a pattern.
//
// Usage:
//
//	result := NewFinder(PureDigit).MatchString("123456")
func NewFinder(pattern Pattern) *regexp.Regexp {
	return regexp.MustCompile(string(pattern))
}

// FindAllNotEmptyString returns all non-empty strings which match the regressions.
func FindAllNotEmptyString(reg *regexp.Regexp, data string) (result []string) {
	tmp := reg.FindAllString(data, -1)
	for _, v := range tmp {
		if v != "" {
			result = append(result, v)
		}
	}
	return result
}
