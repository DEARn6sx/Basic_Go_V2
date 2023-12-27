package main

import (
	"fmt"

	lesson "github.com/DEAR/go-example/lesson/ep2"
)

func main() {
	err := lesson.Login("user","pass")
	if err != nil {
		switch e := err.(type){
		case *lesson.LoginError:
			fmt.Println("Custom error occurred:",e)
		default:
			fmt.Println("Generic error occurred:",e)
		}
		return
	}
}
