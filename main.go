package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"  // or the Docker service name if running in another container
	port     = 5432         // default PostgreSQL port
	user     = "myuser"     // as defined in docker-compose.yml
	password = "mypassword" // as defined in docker-compose.yml
	dbname   = "mydatabase" // as defined in docker-compose.yml
)

var db *sql.DB

type Product struct {
	ID    int
	Name  string
	Price int
}

func main() {
	// Connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open a connection
	sdb, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	db = sdb

	// Check the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected!")

	//insert

	// err = createProduct(&Product{Name: "Go product", Price: 222})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Successfully Created!")

	//search

	product, err := getProduct(2)
	fmt.Println(product)
}

func createProduct(product *Product) error {

	_, err := db.Exec(
		"INSERT INTO public.products( name, price)VALUES ($1, $2 );",
		product.Name,
		product.Price,
	)

	return err
}

func getProduct(id int) (Product, error)  {
	var p Product
	row := db.QueryRow("SELECT id, name, price FROM products WHERE id = $1",
	id,
	)
	err := row.Scan(&p.ID,&p.Name, &p.Price)
	if err != nil {
		return Product{}, err
	}

	return p, nil

}