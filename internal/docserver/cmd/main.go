// Package main — Development server entry point for Bulma-Templ.
//
// This server provides three main routes:
//   - / — Starter example page
//   - /kitchensink — Kitchen Sink example with all components
//   - /knowledgebase/ — Markdown documentation browser
package main

import (
	"log/slog"
	"os"

	"github.com/mtzvd/bulma-templ/internal/docserver"
)

func main() {
	// Initialize structured logger.
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)

	// Create server with default configuration.
	cfg := docserver.DefaultConfig()
	server := docserver.New(cfg)

	// Start the server.
	if err := server.Start(); err != nil {
		slog.Error("server failed", "error", err)
		os.Exit(1)
	}
}
