package lesson

type Std struct {
	Firstname string
	Lastname  string
}

func (student Std) Fullname() string {
	fullname := student.Firstname + " " + student.Lastname
	return fullname
}

// func main() {
// 	// Create a new Std instance
// 	studentInstance := lesson.Std{
// 		Firstname: "DEAR",
// 		Lastname:  "Nantawat",
// 	}

// 	// Access the Fullname method
// 	fullname := studentInstance.Fullname()

// 	// Print the result
// 	fmt.Println("Fullname:", fullname)
// }
