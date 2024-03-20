package main

import (
	"fmt"
	"net/http"
)

type message struct {
	msg string
}

func (m message) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, e := fmt.Fprintf(w, m.msg)
	if e != nil {
		fmt.Println(e)
	}
}

func main() {
	msg := message{
		msg: "Hello, world!",
	}
	mux := http.NewServeMux()
	mux.Handle("/", msg)
	e := http.ListenAndServe(":8080", mux)
	if e != nil {
		fmt.Println(e)
	}
}
