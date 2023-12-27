package lesson

import "fmt"

func Ep2_Array() {
	var myArray [3]int
	myArray[0] = 10
	myArray[1] = 20
	myArray[2] = 30
	fmt.Println(myArray)
	for i := 0; i < len(myArray); i++ {
		fmt.Println(myArray[i])
	}
}

func Ep2_Slice() {
	mySlice := []int{10, 20, 30, 40, 50, 60}

	fmt.Println(mySlice)
	fmt.Println(len(mySlice)) //length
	fmt.Println(cap(mySlice)) //capacity
	fmt.Println(mySlice[:])
	fmt.Println(mySlice[3:])
	fmt.Println(mySlice[:3])

	subSlice := mySlice[1:3]
	fmt.Println(subSlice)
	fmt.Println(len(subSlice)) //length
	fmt.Println(cap(subSlice)) //capacity

}

func Ep2_AppendSlice() {
	var mySlice []int

	mySlice = append(mySlice, 10)
	fmt.Println(mySlice)

	mySlice = append(mySlice, 20, 30)
	fmt.Println(mySlice)

	
}

func Ep2_ArrayConvertSlice()  {
	var myArray [3]int
	myArray[0] = 10
	myArray[1] = 20
	myArray[2] = 30
	fmt.Println(myArray)

	mySlice := myArray[:]
	mySlice = append(mySlice, 40,50)
	fmt.Println(mySlice)
}
