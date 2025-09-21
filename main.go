package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// build a simple server that binds localhost:8080 to a 404
func main() {
	port := 8080
	// create new http.ServeMux
	mux := http.NewServeMux()

	// create new http.Server
	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// use server's ListenAndServe
	fmt.Printf("Serving on port: %d\n", port)
	log.Fatal(s.ListenAndServe())
}
