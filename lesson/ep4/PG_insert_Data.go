package lesson

import (
	"fmt"
	"log"
)

func Ep4_createProduct(n string, p int, sid int) {
	// insert
	err := createProduct(&Product{Name: n, Price: p, Supplier_id: sid})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully Created!")
}

func createProduct(product *Product) error {
	_, err := db.Exec(
		"INSERT INTO public.products(name, price,supplier_id) VALUES ($1, $2,$3);",
		product.Name,
		product.Price,
		product.Supplier_id,
	)

	if err != nil {
		log.Printf("Error inserting product: %v", err)
	}

	return err
}
