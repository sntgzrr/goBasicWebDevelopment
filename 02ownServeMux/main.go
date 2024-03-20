package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloWord)
	e := http.ListenAndServe(":8080", mux)
	if e != nil {
		fmt.Println(e)
	}
}
func helloWord(w http.ResponseWriter, r *http.Request) {
	_, e := fmt.Fprintf(w, "Hello, world!")
	if e != nil {
		fmt.Println(e)
	}
}
