package main

import (
	"fmt"

	"github.com/gtank/cryptopasta"
)

func main4() {
	// Генерация ключа
	key := cryptopasta.NewEncryptionKey()
	// Текст для шифрования
	plaintext := []byte("Hello, cryptopasta!Hello, cryptopasta!Hello, cryptopasta!Hello, cryptopasta!")
	// Шифрование (автоматически использует GCM)
	ciphertext, err := cryptopasta.Encrypt(plaintext, key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Encrypted: %x\n", ciphertext)
	// Расшифрование
	decrypted, err := cryptopasta.Decrypt(ciphertext, key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decrypted: %s\n", decrypted)
}
