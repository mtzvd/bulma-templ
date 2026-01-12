// Package docserver — Development server for Bulma-Templ examples and documentation.
package docserver

import "os"

// Config holds the configuration for the documentation server.
type Config struct {
	// Port is the server port (defaults to 8080).
	Port string

	// StarterPath is the route path for the starter example.
	StarterPath string

	// KitchenPath is the route path for the kitchen sink example.
	KitchenPath string

	// DocsPath is the route path for the documentation browser.
	DocsPath string
}

// DefaultConfig returns a Config with default values.
func DefaultConfig() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return Config{
		Port:        port,
		StarterPath: "/",
		KitchenPath: "/kitchensink",
		DocsPath:    "/knowledgebase/",
	}
}
