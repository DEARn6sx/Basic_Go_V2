package lesson

import "fmt"

func Ep2_Map() {
	myMap := make(map[string]int) //key: string , value: int
	myMap["a"] = 5
	myMap["b"] = 10
	myMap["c"] = 20

	fmt.Println(myMap)
	fmt.Println("a is : ", myMap["a"])

	myMap["d"] = 50

	delete(myMap, "c")

	fmt.Println(myMap)

	for key, value := range myMap {
		fmt.Printf("%s is : %d \n", key, value)
	}

	myMap["c"] = 508
	//checking "key" is exitsts
	value, key := myMap["c"]
	if key {
		fmt.Printf("Value is : %d", value)
	} else {
		fmt.Println("Not found")
	}

}
