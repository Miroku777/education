package main

import "fmt"

func binarySearch(array []int, el int) (r int) {
	low := 0
	high := len(array) - 1
	for low <= high {
		mid := (low + high) / 2
		if el > array[mid] {
			low = mid + 1
		} else if el < array[mid] {
			high = mid - 1
		} else {
			r = mid
			break
		}
	}
	return
}

func main() {
	in := make([]int, 101)
	for i := 0; i < 101; i++ {
		in[i] = i
	}
	fmt.Println(binarySearch(in, 44))
}
