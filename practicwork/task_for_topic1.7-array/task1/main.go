package main

import (
	"fmt"
	"math/rand"
	"sort"
)

//Напишите программу, которая создает массив из 10 целых чисел,
// заполняет его случайными значениями от 1 до 100.
// Затем скопируйте этот массив в слайс и отсортируйте его по возрастанию.
// Выведите исходный массив и отсортированный слайс.

func main() {
	array := [10]int{}
	for i := 0; i < len(array); i++ {
		generate := rand.Intn(100) + 1
		array[i] = generate
	}
	slice := make([]int, len(array))
	copy(slice, array[:])
	sort.Ints(slice)

	fmt.Println(array)
	fmt.Println(slice)
}
