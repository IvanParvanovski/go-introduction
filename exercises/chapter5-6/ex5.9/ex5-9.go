package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	s := "Welcome $foo $bar to ..."

	fmt.Println(expand(s, toUpperCase))
}

func expand(s string, f func(string) string) string {
	re := regexp.MustCompile(`\$(\w+)`)

	return re.ReplaceAllStringFunc(s, f)
}

func toUpperCase(s string) string {
	return strings.ToUpper(s[1:])
}
