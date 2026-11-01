package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
)

// RSA шифрование (асимметричное), для https. Шифрование публичны ключом, расшифровка приватным

func main() {
	// Генерация ключей
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Println(err)
		return
	}
	publicKey := &privateKey.PublicKey
	// Сообщение для шифрования
	message := []byte("Hello, RSA Encryption!")
	// Шифрование с помощью публичного ключа
	ciphertext, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		publicKey,
		message,
		nil,
	)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("Original: %s\n", message)
	fmt.Printf("Encrypted (base64): %s\n", base64.StdEncoding.EncodeToString(ciphertext))
	// Расшифрование с помощью приватного ключа
	plaintext, err := rsa.DecryptOAEP(
		sha256.New(),
		rand.Reader,
		privateKey,
		ciphertext,
		nil,
	)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("Decrypted: %s\n", plaintext)
}
