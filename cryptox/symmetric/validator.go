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
