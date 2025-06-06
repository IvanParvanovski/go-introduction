package main

import (
	"fmt"
	"strings"
	// "io/ioutil"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		hasHttp := strings.HasPrefix(url, "http")
		
		if (!hasHttp) {
			url = "http://" + url
		}
		
		resp, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		// b, err := ioutil.ReadAll(resp.Body)
		// io.Copy(dst, src)
		io.Copy(os.Stdout, resp.Body)

		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
