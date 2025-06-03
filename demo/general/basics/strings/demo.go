package main

import (
	"fmt"
	"strings"
)

func main() {
	// var myString = "råsumç"
	// var indexed = myString[0]
	// fmt.Printf("%v, %T\n", indexed, indexed)

	// for i, v := range myString {
	// 	fmt.Println(i, v)
	// }

	// fmt.Printf("\nThe length of 'mystring' is %v", len(myString))

	// var myRune = 'a'
	// fmt.Printf("\nmyRune = %v", myRune)

	// var strSlice = []string{"s", "u", "b", "s", "c", "r", "i", "b", "e"}
	// var catStr = ""
	
	// for i := range strSlice {
	// 	catStr += strSlice[i]
	// }
	
	// fmt.Printf("\n%v", catStr)

	var strSlice = []string{"s", "u", "b", "s", "c", "i", "b", "e"}
	var strBuilder strings.Builder
	
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var catStr = strBuilder.String()
	fmt.Printf("\n%v", catStr)
}
