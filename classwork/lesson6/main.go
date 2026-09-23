package main

import "fmt"

func main() {
	animals := [...]string{"Йо", "Тема", "Федя", "Тузик", "Шарик", "Муся"}
	myPets := animals[0:3]
	myPets[2] = "Федор"
	fmt.Println(animals)
}
