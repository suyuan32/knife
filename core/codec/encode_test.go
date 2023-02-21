package codec

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBase32Encode(t *testing.T) {
	assert.Equal(t, "JBSWY3DPFQQHO33SNRSCC===", Base32Encode("Hello, world!"))
}

func TestBase64Encode(t *testing.T) {
	assert.Equal(t, "SGVsbG8sIOS4lueVjA==", Base64Encode("Hello, 世界"))
}
