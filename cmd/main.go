package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

// main initializes and runs the HTTP server.
// It creates a new logger that writes to standard output,
// builds a server instance using the server package,
// and starts listening for incoming requests.
// If the server fails to start, the application exits with a fatal error.
func main() {
	logger := log.New(os.Stdout, "SERVER: ", log.LstdFlags)

	srv := server.NewServer(logger)

	logger.Println("Starting server on :8080...")

	if err := srv.HttpSrv.ListenAndServe(); err != nil {
		logger.Fatal("Server failed to start: ", err)
	}
}
