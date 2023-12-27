package lesson

import "fmt"

func Ep2_FuncInputParams(name string, round int) {
	for i := 0; i < round; i++ {
		fmt.Printf("Hello %s\n", name)
	}
}

func Ep2_FuncReturn(num1 int, num2 int) int {
	return num1 + num2
}
