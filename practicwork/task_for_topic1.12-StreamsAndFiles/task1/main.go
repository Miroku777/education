package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// Задача 1
// Создайте программу, которая читает лог-файл server.log, подсчитывает количество строк, содержащих слово "error", и выводит это число.

func main() {
	os.Rename("./logfile.log", "./server.log")

	data, err := os.ReadFile("./server.log")
	if err != nil {
		log.Println(err)
		return
	}
	s := string(data)

	count := 0
	for _, v := range strings.Split(s, "\n") {
		if strings.Contains(v, "error") {
			count++
		}
	}
	fmt.Println(count)
}
