package lesson

import "fmt"

func Ep2_SwichCase() {
	var dayOfWeek int
	fmt.Print("Input date 1 - 7 :")
	fmt.Scanf("%d",&dayOfWeek)

	switch dayOfWeek {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid Day")
	}
}
