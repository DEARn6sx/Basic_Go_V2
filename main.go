package main

import (
	lesson "Go_lang/lesson/ep4"

	_ "github.com/lib/pq"
)

func main() {
	lesson.Connect_DB()

	//lesson.Ep4_createProduct("99",9)

	//lesson.Ep4_selectProduct(8)	

	lesson.Ep4_updateProduct(11,"jjjjj",1111)
}


