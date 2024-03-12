package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// List of allowed file extensions (lowercase)
var allowedExtensions = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".gif":  true,
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		// Get the file from form data
		file, handler, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Could not get file from form data", http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Check file extension
		ext := strings.ToLower(filepath.Ext(handler.Filename))
		if !allowedExtensions[ext] {
			http.Error(w, "File type not allowed", http.StatusBadRequest)
			return
		}

		uploadDir := "./uploads"
		os.MkdirAll(uploadDir, os.ModePerm)
		filePath := filepath.Join(uploadDir, handler.Filename)

		// Opening a new file to store data from the uploaded file.
		newFile, err := os.Create(filePath)
		if err != nil {
			http.Error(w, "Could not create file", http.StatusInternalServerError)
			return
		}
		defer newFile.Close()

		// Copy data from the uploaded file to the new file.
		_, err = io.Copy(newFile, file)
		if err != nil {
			http.Error(w, "Could not copy file data", http.StatusInternalServerError)
			return
		}

		// fmt.Fprintf(w, "File uploaded successfully")
		respondWithJSON(w, http.StatusOK, map[string]string{"message": "File uploaded successfully"})

	} else {
		// http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		respondWithError(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Helper function to send an error message as JSON
func respondWithError(w http.ResponseWriter, message string, statusCode int) {
	respondWithJSON(w, statusCode, map[string]string{"error": message})
}

// Helper function to send a response as JSON
func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error marshalling the response data"}`))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func main() {
	http.HandleFunc("/upload", uploadHandler)
	fmt.Println("Server started at :1010")
	http.ListenAndServe(":1010", nil)
}
