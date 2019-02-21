package crypt

import (
	"math/rand"
	"time"
)

var charset = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandString(length int) string {
	data := make([]byte, length)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for index := range data {
		data[index] = charset[r.Intn(len(charset))]
	}

	return string(data)
}
