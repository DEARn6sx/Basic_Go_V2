package main

import (
	"fmt"

	"github.com/DEAR/go-example/package_go"
	"github.com/google/uuid"
)

func main() {
	ep2()
}

func ep2() {
	var fullname string = "DEARZA"
	var age = 26
	address := "Roiet"
	fmt.Println(fullname, age)
	fmt.Printf("Hello %s Yaaaay my age is %d myaddress is %s\n",
		fullname, age, address)
	fmt.Printf("Type name is %T \nType age is %T \nType address is %T \n",
		fullname, age, address)

}

func ep1() {
	id := uuid.New()
	id_string := uuid.New().String()
	fmt.Println("Gogogogogogohohohozazac")
	fmt.Printf("UUID: %s\n", id)
	fmt.Printf("UUID_String: %s\n", id_string)
	package_go.SayHello() //read from package not folder
	package_go.SayHello2()
}
