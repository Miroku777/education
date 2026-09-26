package main

import (
	"fmt"
	"math"
)

//Задача 2: Полиморфное отображение данных

// Описание: Создайте интерфейс Shape с методом Area() float64. Реализуйте структуры Circle и Rectangle,
// которые реализуют этот интерфейс. Напишите функцию, которая принимает слайс фигур ([]Shape) и выводит площадь каждой фигуры.

// Требования:

//     Интерфейс Shape.
//     Структуры Circle (с радиусом) и Rectangle (с длиной и шириной).
//     Функция для вывода площадей всех фигур в слайсе.

type Shape interface {
	Area() float64
	String() string
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Length float64
	Width  float64
}

func (c Circle) Area() (square float64) {
	square = c.Radius * c.Radius * math.Pi
	return
}
func (r Rectangle) Area() (square float64) {
	square = r.Length * r.Width
	return
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle с окружностью: %f", c.Radius)
}
func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle длина: %f, ширина: %f", r.Length, r.Width)
}

func Result(figureArray []Shape) {
	for _, v := range figureArray {
		fmt.Printf("Фигура: %-45s | её площадь: %.4f\n", v, v.Area())
	}
}

func main() {
	c := Circle{Radius: 2.54}
	r := Rectangle{Length: 44, Width: 20}
	figureArray := []Shape{c, r}
	Result(figureArray)
}
