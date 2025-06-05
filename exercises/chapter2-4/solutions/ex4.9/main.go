package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <file>\n", os.Args[0])
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "error opening file:", err)
		os.Exit(1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords) // break the input into words

	freq := make(map[string]int)

	for scanner.Scan() {
		word := normalize(scanner.Text())
		if word == "" {
			continue
		}
		freq[word]++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "scan error:", err)
		os.Exit(1)
	}

	type pair struct {
		word  string
		count int
	}
	var list []pair
	for w, c := range freq {
		list = append(list, pair{w, c})
	}

	sort.Slice(list, func(i, j int) bool {
		if list[i].count == list[j].count {
			return list[i].word < list[j].word
		}
		return list[i].count > list[j].count
	})

	for _, p := range list {
		fmt.Printf("%s\t%d\n", p.word, p.count)
	}
}

func normalize(s string) string {
	s = strings.ToLower(s)
	return strings.Trim(s, ".,;:!?\"'()[]{}<>")
}
