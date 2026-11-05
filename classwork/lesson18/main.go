package main

//github.com/mattn/go-sqlite3

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

type Phone struct {
	id     int
	vendor string
	model  string
	price  int64
}

func main() {
	db, err := sql.Open("sqlite", "shop.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	res, err := db.Exec(`
	Insert into phones (Vendor, Model, Phone)
	Values (?, ?, ?), "Huawey", "Simple 1", 1000`)
	if err != nil {
		fmt.Println(err)
		return
	}
	i, err := res.RowsAffected()
	fmt.Print(i)

	rows, err := db.Query(`Select * FROM phones;`)
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()

	phones := []Phone{}
	for rows.Next() {
		p := Phone{}
		var price sql.NullInt64
		err := rows.Scan(&p.id, &p.vendor, &p.model, &p.price)
		if err != nil {
			log.Println(err)
			continue
		}
		if price.Valid {
			p.price = price.Int64
		} else {
			p.price = 0
		}
		phones = append(phones, p)
	}
	fmt.Println(phones)

}
