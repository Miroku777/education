package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

//Напишите программу, которая читает строку текста и подсчитывает количество вхождений каждого слова.
// Используйте мапу (map[string]int) для хранения результатов. Выведите полученную статистику.

// Hello world!!! Hello world ! ! !
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()
	line, _ := reader.ReadString('\n')
	dict := countString(line)
	var words string
	for key, value := range dict {
		words += key + ", " + strconv.Itoa(value) + "\n"
	}
	writer.WriteString(words)
	writer.WriteByte('\n')
}

func countString(line string) map[string]int {
	words := strings.Fields(line)
	var dict = map[string]int{}
	for _, val := range words {
		dict[val]++
	}
	return dict
}
