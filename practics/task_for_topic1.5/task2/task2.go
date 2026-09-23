package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// Напишите программу, которая подсчитывает количество гласных букв
// (а, е, ё, и, о, у, ы, э, ю, я) в введённой пользователем строке.
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	text := sc.Text()
	str := "аеёиоуыэюя"
	countVowel := 0
	for _, v := range text {
		if strings.ContainsRune(str, unicode.ToLower(v)) {
			countVowel++
		}
	}
	fmt.Printf("%d", countVowel)
}
