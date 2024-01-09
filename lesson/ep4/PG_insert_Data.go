package lesson

import (
	"database/sql"
	"fmt"
	"log"
)

func Ep4_createProduct(n string, p int, sid int) {
	// Check if the product already exists
	if !productExists(n, p, sid) {
		fmt.Println("Before insert product")
		// Insert only if the product doesn't exist
		// err := createProduct(&Product{Name: n, Price: p, Supplier_id: sid})
		// if err != nil {
		// 	log.Fatal(err)
		// }

		fmt.Println("Successfully Created!")
	} else {
		fmt.Println("Product already exists.")
	}
}

func createProduct(product *Product) error {
	_, err := db.Exec(
		"INSERT INTO public.products(name, price, supplier_id) VALUES ($1, $2, $3);",
		product.Name,
		product.Price,
		product.Supplier_id,
	)

	if err != nil {
		log.Printf("Error inserting product: %v", err)
	}

	return err
}

func productExists(name string, price int, supplierID int) bool {
	var count int
	err := db.QueryRow(
		"SELECT COUNT(*) FROM public.products WHERE name = $1 AND price = $2 AND supplier_id = $3",
		name, price, supplierID,
	).Scan(&count)

	if err != nil && err != sql.ErrNoRows {
		log.Printf("Error checking if product exists: %v", err)
	}

	fmt.Println("count is : ", count)

	return count > 0
}
