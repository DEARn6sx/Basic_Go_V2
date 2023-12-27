package main

import (
	"fmt"

	lesson "github.com/DEAR/go-example/lesson/ep2"
)

func main() {
	x := 20
	lesson.Ep2_changeValue(x)
	fmt.Println(x)

	lesson.Ep2_changeValuePointer(&x) //ส่งที่อยู่ไป
	fmt.Println(x)

	y := 10
	var p *int = &y
	fmt.Println("value of y : ",y)
	fmt.Println("value of p : ",*p)
	*p = 20
	fmt.Println("value of y : ",y)
}
