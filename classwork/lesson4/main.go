package main

import "fmt"

func main() {
	message := "jr#wkh#ehvw"
	for _, v := range message {
		fmt.Printf("%c", decrypt(v))
	}

}

func crypt(r rune) rune {
	var res rune = r + 3
	if res > 'z' {
		res -= 26
	}
	return res
}
func decrypt(r rune) rune {
	var res rune = r - 3
	if res > 'z' {
		res += 26
	}
	return res
}
