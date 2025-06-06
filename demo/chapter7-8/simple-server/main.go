package main 

import (
	"net/http"
	"fmt"
)

func main() {
	mux := http.NewServeMux()

	fmt.Println("Server listening to :8080")
	mux.HandleFunc("/", handleRoot)
	http.ListenAndServe(":8080",  mux)
}

func handleRoot(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}
