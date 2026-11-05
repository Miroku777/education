package main

//go get github.com/lib/pq
//github.com/mattn/go-sqlite3

import (
	"database/sql"
	"fmt"
	"log"
	//_ "modernc.org/sqlite"
)

type Phone1 struct {
	id     int
	vendor string
	model  string
	price  int64
}

func main1() {
	const pgConfig = "user=postgres password=postgres dbname=mydb sslmode=disable port=5432"
	db, err := sql.Open("postgres", pgConfig)
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

	phones := []Phone1{}
	for rows.Next() {
		p := Phone1{}
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
