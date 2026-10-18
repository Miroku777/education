package main

import (
	"database/sql"
	"fmt"
	"log"
)

func insertUser(db *sql.DB, array []string) (b bool) {
	result, err := db.Exec(`Insert into users(email, password) values(?, ?)`, array[1], array[2])
	if err != nil {
		log.Println(err)
		return
	}
	ins, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("Isert %d string\n", ins)
	return true
}

func findByEmail(db *sql.DB, array []string) {
	row := db.QueryRow(`Select id, email, password From users Where email = ?`, array[1])
	var user users
	err := row.Scan(&user.id, &user.email, &user.password)
	if err != nil {
		fmt.Printf("%s: пользователь не найден\n\n", err)
		return
	}
	if !user.password.Valid || user.password.String != array[2] {
		fmt.Printf("пароль неверный\n\n")
		return
	}
	fmt.Printf("пароль верный\n\n")
}

func findAll(db *sql.DB) []users {
	rows, err := db.Query(`Select id, email, password From users`)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer rows.Close()
	var userlist = []users{}
	for rows.Next() {
		var user users
		var password sql.NullString
		err := rows.Scan(&user.id, &user.email, &password)
		if err != nil {
			log.Println(err)
			continue
		}
		if password.Valid {
			user.password = password
		}
		userlist = append(userlist, user)
	}
	if err := rows.Err(); err != nil {
		log.Println(err)
		return nil
	}
	return userlist
}
