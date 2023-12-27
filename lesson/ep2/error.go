package lesson

import "errors"

func Ep2_error_Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannit divide by zero")
	}
	return a / b , nil
}


// func main() {
// 	result, err := lesson.Ep2_error_Divide(10,2)
// 	if err != nil {
// 		fmt.Println("Error: ",err)
// 		return
// 	}
// 	fmt.Println("Result: ",result)
// }