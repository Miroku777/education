package main

import (
	"fmt"
	"strings"
)

//Создайте слайс строк, содержащий названия городов.
// Реализуйте функции для добавления нового города,
// удаления города по имени и поиска города в списке.
// Продемонстрируйте работу этих функций на примере.

func main() {
	city := []string{"Алчевск", "Луганск"}
	addNewCity(&city, "Перевальск", "Стаханов")
	listCity := strings.Join(city, ", ")
	fmt.Printf("Спсиок городов: %s\n", listCity)
	i, ok := getCity(&city, "Перевальск")
	fmt.Printf("Искомый город по индексу: %d - %s\n", i, ok)
	cityName := deleteCity(&city, "Перевальск")
	fmt.Printf("Удалегие города: %s", cityName)
}
func addNewCity(slice *[]string, cityName ...string) bool {
	for _, city := range *slice {
		for _, name := range cityName {
			if city == name {
				return false
			}
		}
	}
	*slice = append(*slice, cityName...)
	return true
}

func getCity(slice *[]string, cityName string) (int, string) {
	for i, city := range *slice {
		if city == cityName {
			return i, "есть"
		}
	}
	return -1, "нет"
}

func deleteCity(slice *[]string, cityName string) string {
	for i, city := range *slice {
		if city == cityName {
			*slice = append((*slice)[:i], (*slice)[i+1:]...)
			return cityName
		}
	}
	return "такого города нет"
}
