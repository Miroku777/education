package main

import (
	"bufio"
	"fmt"
	"log"
	"net/mail"
	"os"
	"strings"

	_ "modernc.org/sqlite"
)

func main() {
	db := connectDB()
	defer db.Close()

	sc := bufio.NewScanner(os.Stdin)
	insert := ">"
	findOne := "<"
	for {
		fmt.Println("All users")
		users := findAll(db)
		for _, u := range users {
			fmt.Printf("%2d email: %-15s passwd: %s;\n", u.id, u.email, u.password.String)
		}
		fmt.Printf("Добавить запись: %s | Получить запись: %s\n\n", insert, findOne)

		input := make([]string, 0)
		fmt.Print("операция выполнения: ")
		sc.Scan()
		operate := strings.TrimSpace(sc.Text())

		fmt.Print("email: ")
		sc.Scan()
		email := strings.TrimSpace(sc.Text())
		addres, err := mail.ParseAddress(email)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Print("passwd: ")
		sc.Scan()
		password := strings.TrimSpace(sc.Text())
		input = []string{operate, addres.Address, password}
		if len(input) < 3 {
			log.Println("Input cannot to process, len < ", len(input))
			continue
		}
		if input[0] == insert {
			bool := insertUser(db, input)
			if bool {
				log.Println("Insert success")
			}
		} else if input[0] == findOne {
			findByEmail(db, input)
		}
		//fmt.Print("\033[H\033[2J")
	}
}
