package main 

import "net/http"

func main() {
	http.HandleFunc("/", indexView)
	
	err := http.ListenAndServe(":8080", nil)
	
	if err != nil {
		panic(err.Error())
	}
}

func indexView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Go!"))
	
	
}
