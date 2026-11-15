package main

import (
	"fmt"
)

// Вам необходимо написать тесты для задач 1 и 4 к теме 8.
// Задача 1 из темы 8
// Реализуйте функцию divide(a, b float64) (float64, error), которая делит a на b. Если b равно нулю, возвращайте ошибку.
// В основной программе вызовите эту функцию и обработайте возможную ошибку.

func main() {
	fmt.Println(Divide(54, -32))
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Деление на %f", b)
	}
	return a / b, nil
}
