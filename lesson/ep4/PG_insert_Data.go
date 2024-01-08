package lesson

import (
	"fmt"
	"log"
)

func Ep4_createProduct(n string, p int) {
	// insert
	err := createProduct(&Product{Name: n, Price: p})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully Created!")
}

func createProduct(product *Product) error {
	_, err := db.Exec(
		"INSERT INTO public.products(name, price) VALUES ($1, $2);",
		product.Name,
		product.Price,
	)

	if err != nil {
		log.Printf("Error inserting product: %v", err)
	}

	return err
}
