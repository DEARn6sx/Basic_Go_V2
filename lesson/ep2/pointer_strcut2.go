package lesson

type ListNode struct {
	Value int
	Next  *ListNode
}

// Function to add a node to the front of the list
func Ep2_listNode_prepend(head **ListNode, value int) *ListNode {
	newNode := ListNode{Value: value, Next: *head}
	*head = &newNode
	return *head
}

// func main() {
// 	var head *lesson.ListNode

// 	lesson.Ep2_listNode_prepend(&head, 10)
// 	lesson.Ep2_listNode_prepend(&head, 20)

// 	current := head
// 	for current != nil {
// 		fmt.Println(current.Value)
// 		current = current.Next
// 	}
// }
