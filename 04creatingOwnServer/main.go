package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloWorld)
	server := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second, // Time the server has to wait to Read.
		WriteTimeout:   10 * time.Second, // Time the server has to wait to Write.
		MaxHeaderBytes: 1 << 20,          // Max num of header (1mb)
	}
	log.Println("Listening...")
	e := server.ListenAndServe()
	if e != nil {
		log.Fatal(e)
	}
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	_, e := fmt.Fprintf(w, "Hello, world!")
	if e != nil {
		log.Fatal(e)
	}
}
