package main

import (
	"log"
	"os"
	"path/filepath"
)

// Задача 2
// Напишите программу, которая принимает список имен файлов в текущей директории,
// объединяет их содержимое и сохраняет результат в новый файл combined.txt (работать только с текстовыми файлами)

func main() {
	f, err := os.Create("./combined.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	suffix, _ := filepath.Glob("./newDirectory/*.txt")
	for _, file := range suffix {
		data, err := os.ReadFile(file)
		if err != nil {
			log.Println(err)
			continue
		}
		_, errr := f.Write(data)
		if errr != nil {
			log.Println(errr)
		}
		f.WriteString("\n")
	}
	os.Remove("./combined.txt")
}
