package package_go

import (
	"fmt"

	internal "github.com/DEAR/go-example/package_go/internal/intenal_pack"
)

func testHello() {
	fmt.Println("Hello Gogogogogogog22222")
}

func SayHello2() {
	testHello()
	internal.SayHello()
}
