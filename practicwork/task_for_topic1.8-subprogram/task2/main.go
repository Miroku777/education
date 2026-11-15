package main

import "fmt"

//Функция высшего порядка: передача функции как аргумента

//Создайте функцию applyOperation(a, b int, op func(int, int) int) int,
// которая применяет переданную функцию op к числам a и b.
//Создайте несколько функций-операций: сложение, вычитание, умножение.
//В основной программе вызовите applyOperation с разными операциями и выведите результаты.

func main() {
	fmt.Println(applyOperation(12, 5, sum))
	fmt.Println(applyOperation(12, 5, difference))
	fmt.Println(applyOperation(12, 5, multiplicate))
}

func applyOperation(a, b int, op func(int, int) int) int {
	return op(a, b)
}

func sum(a, b int) int {
	return a + b
}

func difference(a, b int) int {
	return a - b
}

func multiplicate(a, b int) int {
	return a * b
}
