package lesson

import (
	"fmt"
	"log"
)

func Ep4_getAllProduct() {
	defer db.Close()
	products, err := getAllProduct()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(products)
}

func getAllProduct() ([]Product, error) {

	rows, err := db.Query("SELECT id, name, price, supplier_id FROM products")

	if err != nil {
		return nil, err
	}

	var products []Product

	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Supplier_id)
		if err != nil {
			return nil, err
		}
		products = append(products, p)

	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}
