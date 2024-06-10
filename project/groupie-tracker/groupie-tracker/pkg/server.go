package pkg

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var Mux *http.ServeMux

func Server() *http.Server {

	Mux = http.NewServeMux()

	Mux.HandleFunc("/", HomeHandler)

	Mux.HandleFunc("/band", BandHandler)

	fileServer := http.FileServer(http.Dir("web/static"))

	Mux.Handle("/web/static/", http.StripPrefix("/web/static/", fileServer))

	S := &http.Server{
		Addr:         ":5000",           
		ReadTimeout:  30 * time.Second,  
		WriteTimeout: 90 * time.Second,  
		IdleTimeout:  120 * time.Second, 
		Handler:      Mux,               
	}

	fmt.Printf("Тестовый URL: %s"+"\n", "http://localhost:5000")
	fmt.Printf("Основной URL: %s"+"\n", "http://localhost:8080")
	go func() {
		err := S.ListenAndServe() 
		if err != nil {
			if err != http.ErrServerClosed {
				log.Fatal(err)
			}
		}
	}()
	return S
}
