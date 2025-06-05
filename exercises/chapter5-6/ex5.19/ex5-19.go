package main

import "fmt"


// func main() {
// 	x := 0
//     defer recoveringFunc()
	
// 	y := 1 / x
// 	fmt.Println(y)
// }

// func recoveringFunc() string {
// 	if r := recover(); r != nil {
// 		fmt.Println("cannot divide by 0")
// 		return "cannot divide by 0"
// 	}
// 	fmt.Println("Error msg")
// 	return "Error msg"
// }

func main() {
	fmt.Println("Result: ", recoverReturn())
}
func recoverReturn() (result int) {
	defer func() {
		if r := recover(); r != nil {
			result = 42
		}
	}()

	panic("force a return without return")
}


// func main() {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Println("Recovered: cannot divide by 0")
// 		}
// 	}()

// 	x := 0
// 	y := 1 / x // This causes a panic
// 	fmt.Println("y =", y) // Never reached
// }
