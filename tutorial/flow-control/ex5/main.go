package main

import (
	"fmt"
)

func main() {
	var i interface{} = [3]int{}
	switch i.(type) {
		case int:
			fmt.Println("i is an int")
			fmt.Println("This will print too")
		case float64:
			fmt.Println("i is a float64")
		case string:
			fmt.Println("i is a string")
		default:
			fmt.Println("it is random")
	}
}