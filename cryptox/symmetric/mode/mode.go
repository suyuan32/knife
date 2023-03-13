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

package mode

// ModeType is the encrypted mode.
type ModeType uint8

const (
	// CBC mode, each block of plaintext is XORed with the previous ciphertext block before being encrypted.
	// This way, each ciphertext block depends on all plaintext blocks processed up to that point.
	// To make each message unique, an initialization vector must be used in the first block.
	CBC ModeType = 1 + iota

	// CFB mode, in its simplest form uses the entire output of the block cipher. In this
	// variation, it is very similar to CBC, makes a block cipher into a self-synchronizing stream cipher.
	CFB

	// OFB (short for output feedback) is an AES block cipher mode similar to the CFB mode.
	// What mainly differs from CFB is that the OFB mode relies on XOR-ing plaintext and ciphertext blocks
	// with expanded versions of the initialization vector.
	OFB

	// CTR is a Stream which encrypts/decrypts using the given Block in
	// counter mode. The length of iv must be the same as the Block's block size.
	CTR
)
