package main

import (
	"bufio"
	"os"
	"unicode"
)

// Создайте функцию capitalizeWords(s string) string,
// которая преобразует каждое слово в строке так, чтобы первая буква была заглавной,
// а остальные — строчными. Например: "привет мир" → "Привет Мир".

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()
	line, _ := reader.ReadString('\n')
	writer.WriteString(capitalizeWords(line))
}

func capitalizeWords(s string) string {
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	for i := 1; i < len(runes); i++ {
		runes[i] = unicode.ToLower(runes[i])
		if runes[i-1] == ' ' && runes[i] != ' ' {
			runes[i] = unicode.ToUpper(runes[i])
		}
	}
	return string(runes)
}
