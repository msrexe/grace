package main

import (
	"grace"
	"log"
	"net/http"
)

func main() {
	handler := http.DefaultServeMux

	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	go func() {
		grace.Shutdown(func() {
			srv.Close()
			log.Println("Server gracefully stopped")
		})
	}()

	log.Printf("Server starting on port %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("%s\n", err)
	}
}
