package main

import (
	"fmt"
	"ex9-2/utils"
)

func main() {
	fmt.Println("hi")
	var x uint64 = 0xF0F0F0F0F0F0F0F0
	count := popCount.PopCount(x)
	fmt.Printf("PopCount(%x) = %d\n", x, count)
}
