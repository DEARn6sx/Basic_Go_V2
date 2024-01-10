package lesson

import (
	"fmt"
	"log"
)

func Ep4_updateProduct(id int, n string, p int, sp int) {
	// insert
	product, err := updateProduct(id, &Product{Name: n, Price: p, Supplier_id: sp})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully Updated!", product)
}

func updateProduct(id int, product *Product) (Product, error) {
	var p Product
	row := db.QueryRow(
		"UPDATE  public.products SET name=$1, price=$2, supplier_id=$3 WHERE id =$4 RETURNING id, name, price,supplier_id;",
		product.Name,
		product.Price,
		product.Supplier_id,
		id,
	)
	err := row.Scan(&p.ID, &p.Name, &p.Price, &p.Supplier_id)
	if err != nil {
		return Product{}, err
	}

	return p, nil

}
