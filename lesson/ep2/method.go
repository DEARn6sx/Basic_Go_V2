package lesson

type Std struct {
	Firstname string
	Lastname  string
}

func (student Std) Fullname() string {
	fullname := student.Firstname + " " + student.Lastname
	return fullname
}
