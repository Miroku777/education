package main

// Task0:
// Задача: написать конольное приложение на яп GO с следующими требованиями:
// 	1) Пользователь-кассир он вводит в консоль по очереди строки (всего 3 строки):
// 		$enter command:молоко 100(цена) 1(количество)
// 	2) введенные данные сохраняются в переменные
// 	3) после введения третьего товара - консоль спрашивает: есть ли карта постоянного покупателя?
// 		если да- на все скидка 5%
// 	4) резульатом является красиво оформленный чек, как в примере
// 	=======================================
// 				ОАО Golang
// 	=======================================
// 		Молоко
// 		  1 х 100
// 		  Скидка: 5%
// 		  Сумма: 95р

// 		Хлеб
// 		  2 х 50
// 		  Скидка: 5%
// 		  Сумма: 95р
// 	----------------------------------------
// 	Итого без скидки: 200
// 	Скидка: 5%
// 	Итог: 190р
// 	-----------------------------------------
// 			    Спасибо за покупку
// 	-----------------------------------------
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var product []string
	for i := 0; i < 3; i++ {
		fmt.Printf("%s", "$enter command:")
		sc.Scan()
		product = append(product, sc.Text())
	}

	fmt.Printf("есть ли карта постоянного покупателя? ")
	sc.Scan()
	card := sc.Text() == "да"

	chek(product, card)
}

func chek(products []string, card bool) {

	var sumPrice int
	var discont int
	if card {
		discont = 5
	}

	fmt.Print(`
				=======================================
						ОАО Golang
				=======================================`)
	for _, val := range products {
		prod := strings.Split(val, " ")
		if len(prod) < 3 {
			continue
		}
		name := prod[0]
		price, _ := strconv.Atoi(prod[1])
		count, _ := strconv.Atoi(prod[2])

		fmt.Printf(`
				%s
				  %d х %d
				  Скидка: %d%%
				  Сумма: %dр
				  `, name, count, price, discont, count*price)

		sumPrice += price * count
	}

	fmt.Printf(`----------------------------------------
				Итого без скидки: %d
				Скидка: %d%%
				Итог: %dр
				-----------------------------------------
			 			Спасибо за покупку
				-----------------------------------------`, sumPrice, discont, sumPrice-(sumPrice*discont/100))
}
