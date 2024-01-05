package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host         = "localhost"
	port         = 5432
	databaseName = "mydatabase"
	username     = "myuser"
	password     = "mypassword"
)

func main() {
	// Connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, username, password, databaseName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}	

	print("conntecion db:",db)

}
