package main

import (
	"fmt"

	lesson "github.com/DEAR/go-example/lesson/ep2"
)

func main() {
	// Create a new Std instance
	studentInstance := lesson.Std{
		Firstname: "DEAR",
		Lastname:  "Nantawat",
	}

	// Access the Fullname method
	fullname := studentInstance.Fullname()

	// Print the result
	fmt.Println("Fullname:", fullname)
}
