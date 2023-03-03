package symmetric

import (
	"crypto/cipher"
	"errors"
	"fmt"

	"github.com/suyuan32/knife/cryptox/symmetric/mode"
	"github.com/suyuan32/knife/cryptox/symmetric/padding"
)

func (s *CryptoS) Encrypt() *CryptoS {
	if len(s.InputData) == 0 {
		s.Errors = errors.Join(s.Errors, errors.New("input data cannot be empty"))
		return s
	}

	block, err := s.NewCipher()
	if err != nil {
		s.Errors = errors.Join(s.Errors, fmt.Errorf("failed to create cipher from the data, error:%s", err))
		return s
	}

	err = s.Validate(block.BlockSize())
	if err != nil {
		s.Errors = errors.Join(s.Errors, fmt.Errorf("failed to validate data, error:%s", err))
		return s
	}

	paddingData, err := padding.Padding(s.InputData, s.Padding, block.BlockSize())
	if err != nil {
		s.Errors = errors.Join(s.Errors, fmt.Errorf("failed to pad data, error:%s", err))
		return s
	}

	s.OutputData = make([]byte, len(paddingData))

	switch s.Mode {
	case mode.CBC:
		cipher.NewCBCEncrypter(block, s.IV).CryptBlocks(s.OutputData, paddingData)
	case mode.CFB:
		cipher.NewCFBEncrypter(block, s.IV).XORKeyStream(s.OutputData, paddingData)
	case mode.OFB:
		cipher.NewOFB(block, s.IV).XORKeyStream(s.OutputData, paddingData)
	case mode.CTR:
		cipher.NewCTR(block, s.IV).XORKeyStream(s.OutputData, paddingData)
	default:
		s.Errors = errors.Join(s.Errors, errors.New("the mode is not supported"))
		return s
	}

	return s
}
