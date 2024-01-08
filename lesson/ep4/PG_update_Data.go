package lesson

import (
	"fmt"
	"log"
)

func Ep4_updateProduct(id int, n string, p int) {
	// insert
	product, err := updateProduct(id, &Product{Name: n, Price: p})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully Updated!", product)
}


func updateProduct(id int, product *Product) (Product, error) {
	var p Product
	row := db.QueryRow(
		"UPDATE  public.products SET name=$1, price=$2 WHERE id =$3 RETURNING id, name, price;",
		product.Name,
		product.Price,
		id,
	)
	err := row.Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return Product{}, err
	}

	return p, nil

}
