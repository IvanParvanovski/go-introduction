package main

import (
	"errors"
	"fmt"
)

func main() {
	var printValue string = "Hello World"
	printMe(printValue)

	var numerator int = 11
	var denominator int = 2
	var result, remainder, err = intDivision(numerator, denominator)
	
	// if err != nil {
	// 	fmt.Print(err.Error())
	// } else if remainder == 0 {
	// 	fmt.Printf("The result of the integer division is %v", result)
	// }

	switch {
		case err != nil:
			fmt.Printf(err.Error())
		case remainder == 0:
			fmt.Printf("The result of the integer division is %v", result)
		default:
			fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
	}

	switch remainder {
		case 0:
			fmt.Printf("The divison was exact")
		case 1, 2:
			fmt.Printf("The divison was close")
		default:
			fmt.Printf("Default statements")
	}

	fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Cannot Divide by Zero")
		return 0, 0, err
	}

	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}

