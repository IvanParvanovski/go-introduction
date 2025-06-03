package main

import (
	"os"
	"bufio"
	"fmt"
)

func main() {
	file, ferr := os.Open("customers-100.csv")

	if ferr != nil {
		panic(ferr)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	
}