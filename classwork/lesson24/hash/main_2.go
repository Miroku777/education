package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// GenerateHMAC создает HMAC подпись для сообщения
func GenerateHMAC(key, message []byte) string {
	h := hmac.New(sha256.New, key)
	h.Write(message)
	return hex.EncodeToString(h.Sum(nil))
}

// VerifyHMAC проверяет HMAC подпись
func VerifyHMAC(key, message []byte, signature string) bool {
	// Вычисляем ожидаемую подпись
	h := hmac.New(sha256.New, key)
	h.Write(message)
	expectedSignature := hex.EncodeToString(h.Sum(nil))
	// Сравниваем (используем безопасное сравнение)
	expectedBytes, err1 := hex.DecodeString(expectedSignature)
	receivedBytes, err2 := hex.DecodeString(signature)
	if err1 != nil || err2 != nil {
		return false
	}
	return hmac.Equal(expectedBytes, receivedBytes)
}
func main2() {
	key := []byte("secret-key-123")
	message := []byte("Hello, HMAC World!")
	// Генерация подписи
	signature := GenerateHMAC(key, message)
	fmt.Printf("Message: %s\n", string(message))
	fmt.Printf("HMAC Signature: %s\n\n", signature)
	// Проверка корректной подписи
	isValid := VerifyHMAC(key, message, signature)
	fmt.Printf("Valid signature check: %v\n", isValid)
	// Проверка неверной подписи
	isValid = VerifyHMAC(key, message, "wrong-signature-here")
	fmt.Printf("Invalid signature check: %v\n", isValid)
	// Проверка с измененным сообщением
	isValid = VerifyHMAC(key, []byte("Modified message"), signature)
	fmt.Printf("Modified message check: %v\n", isValid)
}
