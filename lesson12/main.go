package main

func InsertInSliceZero(v int) {
	mySlice := make([]int, 0)
	for i := 0; i < v; i++ {
		mySlice = append(mySlice, v)
	}
}
func InsertInSliceV(v int) {
	mySlice := make([]int, v)
	for i := 0; i < v; i++ {
		mySlice = append(mySlice, v)
	}
}

func IsEven(value int) string {
	if value%2 == 0 {
		return "Yes"
	}
	return "No"
}
func main() {
}
