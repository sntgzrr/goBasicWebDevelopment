package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/users", GetUsers).Methods("GET")
	r.HandleFunc("/api/users", PostUsers).Methods("POST")
	r.HandleFunc("/api/users", PutUsers).Methods("PUT")
	r.HandleFunc("/api/users", DeleteUsers).Methods("DELETE")
	server := http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	e := server.ListenAndServe()
	if e != nil {
		log.Fatal(e)
	}
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	_, e := fmt.Fprintf(w, "Mensaje desde Get")
	if e != nil {
		log.Fatal(e)
	}
}

func PostUsers(w http.ResponseWriter, r *http.Request) {
	_, e := fmt.Fprintf(w, "Mensaje desde Post")
	if e != nil {
		log.Fatal(e)
	}
}

func PutUsers(w http.ResponseWriter, r *http.Request) {
	_, e := fmt.Fprintf(w, "Mensaje desde Put")
	if e != nil {
		log.Fatal(e)
	}
}

func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	_, e := fmt.Fprintf(w, "Mensaje desde Delete")
	if e != nil {
		log.Fatal(e)
	}
}
