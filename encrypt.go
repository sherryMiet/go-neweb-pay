package neweb_pay

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

func Aes256(plaintext string, key string, iv string) string {
	bKey := []byte(key)
	bIV := []byte(iv)
	bPlaintext := PKCS7Padding([]byte(plaintext))
	block, _ := aes.NewCipher(bKey)
	ciphertext := make([]byte, len(bPlaintext))
	mode := cipher.NewCBCEncrypter(block, bIV)
	mode.CryptBlocks(ciphertext, bPlaintext)
	return hex.EncodeToString(ciphertext)
}

func SHA256(str string) string {
	sum := sha256.Sum256([]byte(str))
	checkMac := strings.ToUpper(hex.EncodeToString(sum[:]))
	return checkMac
}

func DecodeAes256(cipherText string, key string, iv string) string {
	bIV := []byte(iv)
	bKey := []byte(key)
	cipherTextDecoded, err := hex.DecodeString(cipherText)
	if err != nil {
		fmt.Errorf(err.Error())
		return ""
	}
	block, err := aes.NewCipher(bKey)
	if err != nil {
		panic(err)
	}

	mode := cipher.NewCBCDecrypter(block, bIV)

	mode.CryptBlocks([]byte(cipherTextDecoded), []byte(cipherTextDecoded))
	result := PKCS7UnPadding(cipherTextDecoded, block.BlockSize())
	return string(result)
}

func PKCS7Padding(ciphertext []byte) []byte {
	padding := aes.BlockSize - len(ciphertext)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(plantText []byte, blockSize int) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	return plantText[:(length - unpadding)]
}
