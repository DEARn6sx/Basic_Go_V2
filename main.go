package main

import (
	"fmt"

	"github.com/DEAR/go-example/package_go"
	"github.com/google/uuid"
)

func main() {
	ep1()
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
