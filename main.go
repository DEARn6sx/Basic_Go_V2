package main

import (
	"fmt"

	lesson "github.com/DEAR/go-example/lesson/ep2"
)

func main() {
	result, err := lesson.Ep2_error_Divide(10,2)
	if err != nil {
		fmt.Println("Error: ",err)
		return
	}
	fmt.Println("Result: ",result)
}
