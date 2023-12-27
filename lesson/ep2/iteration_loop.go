package lesson

import "fmt"

func Ep2_ForLoop() {
	for i := 0 ; i < 10 ; i++ {
		fmt.Print(i, " ")
	}
}

func Ep2_DoWhileLoop() {
	i := 0
	for {
		fmt.Print(i, " ")
		i++
		if i >= 10 {
			break
		}
	}
}

func Ep2_WhileLoop() {
	i := 0
	for i < 10{
		fmt.Print(i, " ")
		i++
	}
}