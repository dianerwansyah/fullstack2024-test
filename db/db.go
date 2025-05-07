package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	connStr := "user=dian password=1234 dbname=my_client sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}
