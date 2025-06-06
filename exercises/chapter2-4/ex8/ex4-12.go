package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
	"encoding/json"
)

type XKCD struct {
	Img string `json:"img"`
}

func main() {
	url := "https://xkcd.com/571/info.0.json"

	// Read full json 
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return
	}

	var xkcd XKCD
	err = json.Unmarshal(body, &xkcd)

	fmt.Println(xkcd.Img)

	// Read image
	resp, err = http.Get(xkcd.Img)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// body, err = io.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("Error reading body:", err)
	// 	os.Exit(1)
	// }
	
	// fmt.Println(body)
	
	// Save image to file
	file, err := os.Create("./img.png")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	_, err = io.Copy(file, resp.Body)
	fmt.Println("Success!")
}