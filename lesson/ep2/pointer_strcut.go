package lesson

type Employee struct {
	Name string
	Salary int
}

func Ep2_pointerStruct_giveRaise(e *Employee, raise int) int {
		e.Salary += raise
	return e.Salary
}


// func main() {
// 	emp := lesson.Employee{
// 		Name: "DEAR",
// 		Salary: 20000,
// 	}
// 	lesson.Ep2_pointerStruct_giveRaise(&emp,5000)
// 	fmt.Println("After raise :", emp)
// }