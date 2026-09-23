package main

import (
	"fmt"
	"strings"
)

func source(ch12 chan string) {
	data := []string{"hello", "my", "friends", "bad", "wonderful weather", "bad weather"}
	for _, v := range data {
		ch12 <- v
	}
	ch12 <- ""
}
func analyze(ch12, ch23 chan string) {
	for {
		message := <-ch12
		if message == "" {
			ch23 <- ""
			return
		}
		if !strings.Contains(message, "bad") {
			ch23 <- message
		}
	}
}
func output(ch23 chan string) {
	for {
		message := <-ch23
		if message == "" {
			return
		}
		fmt.Printf("New message: %s\n", message)
	}
}
func main() {
	ch12 := make(chan string)
	ch23 := make(chan string)
	go source(ch12)
	go analyze(ch12, ch23)
	output(ch23)
}
