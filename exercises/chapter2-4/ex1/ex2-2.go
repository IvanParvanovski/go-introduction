package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
	"strconv"
)

func main() {
	// Display menu
	message := []string{
		"Your options are:", 
		"1 -> Convert Celsius to Fahrenheit", 
		"2 -> Convert km to meters", 
		"3 -> Convert BGN to Euro",
	}
	fmt.Println(strings.Join(message, "\n"))

	scanner := bufio.NewScanner(os.Stdin)
	var err error

	// Get user's option and handle errors
	var option int
	fmt.Print("Type an option: ")
	if scanner.Scan() {
		option, err = strconv.Atoi(scanner.Text())
	}

	if err != nil {
		fmt.Println("Invalid input!")
		return
	}
	
	// Read the value that needs conversion
	fmt.Println(message[option])

	var value float64
	fmt.Print("\nType corresponding value: ")
	if scanner.Scan() {
		value, err = strconv.ParseFloat(scanner.Text(), 64)
	}

	// Compute result
	var result float64
	switch option {
		case 1:
			result = (value * 9/5) + 32
		case 2: 
			result = value * 1000
		case 3:
			result = value / 1.95583
	}

	fmt.Println(result)
}
