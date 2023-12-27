package main

import (
	lesson "github.com/DEAR/go-example/lesson/ep2"
)

func main() {

	human := lesson.Human{
		Name: "DEAR",
	}

	dog := lesson.Dog{
		Name: "Bobo",
	}

	lesson.Ep2_Inerface_makeSound(human)
	lesson.Ep2_Inerface_makeSound(dog)
}
