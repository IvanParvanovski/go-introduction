package main 

import (
	"fmt"
	"os"
	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	
	if err != nil {
		fmt.Fprintf(os.Stderr, "didn't manage to read html, %v", err)
		os.Exit(1)
	}

	// printTags(doc)
	// fmt.Println()
	// printTagsNew(doc)
	// html.Render(os.Stdout, doc)
	// visitRecursive(nil, doc)

	for _, link := range visitRecursive(nil, doc) {
		fmt.Println(link)
	}

}

// func printTags(n *html.Node) {
// 	if n.Type == html.ElementNode {
// 		fmt.Println("Tag:", n.Data)
// 	}

// 	for c := n.FirstChild; c != nil; c = c.NextSibling {
// 		printTags(c)
// 	}
// }

// func printTagsNew(n *html.Node) {
// 	if n == nil {
// 		return
// 	}

// 	if n.Type == html.ElementNode {
// 		fmt.Println("Tag:", n.Data)
// 	}

// 	// Visit first child
// 	printTagsNew(n.FirstChild)

// 	// Visit next sibling
// 	printTagsNew(n.NextSibling)
// }

// func visit(links []string, n *html.Node) []string {
// 	if n.Type == html.ElementNode && n.Data == "a" {
// 		for _, a := range n.Attr {
// 			if a.Key == "href" {
// 				links = append(links, a.Val)
// 			}
// 		}
// 	}

// 	for c := n.FirstChild; c != nil; c = c.NextSibling {
// 		links = visit(links, c)
// 	}
// 	return links
// }

func visitRecursive(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			fmt.Println(links)
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	if n.Type == html.ElementNode {
		fmt.Println(links)

	}

	links = visitRecursive(links, n.FirstChild)
	links = visitRecursive(links, n.NextSibling)

	return links
}