package main 

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	fmt.Println("hi")
	doc, err := html.Parse(os.Stdin)

	if err != nil {
		fmt.Fprintf(os.Stderr, "didn't manage to read html, %v", err)
		os.Exit(1)
	}

	for _, r := range visit(nil, doc) {
		fmt.Println(r)
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		switch n.Data {
		case "a", "link":
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
				}
			}
		case "img", "script":
			for _, a := range n.Attr {
				if a.Key == "src" {
					links = append(links, a.Val)
				}
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}