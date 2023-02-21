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
