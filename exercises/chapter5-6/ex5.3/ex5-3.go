package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)


func main() {
	doc, err := html.Parse(os.Stdin)

	if err != nil {
		fmt.Fprintf(os.Stderr, "didn't manage to read html, %v", err)
		os.Exit(1)
	}

	printTagsNew(doc)
}


func printTagsNew(n *html.Node) {
	if n == nil {
		return
	}

	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
		return
	}

	if n.Type == html.TextNode && strings.TrimSpace(n.Data) != "" {
		if n.Data != "" {
			fmt.Println(n.Data)
		}
	}

	printTagsNew(n.FirstChild)
	printTagsNew(n.NextSibling)

}