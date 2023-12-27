package lesson

import "fmt"

func Ep2_If_Else() {
	score := 62
	var grade string

	if score >= 70 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}

	if score >= 80 {
		grade = "A"
	} else if score >= 70 {
		grade = "B"
	} else if score >= 60 {
		grade = "C"
	} else {
		grade = "F"
	}
	fmt.Printf("Your grad is %s\n", grade)
}
