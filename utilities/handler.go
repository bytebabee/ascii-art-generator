package asciiart

import (
	"fmt"
	"log"
	"net/http"
)

func AsciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()
	text := r.FormValue("text")
	font := r.FormValue("font")

	if text == "" {
		http.Error(w, "Input text cannot be empty.", http.StatusBadRequest)
		return
	}
	if font == "" {
		font = "standard"
	}

	if !IsValidASCII(text) {
		http.Error(w, "Input contains unsupported characters.", http.StatusBadRequest)
		return
	}

	result, err := GenerateAsciiArt(text, font)
	if err != nil {
		log.Printf("Error generating ASCII art: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(result))
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		http.ServeFile(w, r, "./templates/index.html")
	case "/about.html":
		http.ServeFile(w, r, "./templates/about.html")
	default:
		// Serve 404 for any undefined route
		log.Printf("404 Page Not Found: URL=%s, RemoteAddr=%s", r.URL.Path, r.RemoteAddr)
		http.ServeFile(w, r, "./templates/404.html")
	}
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/about.html")
}

// Check if the input text contains only valid ASCII characters
func IsValidASCII(text string) bool {
	for _, char := range text {
		if char > 127 {
			return false
		}
	}
	return true
}

// Middleware for error handling
func ErrorHandler(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				log.Printf("Recovered from panic: %v", rec)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		handler(w, r)
	}
}

// ExportHandler handles the export functionality
func ExportHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	text := r.FormValue("text")
	font := r.FormValue("font")

	if text == "" {
		http.Error(w, "Input text cannot be empty", http.StatusBadRequest)
		return
	}

	// Generate ASCII art
	result, err := GenerateAsciiArt(text, font)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set headers for file download
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Disposition", "attachment; filename=ascii-art.txt")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(result)))

	// Write the content to the response
	_, err = w.Write([]byte(result))
	if err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
		return
	}
}
