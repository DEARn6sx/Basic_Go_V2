package main

import (
	"fmt"

	lesson "github.com/DEAR/go-example/lesson/ep2"
)

func main() {

	rect := lesson.Rectangle{Length: 5, Width: 3}
	areaRect := lesson.Ep2_calculateArea(rect)
	fmt.Printf("Rectangle Area: %f\n", areaRect)

	tri := lesson.Triangle{Base: 10, Height: 15}
	areaTri := lesson.Ep2_calculateArea(tri)
	fmt.Printf("Triangle Area: %f\n", areaTri)

	circle := lesson.Circle{Radius: 2}
	areaCircle := lesson.Ep2_calculateArea(circle)
	fmt.Printf("Circle Area: %f\n", areaCircle)
}
