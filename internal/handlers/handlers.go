package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// GetHTML is an HTTP handler for the root endpoint.
// It serves the "index.html" file from the current working directory
// to the client. If the file is missing, it automatically returns
// a 404 Not Found error.
func GetHTML(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

// UploadHandler processes multipart form data to perform Morse code conversion.
// It performs the following steps:
// 1. Parses the multipart form with a 10MB limit.
// 2. Retrieves the file from the "myFile" form field.
// 3. Reads the file content and converts it using the service.DetectedMorse function.
// 4. Saves the conversion result to a local file named after the current UTC time.
// 5. Returns the converted string as a plain text response.
//
// If any step fails, it returns an HTTP 500 Internal Server Error.
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "form parsing error", http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "failed to get file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "error reading file", http.StatusInternalServerError)
		return
	}

	result, err := service.DetectedMorse(string(fileBytes))
	if err != nil {
		http.Error(w, "conversion error", http.StatusInternalServerError)
		return
	}

	ext := filepath.Ext(header.Filename)
	outFileName := fmt.Sprintf("%s%s", time.Now().UTC().String(), ext)

	outFile, err := os.Create(outFileName)
	if err != nil {
		http.Error(w, "failed to create file", http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	_, err = outFile.WriteString(result)
	if err != nil {
		http.Error(w, "recording error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(result))
}
