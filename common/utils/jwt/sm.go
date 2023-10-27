package jwt

import (
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"

	"github.com/tjfoc/gmsm/sm4"
)

func Sm4EncryptCfb(plainData, key []byte) (iv, encryptData []byte, err error) {
	block, err := sm4.NewCipher(key)
	if err != nil {
		return nil, nil, err
	}
	encryptData = make([]byte, len(plainData))
	iv = make([]byte, sm4.BlockSize)
	fmt.Println("iv:", iv)
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return nil, nil, err
	}
	mode := cipher.NewCFBEncrypter(block, iv)
	mode.XORKeyStream(encryptData, plainData)
	fmt.Println("iv", iv)
	return
}

func Sm4DecryptCfb(encryptData, key, iv []byte) (plainData []byte, err error) {
	block, err := sm4.NewCipher(key)
	if err != nil {
		return nil, err
	}
	plainData = make([]byte, len(encryptData))
	mode := cipher.NewCFBDecrypter(block, iv)
	mode.XORKeyStream(plainData, encryptData)
	return
}
