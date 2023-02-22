package finder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDigitPure(t *testing.T) {
	finder := NewFinder(DigitPure)
	result := finder.FindAllString("123 456 10.1", -1)
	assert.Equal(t, "123", result[0])
	assert.Equal(t, "456", result[1])
	assert.Equal(t, "10.1", result[2])
}

func TestChinesePure(t *testing.T) {
	finder := NewFinder(ChinesePure)
	result := FindAllNotEmptyString(finder, "Hi 您好 10000 号！")
	assert.Equal(t, "您好", result[0])
	assert.Equal(t, "号", result[1])
}

func TestLetterPure(t *testing.T) {
	finder := NewFinder(LetterPure)
	result := FindAllNotEmptyString(finder, "Hi 您好 10000 号！ 我是 Jack ")
	assert.Equal(t, "Hi", result[0])
	assert.Equal(t, "Jack", result[1])
}

func TestLetterAndDigit(t *testing.T) {
	finder := NewFinder(LetterAndDigit)
	result := FindAllNotEmptyString(finder, "Hi 您好 10000 号！ 我是 Jack ")
	assert.Equal(t, "Hi", result[0])
	assert.Equal(t, "10000", result[1])
	assert.Equal(t, "Jack", result[2])
}
