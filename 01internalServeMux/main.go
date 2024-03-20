package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, e := fmt.Fprintf(w, "Hello, world!")
		if e != nil {
			fmt.Println(e)
		}
	})
	e := http.ListenAndServe(":8080", nil)
	if e != nil {
		fmt.Println(e)
	}
}
