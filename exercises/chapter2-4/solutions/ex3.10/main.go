package main

import "fmt"

func main() {
	tests := []string{"1", "12", "123", "1234", "12345", "123456", "1234567"}

	for _, t := range tests {
		fmt.Println(commaRecursive(t))
		fmt.Println(comma(t))
	}
}

func commaRecursive(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	return commaRecursive(s[:n-3]) + "," + s[n-3:]
}

func comma(s string) string {
	n := len(s)
	out := ""
	three := 0

	for i := n - 1; i >= 0; i-- {
		if three < 3 {
			out = string(s[i]) + out
			three++
			continue
		}

		out = string(s[i]) + "," + out
		three = 1
	}

	return out
}
