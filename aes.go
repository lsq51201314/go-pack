package gpk

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func aesEncryptCFB(data []byte, key string) (res []byte, err error) {
	var keyStr string
	if keyStr, err = getMd5([]byte(key)); err != nil {
		return
	}
	var keyDat []byte
	if keyDat, err = hex.DecodeString(keyStr); err != nil {
		return
	}
	var block cipher.Block
	if block, err = aes.NewCipher(keyDat); err != nil {
		return
	}
	res = make([]byte, aes.BlockSize+len(data))
	iv := res[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(res[aes.BlockSize:], data)
	return
}

func aesDecryptCFB(data []byte, key string) (res []byte, err error) {
	var keyStr string
	if keyStr, err = getMd5([]byte(key)); err != nil {
		return
	}
	var keyDat []byte
	if keyDat, err = hex.DecodeString(keyStr); err != nil {
		return
	}
	var block cipher.Block
	if block, err = aes.NewCipher(keyDat); err != nil {
		return
	}
	if len(data) < aes.BlockSize {
		err = fmt.Errorf("ciphertext too short")
		return
	}
	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(data, data)
	return data, nil
}
