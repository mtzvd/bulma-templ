package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Component represents a parsed Bulma-Templ component
type Component struct {
	Name        string
	Description string
	AtomicLevel string
	Package     string
	Props       []PropField
	BulmaClass  string
	SourceFile  string
	FullComment string
}

// PropField represents a single prop in a Props struct
type PropField struct {
	Name    string
	Type    string
	Comment string
}

func main() {
	var (
		inputFile  = flag.String("file", "", "Path to .templ file to parse")
		outputDir  = flag.String("output", "docs/knowledgebase", "Output directory for generated docs")
		packageDir = flag.String("package", "", "Parse all .templ files in package directory")
	)
	flag.Parse()

	if *inputFile == "" && *packageDir == "" {
		fmt.Println("Usage: docgen -file <file.templ> [-output <dir>]")
		fmt.Println("   or: docgen -package <dir> [-output <dir>]")
		os.Exit(1)
	}

	if *inputFile != "" {
		// Parse single file (may contain multiple components)
		components, err := parseTemplFile(*inputFile)
		if err != nil {
			fmt.Printf("Error parsing file: %v\n", err)
			os.Exit(1)
		}

		for _, component := range components {
			markdown := generateMarkdown(&component)

			// Create output directory
			pkgDir := filepath.Join(*outputDir, "components", component.Package)
			if err := os.MkdirAll(pkgDir, 0755); err != nil {
				fmt.Printf("Error creating output directory: %v\n", err)
				os.Exit(1)
			}

			// Write output file
			outputFile := filepath.Join(pkgDir, strings.ToLower(component.Name)+".md")
			if err := os.WriteFile(outputFile, []byte(markdown), 0644); err != nil {
				fmt.Printf("Error writing output file: %v\n", err)
				os.Exit(1)
			}

			fmt.Printf("✓ Generated: %s\n", outputFile)
		}
	}

	if *packageDir != "" {
		// Parse all .templ files in directory
		files, err := filepath.Glob(filepath.Join(*packageDir, "*.templ"))
		if err != nil {
			fmt.Printf("Error finding .templ files: %v\n", err)
			os.Exit(1)
		}

		count := 0
		for _, file := range files {
			// Skip elements.templ (base infrastructure)
			if strings.HasSuffix(file, "elements.templ") || strings.HasSuffix(file, "html.templ") {
				continue
			}

			components, err := parseTemplFile(file)
			if err != nil {
				fmt.Printf("Warning: Failed to parse %s: %v\n", file, err)
				continue
			}

			for _, component := range components {
				markdown := generateMarkdown(&component)

				pkgDir := filepath.Join(*outputDir, "components", component.Package)
				if err := os.MkdirAll(pkgDir, 0755); err != nil {
					fmt.Printf("Error creating output directory: %v\n", err)
					continue
				}

				outputFile := filepath.Join(pkgDir, strings.ToLower(component.Name)+".md")
				if err := os.WriteFile(outputFile, []byte(markdown), 0644); err != nil {
					fmt.Printf("Error writing %s: %v\n", outputFile, err)
					continue
				}

				fmt.Printf("✓ Generated: %s\n", outputFile)
				count++
			}
		}

		fmt.Printf("\n✓ Generated %d component docs\n", count)
	}
}

// parseTemplFile parses a .templ file and extracts ALL component metadata
func parseTemplFile(path string) ([]Component, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var components []Component
	var currentComponent *Component

	scanner := bufio.NewScanner(file)
	var currentComment []string
	var inPropsStruct bool

	// Extract package name from path
	packageName := filepath.Base(filepath.Dir(path))
	sourceFile := filepath.Base(path)

	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)

		// Collect comments
		if strings.HasPrefix(trimmed, "//") {
			comment := strings.TrimPrefix(trimmed, "//")
			comment = strings.TrimSpace(comment)
			currentComment = append(currentComment, comment)
			continue
		}

		// Parse type declarations (Props structs)
		if strings.HasPrefix(trimmed, "type ") && strings.Contains(trimmed, "Props struct") {
			// Save previous component if exists
			if currentComponent != nil && currentComponent.Name != "" {
				components = append(components, *currentComponent)
			}

			// Start new component
			currentComponent = &Component{
				Package:    packageName,
				SourceFile: sourceFile,
				Props:      []PropField{},
			}

			// Extract component name from Props struct name
			re := regexp.MustCompile(`type (\w+)Props struct`)
			matches := re.FindStringSubmatch(trimmed)
			if len(matches) > 1 {
				currentComponent.Name = matches[1]

				// Parse component-level comments
				currentComponent.FullComment = strings.Join(currentComment, "\n")
				parseComponentMetadata(currentComponent, currentComment)
				inPropsStruct = true
				currentComment = nil
			}
			continue
		}

		// Parse struct fields
		if inPropsStruct && currentComponent != nil {
			if trimmed == "}" {
				inPropsStruct = false
				currentComment = nil
				continue
			}

			if trimmed != "" {
				// Parse field line
				field := parseStructField(trimmed, currentComment)
				if field != nil {
					currentComponent.Props = append(currentComponent.Props, *field)
				}
				currentComment = nil
			}
			continue
		}

		// Parse templ function declarations (optional, for validation)
		if strings.HasPrefix(trimmed, "templ ") && currentComponent != nil {
			re := regexp.MustCompile(`templ (\w+)\(`)
			matches := re.FindStringSubmatch(trimmed)
			if len(matches) > 1 {
				templName := matches[1]
				// Validate that templ name matches component name
				if templName == currentComponent.Name {
					// Extract Bulma class
					currentComponent.BulmaClass = "." + camelToKebab(currentComponent.Name)
				}
			}
			currentComment = nil
			continue
		}

		// Reset comments if we hit a non-comment, non-target line
		if trimmed != "" && !strings.HasPrefix(trimmed, "package") && !strings.HasPrefix(trimmed, "import") {
			if !inPropsStruct {
				currentComment = nil
			}
		}
	}

	// Save last component
	if currentComponent != nil && currentComponent.Name != "" {
		components = append(components, *currentComponent)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// Set Bulma class for all components
	for i := range components {
		if components[i].BulmaClass == "" {
			components[i].BulmaClass = "." + camelToKebab(components[i].Name)
		}
	}

	return components, nil
}

// parseComponentMetadata extracts metadata from component comments
func parseComponentMetadata(component *Component, comments []string) {
	var descriptionLines []string
	foundTitle := false
	foundAtomicLevel := false

	for _, comment := range comments {
		// Extract title and description (first line with em dash)
		if strings.Contains(comment, "—") && !foundTitle {
			parts := strings.Split(comment, "—")
			if len(parts) == 2 {
				// Don't override component name, it's from Props struct
				component.Description = strings.TrimSpace(parts[1])
				foundTitle = true
			}
			continue
		}

		// Extract atomic level
		if strings.HasPrefix(comment, "Atomic level:") {
			level := strings.TrimPrefix(comment, "Atomic level:")
			component.AtomicLevel = strings.TrimSpace(level)
			foundAtomicLevel = true
			continue
		}

		// Collect description lines (after atomic level, non-empty)
		if foundAtomicLevel && comment != "" {
			descriptionLines = append(descriptionLines, comment)
		}
	}

	// Add detailed description if found
	if len(descriptionLines) > 0 {
		component.FullComment = strings.Join(descriptionLines, "\n\n")
	}
}

// parseStructField parses a struct field line with its comment
func parseStructField(line string, comments []string) *PropField {
	// Skip empty lines and closing braces
	if line == "" || line == "}" {
		return nil
	}

	// Parse field: FieldName Type
	parts := strings.Fields(line)
	if len(parts) < 2 {
		return nil
	}

	field := &PropField{
		Name: parts[0],
		Type: parts[1],
	}

	// Combine multi-line comments
	if len(comments) > 0 {
		field.Comment = strings.Join(comments, " ")
	}

	return field
}

// generateMarkdown generates Markdown documentation for a component
func generateMarkdown(component *Component) string {
	var sb strings.Builder

	// Frontmatter
	sb.WriteString("---\n")
	sb.WriteString(fmt.Sprintf("title: \"%s\"\n", component.Name))
	sb.WriteString(fmt.Sprintf("description: \"%s\"\n", escapeYAML(component.Description)))
	sb.WriteString(fmt.Sprintf("category: \"%s\"\n", capitalize(component.Package)))
	sb.WriteString(fmt.Sprintf("atomicLevel: \"%s\"\n", component.AtomicLevel))
	sb.WriteString(fmt.Sprintf("bulmaClass: \"%s\"\n", component.BulmaClass))
	sb.WriteString("---\n\n")

	// Title
	sb.WriteString(fmt.Sprintf("# %s\n\n", component.Name))

	// Badge for atomic level
	if component.AtomicLevel != "" {
		sb.WriteString(fmt.Sprintf("**Atomic Level:** `%s`\n\n", component.AtomicLevel))
	}

	// Description
	if component.Description != "" {
		sb.WriteString(fmt.Sprintf("%s\n\n", component.Description))
	}

	// Full description from comments
	if component.FullComment != "" {
		sb.WriteString("## Description\n\n")
		sb.WriteString(component.FullComment)
		sb.WriteString("\n\n")
	}

	// Props table
	if len(component.Props) > 0 {
		sb.WriteString("## Props\n\n")
		sb.WriteString("| Property | Type | Description |\n")
		sb.WriteString("|----------|------|-------------|\n")

		for _, prop := range component.Props {
			escapedComment := strings.ReplaceAll(prop.Comment, "|", "\\|")
			escapedComment = strings.ReplaceAll(escapedComment, "\n", " ")
			sb.WriteString(fmt.Sprintf("| `%s` | `%s` | %s |\n",
				prop.Name,
				prop.Type,
				escapedComment,
			))
		}
		sb.WriteString("\n")
	}

	// Usage section
	sb.WriteString("## Usage\n\n")
	sb.WriteString("```go\n")
	sb.WriteString(fmt.Sprintf("@%s.%s(\n", component.Package, component.Name))
	sb.WriteString(fmt.Sprintf("    %s.%sProps{},\n", component.Package, component.Name))
	sb.WriteString(fmt.Sprintf("    %s.Items{%s.Html(\"...\")},\n", component.Package, component.Package))
	sb.WriteString(")\n")
	sb.WriteString("```\n\n")

	// Bulma reference
	sb.WriteString("## Bulma Reference\n\n")
	bulmaPath := strings.ToLower(component.Package)
	componentPath := camelToKebab(component.Name)
	sb.WriteString(fmt.Sprintf("[Bulma %s Documentation](https://bulma.io/documentation/%s/%s/)\n\n",
		component.Name,
		bulmaPath,
		componentPath,
	))

	// Source reference
	sb.WriteString("## Source\n\n")
	sb.WriteString(fmt.Sprintf("**File:** `%s/%s`\n", component.Package, component.SourceFile))

	return sb.String()
}

// Helper functions

func camelToKebab(s string) string {
	var result strings.Builder
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result.WriteRune('-')
		}
		result.WriteRune(r)
	}
	return strings.ToLower(result.String())
}

func capitalize(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func escapeYAML(s string) string {
	s = strings.ReplaceAll(s, "\"", "\\\"")
	return s
}
