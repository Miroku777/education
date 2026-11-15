package main

import (
	"fmt"
	"strings"
)

// Написать программу на ЯП Go. Условие: На вход программе подаются 5 натуральных чисел.
// Задача программы - считать эти числа, отсортировать по убыванию и вывести результат сортировки на экран.
// Также вывести на экран другие параметры считанной последовательности в формате:
// "Самое большое число: {число}", "Самое маленькое число: {число}", "Среднее арифметическое: {число}"
// Пример работы программы:
// выход: 5 2 4 3 1
// выход:
// "Отсортированные элементы: 5 4 3 2 1
// Самое большое число: 5
// Самое маленькое число: 1
// Среднее-арифметическое: 3
// "

func main() {
	sequanceValue, maxValue, minValue, avgValue, error := makeBableSortAndMaxMinAvgElements([]int{5, 4, 3, 2, 1})
	if error != nil {
		fmt.Println("NullPointerException:", error)
	}
	fmt.Printf("Отсортированные элементы: %s\n", sequanceValue)
	fmt.Printf("Самое большое число: %d\n", maxValue)
	fmt.Printf("Самое маленькое число: %d\n", minValue)
	fmt.Printf("Среднее-арифметическое: %d\n", avgValue)
}

func makeBableSortAndMaxMinAvgElements(input []int) (string, int, int, int, error) {
	if len(input) == 0 {
		return "", 0, 0, 0, fmt.Errorf("Length array less then 1")
	}
	avg := 0
	for i := 0; i < len(input)-1; i++ {
		for j := 0; j < len(input)-1-i; j++ {
			if input[j] < input[j+1] {
				temp := input[j+1]
				input[j+1] = input[j]
				input[j] = temp
			}
		}
		avg += input[i]
	}
	avg += input[len(input)-1]
	trimBreakets := strings.Trim(fmt.Sprint(input), "[]")
	return trimBreakets, input[0], input[len(input)-1], avg / len(input), nil
}
