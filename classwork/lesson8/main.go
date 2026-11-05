package main

import (
	"fmt"
)

type cat struct {
	name  string
	voice string
}
type dog struct {
	name  string
	voice string
}

func (c cat) speak() {
	fmt.Printf("Кот %s говорит %s\n", c.name, c.voice)
}
func (d dog) speak() {
	fmt.Printf("Кот %s говорит %s", d.name, d.voice)
}

func main() {
	cat := cat{name: "Барсик", voice: "Мяу"}
	dog := dog{name: "Шарик", voice: "Гав"}
	cat.speak()
	dog.speak()
}
