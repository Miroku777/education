package main

import (
	"crypto/sha512"
	"fmt"
)

// хеширование - это не шифрование, а генерация индивидуальной метки(22 символа) для набора данных
// md5, sha-1 - NOT USED!
// SHA256, SHA-512 - USED!

// Шифрование самое надёжное AES

func main1() {
	message := []byte("12345671")
	hash := sha512.Sum512(message)
	fmt.Printf("%x\n", hash)
}
