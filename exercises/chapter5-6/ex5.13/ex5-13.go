package main

import (
	"ex5-13/links"
	"fmt"
	"log"
	"os"
)

func main() {
	breadthFirst(crawl, os.Args[1:])

	crawl(os.Args[1])
}

func crawl(url string) []string {
	fmt.Println(url)

	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	
	return list
}

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"net/url"
// 	"os"
// 	"path/filepath"
// 	"strings"

// 	"golang.org/x/net/html"
// )

// func main() {
// 	url := "https://go.dev"
// 	res := crawler(url)
// 	fmt.Println(res)
// }

// func crawler(link string) string {
// 	// 1. Download the page
// 	resp, err := http.Get(link)

// 	if err != nil {
// 		panic(err)
// 	}
// 	defer resp.Body.Close()

// 	// 2. Parse HTML
// 	doc, err := html.Parse(resp.Body)
// 	if err != nil {
// 		fmt.Println("Parse")
// 	}

// 	// // 3. Save the page
// 	savePage(link, doc)

// 	// 4. Extract links and follow them recursively

// 	for _, href := range extractLinks(doc, link) {		
// 		fmt.Println(href)

// 		if strings.HasPrefix(href, "https://golang.org") {
// 			// fmt.Println(href)
			
// 			crawler(href)
// 		}
// 	}
// 	return ""
// }

// func extractLinks(doc *html.Node, base string) []string {
// 	var links []string
// 	return visit(links, doc, base)
// }

// func visit(links []string, n *html.Node, base string) []string {
// 	if n.Type == html.ElementNode && n.Data == "a" {
// 		for _, a := range n.Attr {
// 			if a.Key == "href" {
// 				full := resolveURL(base, a.Val)
// 				if full != "" {
// 					links = append(links, full)
// 				}
// 			}
// 		}
// 	}
// 	for c := n.FirstChild; c != nil; c = c.NextSibling {
// 		links = visit(links, c, base)
// 	}
// 	return links
// }

// func resolveURL(base, href string) string {
// 	baseURL, err := url.Parse(base)
// 	if err != nil {
// 		return ""
// 	}
// 	hrefURL, err := url.Parse(href)
// 	if err != nil {
// 		return ""
// 	}
// 	return baseURL.ResolveReference(hrefURL).String()
// }



// func savePage(link string, doc *html.Node) {
// 	u, err := url.Parse(link)

// 	if err != nil {
// 		fmt.Println("Bad URL:", err)
// 		return
// 	}

// 	filePath := filepath.Join("copied", u.Host, u.Path)
// 	if strings.HasSuffix(filePath, "/") || filepath.Ext(filePath) == "" {
// 		filePath = filepath.Join(filePath, "index.html")
// 	}

// 	err = os.MkdirAll(filepath.Dir(filePath), 0755)
// 	if err != nil {
// 		fmt.Println("mkdir errors", err)
// 		return
// 	}

// 	f, err := os.Create(filePath)
// 	if err != nil {
// 		fmt.Println("file create error:", err)
// 		return
// 	}
// 	defer f.Close()

// 	err = html.Render(f, doc)
// 	if err != nil {
// 		fmt.Println("render error:", err)
// 	}
// }

