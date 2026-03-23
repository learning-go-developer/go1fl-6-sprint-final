package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
	"github.com/go-chi/chi/v5"
)

// Server represents the application server, encapsulating the custom logger
// and the underlying HTTP server instance.
type Server struct {
	Log     *log.Logger
	HttpSrv *http.Server
}

// NewServer initializes a new HTTP router, registers application handlers,
// and returns a configured Server instance with predefined timeouts.
// It sets the server to listen on port 8080 and configures Read, Write,
// and Idle timeouts to ensure resource efficiency and security.
func NewServer(logger *log.Logger) *Server {
	r := chi.NewRouter()

	r.Get("/", handlers.GetHTML)
	r.Post("/upload", handlers.UploadHandler)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		Log:     logger,
		HttpSrv: srv,
	}
}
