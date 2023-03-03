package symmetric

import (
	"crypto/cipher"
	"errors"
	"fmt"

	"github.com/suyuan32/knife/cryptox/symmetric/mode"
	"github.com/suyuan32/knife/cryptox/symmetric/padding"
)

func (s *CryptoS) Decrypt() *CryptoS {
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

	if len(s.InputData)%block.BlockSize() != 0 {
		s.Errors = errors.Join(s.Errors, errors.New("the data size needs to be an integer multiple of block size"))
		return s
	}

	s.OutputData = make([]byte, len(s.InputData))

	switch s.Mode {
	case mode.CBC:
		cipher.NewCBCDecrypter(block, s.IV).CryptBlocks(s.OutputData, s.InputData)
	case mode.CFB:
		cipher.NewCFBDecrypter(block, s.IV).XORKeyStream(s.OutputData, s.InputData)
	case mode.OFB:
		cipher.NewOFB(block, s.IV).XORKeyStream(s.OutputData, s.InputData)
	case mode.CTR:
		cipher.NewCTR(block, s.IV).XORKeyStream(s.OutputData, s.InputData)
	}

	dePaddingData, err := padding.DePadding(s.OutputData, s.Padding, block.BlockSize())
	if err != nil {
		s.Errors = errors.Join(s.Errors, fmt.Errorf("failed to depad data, error:%s", err))
		return s
	}

	s.OutputData = dePaddingData

	return s
}
