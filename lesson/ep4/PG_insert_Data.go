package lesson

import (
	"database/sql"
	"fmt"
	"log"
)

func Ep4_createProduct() {
	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// Insert
	err = createProduct(tx, &Product{Name: "x", Price: 1, Supplier_id: 2})
	if err != nil {
		// Rollback the transaction on error
		tx.Rollback()
		log.Fatal(err)
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully Created!")
}

func createProduct(tx *sql.Tx, product *Product) error {
	_, err := tx.Exec(
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
