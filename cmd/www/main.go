package main

import (
	"log"
	"net/http"

	"github.com/bmizerany/pat"
)

func main() {
	// Initialize a router and add the path and handler for the homepage.
	mux := pat.New()
	mux.Get("/:locale", http.HandlerFunc(handleHome))

	// Start the HTTP server using the router.
	log.Println("starting server on :8080...")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
