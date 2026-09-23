package main

import (
	"fmt"
	"time"
)

func SleepGopher(i int, ch chan int) {
	time.Sleep(3 * time.Second)
	ch <- i
}
func main() {
	ch := make(chan int)
	for i := 0; i < 5; i++ {
		go SleepGopher(i, ch)
	}
	for i := 0; i < 6; i++ {
		id := <-ch
		fmt.Printf("#%d finished\n", id)
	}
	fmt.Println("main finished...")
}
