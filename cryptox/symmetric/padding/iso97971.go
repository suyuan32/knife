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

package padding

// PaddingISO97971 add padding at the end of byte slice with zero bytes which are separated by 0x80.
func PaddingISO97971(data []byte, blockSize int) []byte {
	return PaddingZero(append(data, 0x80), blockSize)
}

// DePaddingISO97971 is similar with PKCS7, its block size is 8.
func DePaddingISO97971(data []byte) []byte {
	result := DePaddingZero(data)
	return result[:len(result)-1]
}
