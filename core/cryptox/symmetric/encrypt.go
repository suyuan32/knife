package symmetric

import (
	"errors"
)

func (s *CryptoS) Encrypt() *CryptoS {
	if len(s.InputData) == 0 {
		s.Errors = errors.Join(s.Errors, errors.New("input data cannot be empty"))
		return s
	}

}
