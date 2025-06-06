package main

import "fmt"

func main() {
	var x []int
	go func() {
		x = make([]int, 10)
	}()
	go func() {
		x = make([]int, 1000000)
	}()
	x[999999] = 1
	fmt.Println(x)
}
