package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var url = "https://xkcd.com/571/info.0.json"

type Response struct {
	Img string `json:"img"`
}

func main() {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("failed to fetch data: ", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("failed to read resp: ", err)
		return
	}

	var parsed Response
	err = json.Unmarshal(body, &parsed)
	if err != nil {
		log.Println("can't unmarshal json: ", err)
		return
	}

	spl := strings.Split(parsed.Img, "/")
	filename := spl[len(spl)-1]

	f, err := os.Create(filename)
	if err != nil {
		log.Println("can't create image file: ", err)
		return
	}

	resp, err = http.Get(parsed.Img)
	if err != nil {
		log.Println("failed to fetch image: ", err)
		return
	}
	defer resp.Body.Close()

	io.Copy(f, resp.Body)
}
