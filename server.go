package main

import (
	"log"
	"net/http"
	"asciiart/utilities"
)

func main() {
	// Create a new ServeMux
	mux := http.NewServeMux()

	// Serve static files for styles, fonts, and images
	mux.Handle("/static/fonts/", http.StripPrefix("/static/fonts/", http.FileServer(http.Dir("static/fonts"))))
	mux.Handle("/static/images/", http.StripPrefix("/static/images/", http.FileServer(http.Dir("static/images"))))
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static")))) // For other static files like style.css

	// Define routes with error-handling middleware
	mux.HandleFunc("/generate", asciiart.ErrorHandler(asciiart.AsciiHandler))
	mux.HandleFunc("/about.html", asciiart.ErrorHandler(asciiart.AboutHandler))
	mux.HandleFunc("/", asciiart.ErrorHandler(asciiart.RootHandler))
	mux.HandleFunc("/export", asciiart.ErrorHandler(asciiart.ExportHandler))

	// Start the server
	log.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
