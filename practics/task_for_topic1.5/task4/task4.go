package main

import (
	"bufio"
	"fmt"
	"os"
)

// Напишите программу, которая запрашивает у пользователя ввод строки-формулы,
// а выводит сообщение о правильности написания круглых скобок, например:

// Пример 1
// строка на вход: (1+1)*(2+2)
// вывод: Скобки расставлены верно, 2 открывающиеся, 2 закрывающиеся

// Пример 2
// Строка на вход: ((1+1) + (2+2) ))
// вывод: Скобки расставлены неправильно, 3 открывающиеся, 4 закрывающиеся

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	text := sc.Text()
	countBrecketOpen := 0
	countBrecketClose := 0
	sequanceBrecket := 0
	valid := true
	for _, v := range text {
		if v == '(' {
			sequanceBrecket++
			countBrecketOpen++
		} else if v == ')' {
			sequanceBrecket--
			countBrecketClose++
			if sequanceBrecket < 0 {
				valid = false
			}
		}
	}
	if valid && sequanceBrecket == 0 {
		fmt.Printf("вывод: Скобки расставлены верно, %d открывающиеся, %d закрывающиеся", countBrecketOpen, countBrecketClose)
	} else {
		fmt.Printf("вывод: Скобки расставлены неправильно, %d открывающиеся, %d закрывающиеся", countBrecketOpen, countBrecketClose)
	}
}
