package main

import (
	"net/http"
	"os"
	"path/filepath"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// Load values from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get values from environment variables or use default values
	dir := getEnv("MEDIA_DIR", "./media")
	addr := getEnv("SERVER_ADDR", "localhost:8080")

	// Create a new file server with the given directory
	fileServer := http.FileServer(http.Dir(dir))

	// Create a handler function that serves files
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Get the requested file path
		requestedPath := r.URL.Path
		_, exists := getFilePath(dir, requestedPath)

		// Check if the file exists
		if exists {
			// Serve the file using the file server
			http.StripPrefix("/", fileServer).ServeHTTP(w, r)
		} else {
			// File not found, return a 404 response
			http.NotFound(w, r)
		}
	})

	// Define the server address and start the server
	println("Server listening on", addr)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		panic(err)
	}
}

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}

func getFilePath(baseDir, requestedPath string) (string, bool) {
	// Create the absolute path by joining baseDir and requestedPath
	absolutePath := filepath.Join(baseDir, requestedPath)

	// Check if the file exists
	if fileExists(absolutePath) {
		return absolutePath, true
	}

	return "", false
}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}