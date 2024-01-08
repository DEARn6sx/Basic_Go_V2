package lesson

import (
	"fmt"
	"log"
)

func Ep4_selectProduct(id int) {
	product, err := getProduct(id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(product)

}

func getProduct(id int) (Product, error) {
	var p Product
	row := db.QueryRow("SELECT id, name, price FROM products WHERE id = $1",
		id,
	)
	err := row.Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return Product{}, err
	}

	return p, nil

}
