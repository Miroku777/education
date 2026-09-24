package main

import (
	"fmt"
)

// Реализуйте функцию divide(a, b float64) (float64, error), которая делит a на b. Если b равно нулю, возвращайте ошибку.
// В основной программе вызовите эту функцию и обработайте возможную ошибку.

func main() {
	value, err := divide(46, 15.5)
	if err != nil {
		fmt.Printf("ArifmeticException: %s", err)
		return
	}
	fmt.Printf("%f", value)
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Деление на ноль")
	}
	return a / b, nil
}
