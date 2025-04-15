package main

import (
	"log"
	"net/http"

	"k8s-mcp-server/routes"
)

func main() {
	// Initialize the router
	r := routes.InitializeRouter()

	// Start the server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
