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
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/suyuan32/knife/cryptox/symmetric/method"
	"github.com/suyuan32/knife/cryptox/symmetric/mode"
	"github.com/suyuan32/knife/cryptox/symmetric/padding"
)

var data = NewCryptoS()

func TestNewCryptoS(t *testing.T) {
	c := NewCryptoS()
	assert.Equal(t, method.AES, c.Method)
	assert.Equal(t, mode.CBC, c.Mode)
}

func TestCryptoS_Encrypt_Decrypt(t *testing.T) {
	testStr := bytes.Repeat([]byte{'a'}, 120)
	for _, v := range []method.MethodType{method.AES, method.Twofish, method.SM4} {
		for _, m := range []mode.ModeType{mode.CBC, mode.CFB, mode.CTR, mode.OFB} {
			for _, p := range []padding.PaddingType{padding.Zero, padding.ISO97971, padding.PKCS7} {
				result, err := data.InputFromBytes(testStr).
					WithMethod(v).
					WithMode(m).
					WithPadding(p).
					WithIV(bytes.Repeat([]byte{'b'}, 16)).
					WithKey(bytes.Repeat([]byte{'c'}, 16)).
					Encrypt().
					ToString()

				assert.Nil(t, err)

				decryptResult, err := data.InputFromString(result).
					WithMethod(v).
					WithMode(m).
					WithPadding(p).
					WithIV(bytes.Repeat([]byte{'b'}, 16)).
					WithKey(bytes.Repeat([]byte{'c'}, 16)).
					Decrypt().
					ToString()

				assert.Nil(t, err)
				assert.Equal(t, string(testStr), decryptResult)
			}
		}
	}

	for _, v := range []method.MethodType{method.CAST5, method.TEA, method.XTEA} {
		for _, m := range []mode.ModeType{mode.CBC, mode.CFB, mode.CTR, mode.OFB} {
			for _, p := range []padding.PaddingType{padding.Zero, padding.ISO97971, padding.PKCS5, padding.PKCS7} {
				result, err := data.InputFromBytes(testStr).
					WithMethod(v).
					WithMode(m).
					WithPadding(p).
					WithIV(bytes.Repeat([]byte{'b'}, 8)).
					WithKey(bytes.Repeat([]byte{'c'}, 16)).
					Encrypt().
					ToString()

				assert.Nil(t, err)

				decryptResult, err := data.InputFromString(result).
					WithMethod(v).
					WithMode(m).
					WithPadding(p).
					WithIV(bytes.Repeat([]byte{'b'}, 8)).
					WithKey(bytes.Repeat([]byte{'c'}, 16)).
					Decrypt().
					ToString()

				assert.Nil(t, err)
				assert.Equal(t, string(testStr), decryptResult)
			}
		}
	}

	// test encrypt

	data.Reset()
	data.Encrypt()
	assert.NotNil(t, data.Errors)

	data.Reset()
	data.InputFromString("hello")
	data.Encrypt()
	assert.NotNil(t, data.Errors)

	// test decrypt

	data.Reset()
	data.Decrypt()
	assert.NotNil(t, data.Errors)

	data.Reset()
	data.InputFromString("hello")
	data.Decrypt()
	assert.NotNil(t, data.Errors)

}

func TestCryptoS_Input(t *testing.T) {
	data.Reset()
	data.InputFromString("hello")
	assert.Equal(t, "hello", string(data.InputData))

	data.InputFromBase64String("aGVsbG8=")
	assert.Equal(t, "hello", string(data.InputData))

	data.InputFromBytes([]byte{1, 1})
	assert.Equal(t, []byte{1, 1}, data.InputData)

	data.InputFromHexString("68656C6C6F")
	assert.Equal(t, "hello", string(data.InputData))
}

func TestCryptoS_Key(t *testing.T) {
	data.Reset()
	data.KeyFromString("hello")
	assert.Equal(t, "hello", string(data.Key))

	data.KeyFromBase64String("aGVsbG8=")
	assert.Equal(t, "hello", string(data.Key))

	data.KeyFromBytes([]byte{1, 1})
	assert.Equal(t, []byte{1, 1}, data.Key)

	data.KeyFromHexString("68656C6C6F")
	assert.Equal(t, "hello", string(data.Key))
}

func TestCryptoS_IV(t *testing.T) {
	data.Reset()
	data.IVFromString("hello")
	assert.Equal(t, "hello", string(data.IV))

	data.IVFromBase64String("aGVsbG8=")
	assert.Equal(t, "hello", string(data.IV))

	data.IVFromBytes([]byte{1, 1})
	assert.Equal(t, []byte{1, 1}, data.IV)

	data.IVFromHexString("68656C6C6F")
	assert.Equal(t, "hello", string(data.IV))
}

func TestCryptoS_Output(t *testing.T) {
	data.Reset()
	data.OutputData = []byte("hello")

	result, err := data.ToString()
	assert.Nil(t, err)
	assert.Equal(t, "hello", result)

	result, err = data.ToBase64String()
	assert.Nil(t, err)
	assert.Equal(t, "aGVsbG8=", result)

	byteResult, err := data.ToBytes()
	assert.Equal(t, data.OutputData, byteResult)

	result, err = data.ToHexString()
	assert.Nil(t, err)
	assert.Equal(t, "68656c6c6f", result)
}

func TestCryptoS_Validate(t *testing.T) {
	data.Reset()
	data.WithIV([]byte{1, 1, 1})
	err := data.Validate(10)
	assert.NotNil(t, err)

	err = data.Validate(3)
	assert.NotNil(t, err)

	data.WithKey([]byte{1, 1, 1})
	err = data.Validate(3)
	assert.NotNil(t, err)
}
