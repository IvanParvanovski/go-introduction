package main 

import (
	"fmt" 
	// "log"
)

func main() {
	// fmt.Println("start")
	
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		log.Println("Error:", err)
	// 	}
	// }()

	// panic("something bad happened")
	// fmt.Println("end")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	x := 0
	fmt.Println(1 / x)
	fmt.Println("Hi")

}