package main 

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	dictionary := map[string]int{}

	doc, err := html.Parse(os.Stdin)

	if err != nil {
		fmt.Fprintf(os.Stderr, "didn't manage to read html, %v", err)
		os.Exit(1)
	}

	printTagsNew(dictionary, doc)

	for k, v := range dictionary {
		fmt.Printf("%s -> %d\n", k, v)
	}
}

func printTagsNew(dictionary map[string]int, n *html.Node) {
	if n == nil {
		return
	}

	if n.Type == html.ElementNode {
		dictionary[n.Data]++
	}

	// Visit first child
	printTagsNew(dictionary, n.FirstChild)

	// Visit next sibling
	printTagsNew(dictionary, n.NextSibling)
}