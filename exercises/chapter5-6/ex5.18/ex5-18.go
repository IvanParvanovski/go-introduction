package main

import (
	"io"
	"net/http"
	"os"
	"path"
)

func main() {

}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)

	if local == "/" {
		local = "index.html"
	}

	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	defer f.Close()

	n, err = io.Copy(f, resp.Body)
	return local, n, err
}
