package repositories

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

func setUpConnection() {
	const (
		host     = "localhost"
		port     = 5432
		user     = "admin"
		password = "admin"
		dbname   = "foodate"
	)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
}

func closeConnection() {
	err := db.Close()
	if err != nil {
		panic(err)
	}
}
