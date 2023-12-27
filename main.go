package main

import (
	"fmt"

	lesson "github.com/DEAR/go-example/lesson/ep2"
)

func main() {
	var head *lesson.ListNode

	lesson.Ep2_listNode_prepend(&head, 10)
	lesson.Ep2_listNode_prepend(&head, 20)

	current := head
	for current != nil {
		fmt.Println(current.Value)
		current = current.Next
	}
}
