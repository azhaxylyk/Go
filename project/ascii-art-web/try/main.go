package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/ascii-art", Output)
	log.Println("start server in http://localhost:1337")
	err := http.ListenAndServe(":1337", mux)
	if err != nil {
		log.Fatal(err)
	}
}
