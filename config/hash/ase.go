package Hash

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

var (
	// We're using a 32 byte long secret key.
	// This is probably something you generate first
	// then put into and environment variable.
	secretKey string = "N1PCdw3M2B1TfJhoaY2mL736p2vCUc47"
)

func Encrypt(plaintext string) (string, error) {
	byteMsg := []byte(plaintext)
	// aes, err := aes.NewCipher([]byte(secretKey))
	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		panic(err)
	}
	cipherText := make([]byte, aes.BlockSize+len(byteMsg))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], byteMsg)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}
func Decrypt(ciphertext string) (string, error) {
	cipherText, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		panic(err)
	}
	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		panic(err)
	}
	if len(cipherText) < aes.BlockSize {
		panic(err)
	}
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)
	return string(cipherText), nil
}
