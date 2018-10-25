package main

import (
	"database/sql"
	"log"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "root@localhost/go_learn")

	if err != nil {
		log.Fatal(err)
	}

	return db
}
