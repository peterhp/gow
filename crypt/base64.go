package crypt

import (
	"encoding/base64"
)

func Base64Encrypt(plainText string) string {
	return base64.StdEncoding.EncodeToString([]byte(plainText))
}

func Base64Decrypt(cipherText string) string {
	plain, _ := base64.StdEncoding.DecodeString(cipherText)
	return string(plain)
}
