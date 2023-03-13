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

package method

// MethodType is the encryption method such as AES.
type MethodType uint8

// Only secure algorithms are provided.
const (
	// The AES Encryption algorithm (also known as the Rijndael algorithm) is a symmetric block cipher algorithm with
	// a block/chunk size of 128 bits. It converts these individual blocks using keys of 128, 192, and 256 bits.
	// Once it encrypts these blocks, it joins them together to form the ciphertext.
	AES MethodType = 1 + iota

	// The ChaCha20 Encryption Algorithm ChaCha20 is a stream cipher designed by D. J. Bernstein.
	// It is a refinement of the Salsa20 algorithm, and it uses a 256-bit key.
	// ChaCha20 successively calls the ChaCha20 block function, with the same key and nonce, and with successively
	// increasing block counter parameters.
	//ChaCha20

	// CAST5 is a symmetric block cipher with a block-size of 8 bytes and a variable key-size of up to 128 bits.
	// Its authors and their employer (Entrust Technologies, a Nortel majority-owned company), made it available worldwide
	// on a royalty-free basis for commercial and non-commercial uses.
	CAST5

	// SM4 is secure symmetric-key block cipher, similar to AES (key size: 128 bits), official standard in China, free for public use.
	SM4

	// Twofish is secure symmetric-key block cipher (key sizes: 128, 192 or 256 bits), royalty-free, not patented.
	Twofish

	// TEA (Tiny Encryption Algorithm) is a block cipher notable for its simplicity of description and
	// implementation, typically a few lines of code. TEA operates on two 32-bit unsigned integers
	// (could be derived from a 64-bit data block) and uses a 128-bit key.
	TEA

	// XTEA is a secure encryption algorithm, though not as secure as RSA or others,
	// that uses a 128bit key and requires very little processing power.
	// Like TEA, XTEA is a 64-bit block Feistel cipher with a 128-bit key and a suggested 64 rounds.
	// Several differences from TEA are apparent, including a somewhat more complex key-schedule and a
	// rearrangement of the shifts, XORs, and additions.
	XTEA
)
