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
