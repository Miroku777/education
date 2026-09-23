package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.MkdirAll("../lesson12", 0755)
	if err != nil {
		fmt.Println(err)
		return
	}
	file, err := os.Create("../lesson12/main.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	n, err := file.WriteString("package main\nfunc main(){\n}")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n)
}
