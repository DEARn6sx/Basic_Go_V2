package lesson

import "fmt"

type LoginError struct {
	Username string
	Message  string
}

//implement method error()
func (e *LoginError) Error() string {
	return fmt.Sprintf("Login error for user '%s' : %s" ,e.Username, e.Message)
}

func Login(username, password string) error  {
	if username != "admin" || password != "1234" {
		return &LoginError{Username: username, Message: "invalid credentials"}
	}
	return nil
}


// func main() {
// 	err := lesson.Login("user","pass")
// 	if err != nil {
// 		switch e := err.(type){
// 		case *lesson.LoginError:
// 			fmt.Println("Custom error occurred:",e)
// 		default:
// 			fmt.Println("Generic error occurred:",e)
// 		}
// 		return
// 	}
// }