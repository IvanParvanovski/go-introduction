package main

import "fmt"

func main() {
	// var intArr [3]int32
	// fmt.Println(intArr[0])
	// fmt.Println(intArr[1:3])

	// intArr[1] = 123
	// fmt.Println(&intArr[0])
	// fmt.Println(&intArr[1])
	// fmt.Println(&intArr[2])

	// var intArr [3]int32 = [3]int32{1, 2, 3}
	// fmt.Println(intArr)

	// intArr := [3]int32{1, 2, 3}
	// fmt.Println(intArr)

	// intArr := [...]int32{1, 2, 3}
	// fmt.Println(intArr)

	// var intSlice []int32 = []int32{4, 5, 6}
	// fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))
	// intSlice = append(intSlice, 7)
	// fmt.Printf("\nThe length is %v with capacity %v", len(intSlice), cap(intSlice))

	// var intSlice []int32 = []int32{4, 5, 6}
	// var intSlice2 []int32 = []int32{8, 9}
	// intSlice = append(intSlice, intSlice2...)
	// fmt.Println(intSlice)

	// var intSlice3 []int32 = make([]int32, 3, 8)

	// var myMap map[string]uint8 = make(map[string]uint8)
	// fmt.Println(myMap)

	// var myMap2 = map[string]uint8{"Adam":23, "Sarah": 45}
	// fmt.Println(myMap2)
	// fmt.Println(myMap2["Adam"])
	// fmt.Println(myMap2["JOgn"])
	// var age, ok = myMap2["Adam"]
	// delete(myMap2, "Adam")
	// if ok {
	// 	fmt.Println("The age is %v", age)
	// } else {
	// 	fmt.Println("Invalid Name")
	// }

	// for name, age := range myMap2 {
	// 	fmt.Printf("Name: %v, Age: %v \n", name, age)
	// }

	// for i, v := range myMap2 {
	// 	fmt.Printf("Index: %v, Value: %v \n", i, v)
	// }

	for  i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
