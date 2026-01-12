package handlers

import (
	"bytes"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/a-h/templ"
	"github.com/mtzvd/bulma-templ/internal/docserver/internal/web/pages"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

// md is the goldmark markdown parser configured with GitHub Flavored Markdown.
var md = goldmark.New(
	goldmark.WithExtensions(extension.GFM),
	goldmark.WithParserOptions(
		parser.WithAutoHeadingID(),
	),
	goldmark.WithRendererOptions(
		html.WithHardWraps(),
		html.WithUnsafe(),
	),
)

// HandleDocs serves markdown documentation files from the docs/knowledgebase directory.
func HandleDocs(w http.ResponseWriter, r *http.Request) {
	// Extract path from URL.
	path := strings.TrimPrefix(r.URL.Path, "/knowledgebase/")

	var mdPath string
	if path == "" || path == "/" {
		// Root knowledgebase - show README.
		mdPath = "docs/knowledgebase/README.md"
	} else {
		// Remove trailing .md if present.
		path = strings.TrimSuffix(path, ".md")
		// Build path directly from the URL path.
		mdPath = "docs/knowledgebase/" + path + ".md"
	}
	// Normalize path for Windows.
	mdPath = filepath.FromSlash(mdPath)

	// Read markdown file.
	content, err := os.ReadFile(mdPath)
	if err != nil {
		slog.Warn("documentation file not found",
			"path", mdPath,
			"requestPath", r.URL.Path,
			"error", err,
		)
		http.Error(w, "File not found: "+mdPath, http.StatusNotFound)
		return
	}

	// Strip YAML frontmatter if present.
	contentStr := stripFrontmatter(string(content))

	// Convert markdown to HTML.
	var buf bytes.Buffer
	if err := md.Convert([]byte(contentStr), &buf); err != nil {
		slog.Error("failed to render markdown",
			"path", mdPath,
			"requestPath", r.URL.Path,
			"error", err,
		)
		http.Error(w, "Error rendering markdown", http.StatusInternalServerError)
		return
	}

	// Extract title from path.
	title := extractTitle(path)

	// Render with templ.
	component := pages.DocsPage(title, buf.String())
	templ.Handler(component).ServeHTTP(w, r)
}

// stripFrontmatter removes YAML frontmatter from markdown content.
func stripFrontmatter(content string) string {
	if !strings.HasPrefix(content, "---\n") && !strings.HasPrefix(content, "---\r\n") {
		return content
	}

	// Find the closing ---.
	lines := strings.Split(strings.ReplaceAll(content, "\r\n", "\n"), "\n")
	for i := 1; i < len(lines); i++ {
		if lines[i] == "---" {
			// Skip all lines up to and including this closing ---.
			return strings.Join(lines[i+1:], "\n")
		}
	}

	// If no closing ---, use original content.
	return content
}

// extractTitle extracts a human-readable title from a file path.
func extractTitle(path string) string {
	title := strings.TrimSuffix(filepath.Base(path), ".md")
	if title == "" || title == "." {
		return "Documentation"
	}
	// Replace hyphens with spaces and title-case each word.
	title = strings.ReplaceAll(title, "-", " ")
	return toTitle(title)
}

// toTitle converts a string to title case (capitalizes first letter of each word).
// Replaces deprecated strings.Title.
func toTitle(s string) string {
	var result strings.Builder
	capitalizeNext := true

	for _, r := range s {
		if unicode.IsSpace(r) {
			result.WriteRune(r)
			capitalizeNext = true
			continue
		}

		if capitalizeNext {
			result.WriteRune(unicode.ToUpper(r))
			capitalizeNext = false
		} else {
			result.WriteRune(unicode.ToLower(r))
		}
	}

	return result.String()
}
