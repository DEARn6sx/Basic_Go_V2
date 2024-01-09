package main

import (
	lesson "Go_lang/lesson/ep4"

	_ "github.com/lib/pq"
)

func main() {
	lesson.Connect_DB()

	//lesson.Ep4_createProduct("nooode", 555, 1)

	//lesson.Ep4_selectProduct(8)

	//lesson.Ep4_updateProduct(11,"jjjjj",1111)

	//lesson.Ep4_deleteProductt(19)

	//lesson.Ep4_getAllProduct()

	lesson.Ep4_getJoinProduct()
}
