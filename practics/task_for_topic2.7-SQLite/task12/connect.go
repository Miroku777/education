package main

import (
	"database/sql"
	"log"
)

type users struct {
	id       int
	email    string
	password sql.NullString
}

func connectDB() (db *sql.DB) {
	db, err := sql.Open("sqlite", "auth.db")
	if err != nil {
		log.Println(err)
		return nil
	}
	_, err = db.Exec(`
		Create Table IF NOT EXISTS users(
			id integer primary key autoincrement,
			email text unique not null,
			password text
		);
		Insert into users(email, password) values('test@mail.ru', '123');
			`)
	if err != nil {
		log.Println(err)
		return
	}
	return
}
