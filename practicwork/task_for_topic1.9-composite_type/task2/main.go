package main

import "fmt"

// Задача 2: Структура "Студент" и метод для вычисления среднего балла
// Описание:
// Создайте структуру Student, которая содержит поля: имя (Name) и список оценок (Grades []float64).
// Реализуйте метод AverageGrade() float64, который возвращает средний балл студента.
// Что нужно сделать:
// Объявить структуру и метод.
// Создать студента с несколькими оценками и вывести его средний балл.

type Student struct {
	Name   string
	Grades []float64
}

func (s Student) AverageGrade() (f float64, r error) {
	if len(s.Grades) == 0 {
		return 0.0, fmt.Errorf("Прогульщик! Оценок нет!")
	}
	sum := 0.0
	for _, val := range s.Grades {
		sum += val
	}
	f = sum / float64(len(s.Grades))
	return
}

func main() {
	student := []Student{
		{Name: "Кто-то", Grades: []float64{2.9, 4.4, 2, 8, 5.0}},
		{Name: "Кто-то"},
		{Name: "Кто-то", Grades: []float64{0}}}
	for _, val := range student {
		ball, err := val.AverageGrade()
		if err != nil {
			fmt.Printf("%q %s %d\n", val.Name, err, len(val.Grades))
		} else {
			fmt.Printf("Cтудент - %q балл - %.3f\n", val.Name, ball)
		}
	}

}
