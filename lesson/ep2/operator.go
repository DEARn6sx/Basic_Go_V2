package lesson

import "fmt"

func Ep2_Operator() {
	var fullname string = "DEARZA"
	var age = 26
	address := "Roiet"

	fullname = "DEAR" + "DEAR" + address
	age = age * 2

	fmt.Println(fullname, age)

	a := 10
	b := 3
	fmt.Println(a + b) // 13
	fmt.Println(a - b) // 7
	fmt.Println(a * b) // 30
	fmt.Println(a / b) // 3
	fmt.Println(a % b) // 1

}

func Ep2_RelationalOperator()  {
	a := 10
	b := 3
	fmt.Println(a == b) // false
	fmt.Println(a != b) // true
	fmt.Println(a > b)  // true
	fmt.Println(a < b)  // false
	fmt.Println(a >= b) // true
	fmt.Println(a <= b) // false
}

func Ep2_LogicOperator()  {
	c := true
	d := false
	fmt.Println(c && d) // false
	fmt.Println(c || d) // true
	fmt.Println(!c)     // false
}