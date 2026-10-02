package main

import "fmt"

// Вам необходимо написать тесты для задач 1 и 4 к теме 8.
// Задача 4 из темы 8
// Создайте функцию sumAll, которая принимает произвольное количество целых чисел и возвращает их сумму.
// Пример использования:
//     fmt.Println(sumAll(1, 2, 3)) // 6
//     fmt.Println(sumAll(10, -2, 4, 7)) // 19

func main() {
	fmt.Println(sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9))
}

func sumAll(line ...int) (r int) {
	for _, v := range line {
		r += v
	}
	return
}
