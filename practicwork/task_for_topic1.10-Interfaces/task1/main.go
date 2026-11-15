package main

import "fmt"

// Задача 1: Реализация системы оплаты с использованием интерфейсов

// Описание: Создайте интерфейс PaymentProcessor с методом Process(amount float64) string,
// который возвращает строку с результатом обработки платежа.
// Реализуйте два типа платежных систем: CreditCard и CryptoWallet, каждый из которых реализует интерфейс PaymentProcessor.
// В функции main создайте список платежных систем и вызовите метод Process для каждого, выводя результат на экран.
// Требования:

//     Интерфейс PaymentProcessor.
//     Структуры CreditCard и CryptoWallet, реализующие интерфейс.
//     В main создайте массив/слайс этих структур и вызовите их методы.
//     P.S. Подумайте, какие поля могут быть у каждой структуры

type PaymentProcessor interface {
	Process(amount float64) string
}

type CreditCard struct {
	NumberCard string
	Holder     string
	Currency   string
}

type CryptoWallet struct {
	AddresWallet string
	Holder       string
	Currency     string
}

func (cc CreditCard) Process(amount float64) string {
	return fmt.Sprintf(cc.NumberCard+" "+cc.Holder+" "+cc.Currency+" Операция CreditCard: %.2f", amount)
}

func (cw CryptoWallet) Process(amount float64) string {
	return fmt.Sprintf(cw.AddresWallet+" "+cw.Holder+" "+cw.Currency+" Операция CryptoWallet: %.2f", amount)
}

func Result(processor PaymentProcessor, amount float64) string {
	return processor.Process(amount)
}

func main() {
	creditCard := CreditCard{NumberCard: "123", Holder: "NameAvtor", Currency: "RUB"}
	cryptoWallet := CryptoWallet{AddresWallet: "555", Holder: "AvtorName", Currency: "USD"}
	pp := []PaymentProcessor{creditCard, cryptoWallet}
	for _, v := range pp {
		fmt.Println(Result(v, 44.2))
	}
}
