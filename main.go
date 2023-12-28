package main

import (
	"fmt"
	"log"
	"net/http"

	lesson "github.com/DEAR/go-example/lesson/ep3"
)

func main() {
	http.HandleFunc("/hello", lesson.Ep3_netHttpHelloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080",nil); err != nil {
		log.Fatal(err)
	}
}



//go mod init github.com/DEAR/go-example ลงอันนี้ก่อนเพื่อเรียกใช้งาน module ของ GO
