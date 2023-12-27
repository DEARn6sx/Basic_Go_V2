package lesson

import "fmt"

type Speaker interface {
	Speak() string
	Leg() string
}

type Dog struct {
	Name string
}

// สร้าง method
func (d Dog) Speak() string {
	return "wooof"
}

func (d Dog) Leg() string {
	return "4 leg"
}

type Human struct {
	Name string
}

func (h Human) Speak() string {
	return "Hello"
}

func (h Human) Leg() string {
	return "2 Leg"
}

func Ep2_Inerface_makeSound(s Speaker) {
	fmt.Println(s.Speak())
	fmt.Println(s.Leg())
}




// func main() {
	
// 	humanSpeaker := lesson.Human{
// 		Name: "DEAR",
// 	}

// 	dognSpeaker := lesson.Dog{
// 		Name: "Bobo",
// 	}

// 	lesson.Ep2_Inerface_makeSound(humanSpeaker)
// 	lesson.Ep2_Inerface_makeSound(dognSpeaker)
// }