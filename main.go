package main

import (
	"fmt"

	lesson "github.com/DEAR/go-example/lesson/ep2"
)

func main() {
	err := lesson.Login("admin", "1234")
	if err != nil {
		switch e := err.(type) {
		case *lesson.LoginError:
			fmt.Println("Custom error occurred:", e)
		default:
			fmt.Println("Generic error occurred:", e)
		}
		return
	}
	fmt.Println("Login successful!")
}
