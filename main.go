package main

import (
	lesson "Go_lang/lesson/ep4"

	_ "github.com/lib/pq"
)

func main() {
	lesson.Connect_DB()

	lesson.Ep4_createProduct("TESTinsert2",2,2)

	//lesson.Ep4_selectProduct(8)	

	//lesson.Ep4_updateProduct(11,"jjjjj",1111)
}


