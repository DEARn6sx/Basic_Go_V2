package lesson

import (
	"fmt"
)

type Student struct {
	Name   string
	Weight float32
	Height int
	Grade  string
}

func Ep2_Struct() {
	var student1 Student
	student1.Name = "DEAR"
	student1.Weight = 67.5
	student1.Height = 176
	student1.Grade = "C"

	fmt.Println(student1)
	fmt.Printf("Name : %s\n", student1.Name)
	fmt.Printf("Weight : %.2f\n", student1.Weight)
	fmt.Printf("Height : %d\n", student1.Height)
	fmt.Printf("Grade : %s\n", student1.Grade)
}

func Ep2_StructArray() {
	var student [3]Student
	student[0].Name = "DEAR"
	student[0].Weight = 67.5
	student[0].Height = 176
	student[0].Grade = "C"

	student[1].Weight = 78
	student[1].Name = "xza"
	student[1].Height = 222
	student[1].Grade = "A"

	for i := 0; i < len(student); i++ {
		fmt.Println(student[i])
	}

}

func Ep2_StructMap() {
	students := make(map[string]Student)

	students["std01"] = Student{Name: "DEAr", Weight: 67, Height: 188, Grade: "C"}
	students["std02"] = Student{Name: "zazaza", Weight: 55, Height: 555, Grade: "Xa"}
	students["std03"] = Student{Name: "Sqa", Weight: 33, Height: 256, Grade: "A"}

	fmt.Println("Students Map: ", students)

	fmt.Println("Student std01:", students["std01"])
	fmt.Println("Student std02:", students["std02"])
	fmt.Println("Student std03:", students["std03"])

}

type Person struct {
	Name    string
	Age     int
	Address Adress
}

type Adress struct {
	Street  string
	City    string
	Zipcode int
}

func Ep2_StructInStruct()  {
	var person Person
	person.Name =" DEAR"
	person.Age = 26
	person.Address = Adress{
		Street: "478 m 1",
		City: "Atsamart Roi et",
		Zipcode: 45160,
	}

	Bob := Person{
		Name: "Bob",
		Age: 50,
		Address: Adress{
			Street: "123 m 4",
			City: "kk",
			Zipcode: 40000,
		},
	}

	fmt.Println(person)
	fmt.Println(Bob)
}
