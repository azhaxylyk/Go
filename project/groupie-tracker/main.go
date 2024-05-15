package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"pkg"
	"time"
)

var (
	Mux *http.ServeMux
)

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	parentCtx := context.Background()

	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()

	pkg.UpdateCache()

	server := Server()
	go serverStart(ctx, server)

	<-stop
	cancel()
}

func serverStart(ctx context.Context, server *http.Server) {
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Error starting the server:", err)
	}
}

func Server() *http.Server {
	Mux = http.NewServeMux()
	Mux.HandleFunc("/", pkg.HomeHandler)
	Mux.HandleFunc("/band", pkg.BandHandler)
	fileServer := http.FileServer(http.Dir("web/static"))
	Mux.Handle("/web/static/", http.StripPrefix("/web/static/", fileServer))
	S := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      Mux,
	}
	fmt.Printf("Server: %s"+"\n", "http://localhost:8080")
	return S
}
