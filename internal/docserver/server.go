// Package docserver — Development server for Bulma-Templ examples and documentation.
//
// This package provides a development server that serves:
//   - Starter example page
//   - Kitchen Sink example with all components
//   - Markdown documentation browser
package docserver

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	"github.com/mtzvd/bulma-templ/examples/kitchensink"
	"github.com/mtzvd/bulma-templ/examples/starter"
	"github.com/mtzvd/bulma-templ/internal/docserver/internal/handlers"
)

// Server represents the documentation server.
type Server struct {
	config Config
	mux    *http.ServeMux
	logger *slog.Logger
}

// New creates a new documentation server with the given configuration.
func New(cfg Config) *Server {
	return &Server{
		config: cfg,
		mux:    http.NewServeMux(),
		logger: slog.Default(),
	}
}

// Start initializes routes and starts the HTTP server.
func (s *Server) Start() error {
	s.setupRoutes()
	s.printServerInfo()

	addr := ":" + s.config.Port
	s.logger.Info("starting server", "addr", addr)

	return http.ListenAndServe(addr, s.mux)
}

// setupRoutes configures all HTTP routes for the server.
func (s *Server) setupRoutes() {
	// Starter page (root).
	s.mux.Handle(s.config.StarterPath, templ.Handler(starter.Page()))

	// Kitchen Sink page (component gallery).
	s.mux.Handle(s.config.KitchenPath, templ.Handler(kitchensink.Page()))

	// Documentation browser.
	s.mux.HandleFunc(s.config.DocsPath, handlers.HandleDocs)
}

// printServerInfo logs server startup information.
func (s *Server) printServerInfo() {
	baseURL := fmt.Sprintf("http://localhost:%s", s.config.Port)
	s.logger.Info("development server starting",
		"url", baseURL,
		"port", s.config.Port,
	)
	s.logger.Info("available routes",
		"starter", baseURL+s.config.StarterPath,
		"kitchenSink", baseURL+s.config.KitchenPath,
		"docs", baseURL+s.config.DocsPath,
	)
}
