package lesson

import "fmt"

func Ep2_Variable() {
	var fullname string = "DEARZA"
	var age = 26
	address := "Roiet"
	fmt.Println(fullname, age)
	fmt.Printf("Hello %s Yaaaay my age is %d myaddress is %s\n",
		fullname, age, address)
	fmt.Printf("Type name is %T \nType age is %T \nType address is %T \n",
		fullname, age, address)

}
