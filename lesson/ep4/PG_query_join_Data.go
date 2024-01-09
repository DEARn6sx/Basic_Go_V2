package lesson

import (
	"fmt"
	"log"
)

func Ep4_getJoinProduct() {
	defer db.Close()
	products, err := getJoinProduct()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(products)
}

func getJoinProduct() ([]ProductWithSupplier, error) {

	rows, err := db.Query(`
	SELECT
		p.id AS product_id,
		p.name AS product_name,
		p.price,
		s.name AS supplier_name
	FROM
		products p
	INNER JOIN suppliers s
		ON p.supplier_id = s.id;`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

  var products []ProductWithSupplier
  for rows.Next() {
    var p ProductWithSupplier
    err := rows.Scan(&p.ProductID, &p.ProductName, &p.Price, &p.SupplierName)
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
