package main

import (
	"fmt"
	"sort"
)

func IsPalindrome(s sort.Interface) bool {
	for i := 0; i < s.Len() / 2; i++ {
		j := s.Len() - i - 1

		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}

	return true
}



func main() {
    numbers := sort.IntSlice{1, 2, 3, 2, 1}
	// numbers := sort.IntSlice{1, 5, 3, 4}
	fmt.Println(IsPalindrome(numbers)) 
}