package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"bytes"
	"errors"
	"github.com/Shopify/go-lua"
	"io"
)

func Open(l *lua.State) {
	aesOpen := func(l *lua.State) int {
		lua.NewLibrary(l, aesLibrary)
		return 1
	}
	lua.Require(l, "goluago/crypto/aes", aesOpen, false)
	l.Pop(1)
}

var aesLibrary = []lua.RegistryFunction{
	{"encryptCBC", encryptCBC},
	{"decryptCBC", decryptCBC},
}

func encryptCBC(l *lua.State) int {
	key := []byte(lua.CheckString(l, 1))
	plaintext := []byte(lua.CheckString(l, 2))

	block, err := aes.NewCipher(key)
	if err != nil {
		lua.Errorf(l, err.Error())
		panic("unreachable")
	}

	content := PKCS5Padding(plaintext, aes.BlockSize, len(plaintext))

	ciphertext := make([]byte, aes.BlockSize + len(content))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		lua.Errorf(l, err.Error())
		panic("unreachable")
	}

	encrypter := cipher.NewCBCEncrypter(block, iv)
	encrypter.CryptBlocks(ciphertext[aes.BlockSize:], content)

	l.PushString(string(ciphertext))

	return 1
}

func decryptCBC(l *lua.State) int {
	key := []byte(lua.CheckString(l, 1))
	ciphertext := []byte(lua.CheckString(l, 2))

	block, err := aes.NewCipher(key)
	if err != nil {
		lua.Errorf(l, err.Error())
		panic("unreachable")
	}

	if len(ciphertext) < 2*aes.BlockSize || len(ciphertext)%aes.BlockSize != 0 {
		lua.Errorf(l, "invalid ciphertext length")
		panic("unreachable")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	decrypter := cipher.NewCBCDecrypter(block, iv)
	decrypter.CryptBlocks(ciphertext, ciphertext)

	plaintext, err := PKCS7UnPadding(ciphertext, aes.BlockSize)
	if err != nil {
		lua.Errorf(l, err.Error())
		panic("unreachable")
	}

	l.PushString(string(plaintext))

	return 1
}

func PKCS5Padding(ciphertext []byte, blockSize int, after int) []byte {
	block := (blockSize - len(ciphertext) % blockSize)
	padding := bytes.Repeat([]byte{ byte(block) }, block)
	return append(ciphertext, padding...)
}

// PKCS5UnPadding strips padding without validating it, so malformed input
// (such as the output of decrypting with the wrong key) returns garbage or
// panics.
//
// Deprecated: Use PKCS7UnPadding, which returns an error for invalid padding.
func PKCS5UnPadding(src []byte) []byte {
	src_length := len(src)
	padding_length := int(src[src_length-1])
	return src[:(src_length - padding_length)]
}

// PKCS7UnPadding rejects malformed padding instead of slicing on whatever the
// last byte says, which is what decrypting with the wrong key produces.
func PKCS7UnPadding(src []byte, blockSize int) ([]byte, error) {
	srcLength := len(src)
	if srcLength == 0 {
		return nil, errors.New("invalid padding")
	}
	paddingLength := int(src[srcLength-1])
	if paddingLength == 0 || paddingLength > blockSize || paddingLength > srcLength {
		return nil, errors.New("invalid padding")
	}
	for _, b := range src[srcLength-paddingLength:] {
		if int(b) != paddingLength {
			return nil, errors.New("invalid padding")
		}
	}
	return src[:srcLength-paddingLength], nil
}
