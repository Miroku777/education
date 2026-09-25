package main

import "fmt"

//Задача 1: Создание структуры "Книга"
// Описание:
// Создайте структуру Book, которая содержит поля: Title (строка), Author (строка), Year (целое число).
// Добавьте метод GetInfo(), который возвращает строку с информацией о книге в формате: "Title" by Author, Year.
// Что нужно сделать:
// Объявить структуру Book.
// Реализовать метод GetInfo() для этой структуры.
// Создать экземпляр книги и вывести информацию с помощью метода.

type Book struct {
	Title  string
	Author string
	Year   int
}

func (b Book) GetInfo() string {
	return fmt.Sprintf("\"%s\" by %s, %d", b.Title, b.Author, b.Year)
}

func main() {
	book := Book{Title: "Командная строка Linux", Author: "Шоттс Уильям", Year: 2025}
	fmt.Println(book.GetInfo())
}
