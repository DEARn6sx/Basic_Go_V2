package main

import (
	lesson "Go_lang/lesson/ep4"

	_ "github.com/lib/pq"
)

func main() {
	lesson.Connect_DB()

	lesson.Ep4_createProduct2("5", 1, 2)

	//lesson.Ep4_selectProduct(8)

	//lesson.Ep4_updateProduct(11,"jjjjj",1111)

	//lesson.Ep4_deleteProductt(19)
}
