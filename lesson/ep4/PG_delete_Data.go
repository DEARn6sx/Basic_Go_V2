package lesson

import (
	"fmt"
	"log"
)

func Ep4_deleteProductt(id int) {
	// insert
	err := deleteProduct(id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully Deleted!")
}

func deleteProduct(id int) error {

	_, err := db.Exec(
		"DELETE FROM  public.products WHERE id=$1 ;",
		id,
	)

	return err

}
