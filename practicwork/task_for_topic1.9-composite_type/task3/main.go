package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Задача 3: Структура "Банковский счет" и методы для работы с балансом
// Описание:
// Создайте структуру BankAccount с полями: Owner (владелец счета) и Balance (текущий баланс).
// Реализуйте методы:
// Deposit(amount float64), увеличивающий баланс.
// Withdraw(amount float64), уменьшающий баланс, если хватает средств, иначе выводит сообщение о недостатке средств.
// Что нужно сделать:
// Объявить структуру и методы.
// Создать счет, пополнить его, попытаться снять деньги и вывести итоговый баланс.

type BankAccount struct {
	Owner   string
	Balance float64
}

func (b *BankAccount) Deposit(amount float64) { //увеличивающий баланс
	b.Balance += amount
}
func (b *BankAccount) Withdraw(amount float64) error { //уменьшающий баланс, если хватает средств, иначе выводит сообщение о недостатке средств
	if b.Balance <= amount {
		return fmt.Errorf("Нехватка средств! %f", b.Balance)
	}
	b.Balance -= amount
	return nil
}

func main() {
	//Name + 3.2
	//Name + 55.7
	//Name - 60
	fmt.Println("Формат ввода: Имя +/- 0.0")
	sc := bufio.NewScanner(os.Stdin)
	balanceAccount := BankAccount{Owner: "Иван", Balance: 1000}
	for {
		bol := sc.Scan()
		if !bol {
			return
		}
		line := sc.Text()
		slice := strings.Fields(line)

		balance, _ := strconv.ParseFloat(slice[2], 64)
		if slice[1] == "+" {
			balanceAccount.Deposit(balance)
			fmt.Printf("%s %s %f\n", slice[0], slice[1], balanceAccount.Balance)
		} else if slice[1] == "-" {
			err := balanceAccount.Withdraw(balance)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("%s %s %f\n", slice[0], slice[1], balanceAccount.Balance)
		}
	}
}
