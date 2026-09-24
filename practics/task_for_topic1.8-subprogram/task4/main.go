package main

import "fmt"

// Создайте функцию sumAll,
// которая принимает произвольное количество целых чисел и возвращает их сумму.
// Пример использования:
//    fmt.Println(sumAll(1, 2, 3)) // 6
//    fmt.Println(sumAll(10, -2, 4, 7)) // 19

func main() {
	fmt.Println(sumAll())             // 0
	fmt.Println(sumAll(1, 2, 3))      // 6
	fmt.Println(sumAll(10, -2, 4, 7)) // 19
}

func sumAll(sequanceInt ...int) (ret int) {
	for _, val := range sequanceInt {
		ret += val
	}
	return
}
