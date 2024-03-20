package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	sf := http.FileServer(http.Dir("public"))
	mux.Handle("/", sf)
	server := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	e := server.ListenAndServe()
	if e != nil {
		log.Fatal(e)
	}
}

func message(w http.ResponseWriter, r *http.Request) {
	_, e := fmt.Fprintf(w, "Hello, world!")
	if e != nil {
		log.Fatal(e)
	}
}
