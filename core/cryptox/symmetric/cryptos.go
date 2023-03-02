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

// CryptoS is the struct for symmetric-key algorithm.
type CryptoS struct {
	// InputData is the data to be encrypted.
	InputData []byte

	// OutputData is the data encrypted.
	OutputData []byte

	// Key is the secret used for encrypted
	Key []byte

	// An initialization vector (IV) is an arbitrary number that can be used with a secret key for data encryption
	// to foil cyberattacks. This number, also called a nonce (number used once), is employed only one time in
	// any session to prevent unauthorized decryption of the message by a suspicious or malicious actor.
	IV []byte

	// Method is the encryption method such as AES.
	Method Method

	// Mode is the encrypted mode such as ECB.
	Mode Mode

	// Padding is the padding method such as PKCS7.
	Padding Padding

	// Errors is the errors
	Errors error
}

// Method is the encryption method such as AES.
type Method uint8

// Only secure algorithms are provided.
const (
	// The AES Encryption algorithm (also known as the Rijndael algorithm) is a symmetric block cipher algorithm with
	// a block/chunk size of 128 bits. It converts these individual blocks using keys of 128, 192, and 256 bits.
	// Once it encrypts these blocks, it joins them together to form the ciphertext.
	AES Method = 1 + iota

	// The ChaCha20 Encryption Algorithm ChaCha20 is a stream cipher designed by D. J. Bernstein.
	// It is a refinement of the Salsa20 algorithm, and it uses a 256-bit key.
	// ChaCha20 successively calls the ChaCha20 block function, with the same key and nonce, and with successively
	// increasing block counter parameters.
	ChaCha20

	// CAST5 is a symmetric block cipher with a block-size of 8 bytes and a variable key-size of up to 128 bits.
	// Its authors and their employer (Entrust Technologies, a Nortel majority-owned company), made it available worldwide
	// on a royalty-free basis for commercial and non-commercial uses.
	CAST5

	// RC5 is secure symmetric-key block cipher (key size: 128 to 2040 bits; block size: 32, 64 or 128 bits; rounds: 1 ... 255),
	// insecure with short keys (56-bit key successfully brute-forced), was patented until 2015, now royalty-free.
	RC5

	// RC6 is secure symmetric-key block cipher, similar to RC5, but more complicated (key size: 128 to 2040 bits; block size: 32, 64 or 128 bits; rounds: 1 ... 255),
	// was patented until 2017, now royalty-free.
	RC6

	// SM4 is secure symmetric-key block cipher, similar to AES (key size: 128 bits), official standard in China, free for public use.
	SM4

	// Twofish is secure symmetric-key block cipher (key sizes: 128, 192 or 256 bits), royalty-free, not patented.
	Twofish
)

// Mode is the encrypted mode.
type Mode uint8

const (
	// CBC mode, each block of plaintext is XORed with the previous ciphertext block before being encrypted.
	// This way, each ciphertext block depends on all plaintext blocks processed up to that point.
	// To make each message unique, an initialization vector must be used in the first block.
	CBC Mode = 1 + iota

	// CCM mode (counter with cipher block chaining message authentication code; counter with CBC-MAC) is a mode of
	// operation for cryptographic block ciphers. It is an authenticated encryption algorithm designed to provide both
	// authentication and confidentiality. CCM mode is only defined for block ciphers with a block length of 128 bits.
	CCM

	// CFB mode, in its simplest form uses the entire output of the block cipher. In this
	// variation, it is very similar to CBC, makes a block cipher into a self-synchronizing stream cipher.
	CFB

	// OFB (short for output feedback) is an AES block cipher mode similar to the CFB mode.
	// What mainly differs from CFB is that the OFB mode relies on XOR-ing plaintext and ciphertext blocks
	// with expanded versions of the initialization vector.
	OFB
)

// Padding is the padding method such as PKCS7.
type Padding uint8

const (
	ISO97971 Padding = 1 + iota
	No
	PKCS5
	PKCS7
	Zero
)

func NewCryptoS() CryptoS {
	return CryptoS{
		InputData:  nil,
		OutputData: nil,
		Key:        nil,
		IV:         nil,
		Method:     AES,
		Mode:       CBC,
	}
}

// WithMethod set method for CryptoS.
func (s *CryptoS) WithMethod(method Method) *CryptoS {
	s.Method = method
	return s
}

// WithMode set mode for CryptoS.
func (s *CryptoS) WithMode(mode Mode) *CryptoS {
	s.Mode = mode
	return s
}

// WithPadding set padding for CryptoS.
func (s *CryptoS) WithPadding(padding Padding) *CryptoS {
	s.Padding = padding
	return s
}

// WithIV set padding for CryptoS.
func (s *CryptoS) WithIV(padding Padding) *CryptoS {
	s.Padding = padding
	return s
}
