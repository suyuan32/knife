package symmetric

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"

	"golang.org/x/crypto/cast5"
	"golang.org/x/crypto/tea"
	"golang.org/x/crypto/twofish"
	"golang.org/x/crypto/xtea"

	"github.com/suyuan32/knife/cryptox/symmetric/method"
	"github.com/suyuan32/knife/cryptox/symmetric/method/sm4"
)

// NewCipher returns a cipher block from the cryptos.
func (s *CryptoS) NewCipher() (cipher.Block, error) {
	switch s.Method {
	case method.AES:
		return aes.NewCipher(s.Key)
	case method.SM4:
		return sm4.NewCipher(s.Key)
	case method.CAST5:
		return cast5.NewCipher(s.Key)
	case method.Twofish:
		return twofish.NewCipher(s.Key)
	case method.TEA:
		return tea.NewCipher(s.Key)
	case method.XTEA:
		return xtea.NewCipher(s.Key)
	}
	return nil, errors.New("the method is not supported")
}
