package main

import (
	"crypto/aes"
	"fmt"
)

func main3() {
	key := []byte("examplekey123456") // 16, 24 или 32 байта
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	plaintext := []byte("Hello, AES!!!!!!")
	ciphertext := make([]byte, len(plaintext))
	// ECB-режим в Go не встроен, но для примера
	// показываю блоковое шифрование без режима.
	block.Encrypt(ciphertext, plaintext)
	fmt.Printf("Зашифрованный текст: %x\n", ciphertext)
}
