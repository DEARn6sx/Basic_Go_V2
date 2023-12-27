package lesson

import (
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Length, Width float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

type Triangle struct {
	Base, Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func Ep2_calculateArea(s Shape) float64 {
	return s.Area()
}



// func main() {

// 	rect := lesson.Rectangle{Length: 5, Width: 3}
// 	areaRect := lesson.Ep2_calculateArea(rect)
// 	fmt.Printf("Rectangle Area: %f\n", areaRect)

// 	tri := lesson.Triangle{Base: 10, Height: 15}
// 	areaTri := lesson.Ep2_calculateArea(tri)
// 	fmt.Printf("Triangle Area: %f\n", areaTri)

// 	circle := lesson.Circle{Radius: 2}
// 	areaCircle := lesson.Ep2_calculateArea(circle)
// 	fmt.Printf("Circle Area: %f\n", areaCircle)
// }