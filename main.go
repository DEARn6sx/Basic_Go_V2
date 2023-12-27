package main

import (
	"fmt"

	lesson "github.com/DEAR/go-example/lesson/ep2"
)

func main() {
	emp := lesson.Employee{
		Name: "DEAR",
		Salary: 20000,
	}
	lesson.Ep2_pointerStruct_giveRaise(&emp,5000)
	fmt.Println("After raise :", emp)
}
