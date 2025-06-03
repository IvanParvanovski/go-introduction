package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"regexp"
)

func main() {
	dictionary := map[string]int{}
	fileName := "file.txt"

	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	scanner.Split(bufio.ScanWords)
	re := regexp.MustCompile(`[a-zA-Z0-9]+`)

	for scanner.Scan() {
		rawWord := scanner.Text()

		cleanedWord := re.FindString(rawWord)
		if cleanedWord == "" {
			continue
		}

		cleanedWord = strings.ToLower(cleanedWord)
		dictionary[cleanedWord]++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	for word, count := range dictionary {
		fmt.Printf("%s: %d\n", word, count)
	}
}



