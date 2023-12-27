package lesson

func Ep2_changeValue(val int) int {
	val = 50
	return val
}

//อ้างอิง address
//Passed by referent (เสมือน) เพราะเป็นการ copy value ไว้อีก address
func Ep2_changeValuePointer(ptr *int) *int {
	*ptr = 50
	return ptr
}




// func main() {
// 	x := 20
// 	lesson.Ep2_changeValue(x)
// 	fmt.Println(x)

// 	lesson.Ep2_changeValuePointer(&x) //ส่งที่อยู่ไป
// 	fmt.Println(x)

// 	y := 10
// 	var p *int = &y
// 	fmt.Println("value of y : ",y)
// 	fmt.Println("value of p : ",*p)
// 	*p = 20
// 	fmt.Println("value of y : ",y)
// }
