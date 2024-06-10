package main

import (
	"Go/project/ascii-art-web/ascii-art-web-dockerize/ascii-art/convert"
	"Go/project/ascii-art-web/ascii-art-web-dockerize/function"
	"fmt"
	"log"
	"net/http"
)

func main() {
	if err := convert.Initialize_Format(); err != nil { // initializes the array of styles
		fmt.Println("Function name: main")
		return
	}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../function/static"))))
	http.HandleFunc("/", function.HomeHandler)
	http.HandleFunc("/ascii-art", function.AsciiArtHandler)
	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
