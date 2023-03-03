package symmetric

import (
	"errors"
	"fmt"

	"github.com/suyuan32/knife/cryptox/symmetric/method"
	"github.com/suyuan32/knife/cryptox/symmetric/mode"
)

// Validate validates the CryptoS and returns error if it does not meet the requirements.
func (s *CryptoS) Validate(blockSize int) error {
	switch s.Mode {
	case mode.CBC, mode.CFB, mode.OFB:
		if len(s.IV) != blockSize {
			return fmt.Errorf("the IV is not the same as block size, IV size: %d, block size: %d", len(s.IV), blockSize)
		}
	default:
		return nil
	}

	if len(s.Key) == 0 {
		return errors.New("the key cannot be empty")
	}

	switch s.Method {
	case method.AES:
		if len(s.Key) != 16 && len(s.Key) != 24 && len(s.Key) != 32 {
			return errors.New("the length of key of AES can only be 16, 24, 32 ")
		}
	}

	return nil
}
