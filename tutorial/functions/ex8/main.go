package main 

import (
	"fmt"
)

func main() {
	func() {
		msg := "Hello Go!"
		fmt.Println(msg)
	}()

	for i := 0; i < 5; i++ {
		func() {
			fmt.Println(i)
		}()
		// func(i int) {
		// 	fmt.Println(i)
		// }(i)
	}
}
