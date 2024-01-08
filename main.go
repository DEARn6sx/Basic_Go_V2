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

// ประกาศ db เป็น global เพื่อให้ db เป็นตัวแทนในการพูดคุย
var db *sql.DB

type Product struct {
	ID         int
	Name       string
	Price      int
}

func main() {
	// Connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, username, password, databaseName)

	sdb, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal(err)
	}

	db = sdb

	//check db ส่งสัญญาณมาหรือไม่
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("conntecion sucess")

	err = createProducts(&Product{Name: "6", Price: 6})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Create Successful")
}

func createProducts(product *Product) error {
	//Exec เป็นการ run คำสั่งที่ไม่ต้องการ return กลับมา แค่ต้องการ return ว่า error มั้ยแค่นั้น
	_, err := db.Exec(
		"INSERT INTO public.products(name, price)VALUES ($1, $2);",
		product.Name,
		product.Price,
	)
	return err
}
