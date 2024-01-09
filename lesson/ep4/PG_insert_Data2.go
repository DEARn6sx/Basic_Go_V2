package lesson

import (
	"fmt"
	"log"
)

func Ep4_createProduct2(n string, p int, sid int) {
	// Check if the product already exists

	// Insert only if the product doesn't exist
	product, err := createProduct2(&Product{Name: n, Price: p, Supplier_id: sid})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully Created!", product)

}

func createProduct2(product *Product) (Product, error) {
	var p Product
	row := db.QueryRow(
		"INSERT INTO public.products(name, price, supplier_id) VALUES ($1, $2, $3) RETURNING id, name, price, supplier_id;",
		product.Name,
		product.Price,
		product.Supplier_id,
	)
	err := row.Scan(&p.ID, &p.Name, &p.Price, &p.Supplier_id)
	if err != nil {
		return Product{}, err
	}

	return p, nil
}
