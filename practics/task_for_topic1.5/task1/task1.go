package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

// Напишите программу, которая запрашивает у пользователя ввод строки,
// а затем выводит число - количество символов в строке

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	text := sc.Text()
	countRune := utf8.RuneCountInString(text)
	fmt.Printf("%d\n", countRune)
}
