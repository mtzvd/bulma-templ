package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

type Component struct {
	Name       string             `json:"name"`
	FilePath   string             `json:"file_path"`
	FolderPath string             `json:"folder_path"`
	Relations  ComponentRelations `json:"relations"`
}

type ComponentRelations struct {
	IsChildOf  string   `json:"is_child_of,omitempty"`
	IsParentOf []string `json:"is_parent_of"`
}

type Output struct {
	Metadata   Metadata    `json:"metadata"`
	Components []Component `json:"components"`
}

type Metadata struct {
	Generator string `json:"generator"`
	Timestamp string `json:"timestamp"`
}

// componentNamePattern matches templ component definitions
var componentNamePattern = regexp.MustCompile(`^templ\s+([A-Z][a-zA-Z0-9]*)\s*\(`)

func main() {
	projectRoot := "."
	outputPath := filepath.Join(projectRoot, "docs", "_generated", "components.json")

	// Discover all components
	components, err := discoverComponents(projectRoot)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error discovering components: %v\n", err)
		os.Exit(1)
	}

	// Infer relations
	inferRelations(components)

	// Generate output
	output := Output{
		Metadata: Metadata{
			Generator: "component-extractor-v1",
			Timestamp: time.Now().UTC().Format(time.RFC3339),
		},
		Components: components,
	}

	// Ensure output directory exists
	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating output directory: %v\n", err)
		os.Exit(1)
	}

	// Write output
	data, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	if err := os.WriteFile(outputPath, data, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing output file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Generated %s\n", outputPath)
	fmt.Printf("Found %d components\n", len(components))
}

// discoverComponents scans the project for .templ files and extracts component definitions
func discoverComponents(projectRoot string) ([]Component, error) {
	componentsMap := make(map[string]Component)

	// Component folders to scan
	folders := []string{"components", "elements", "form", "layout", "columns", "grid"}

	for _, folder := range folders {
		folderPath := filepath.Join(projectRoot, folder)

		entries, err := os.ReadDir(folderPath)
		if err != nil {
			if os.IsNotExist(err) {
				continue // folder doesn't exist, skip
			}
			return nil, err
		}

		for _, entry := range entries {
			if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".templ") {
				continue
			}

			filePath := filepath.Join(folderPath, entry.Name())

			// Read file and extract component names
			content, err := os.ReadFile(filePath)
			if err != nil {
				return nil, err
			}

			lines := strings.Split(string(content), "\n")
			for _, line := range lines {
				matches := componentNamePattern.FindStringSubmatch(strings.TrimSpace(line))
				if len(matches) > 1 {
					componentName := matches[1]

					// Only include public components (capitalized)
					if componentName[0] >= 'A' && componentName[0] <= 'Z' {
						// Normalize paths to forward slashes
						normalizedFilePath := filepath.ToSlash(filepath.Join(folder, entry.Name()))
						normalizedFolderPath := folder

						componentsMap[componentName] = Component{
							Name:       componentName,
							FilePath:   normalizedFilePath,
							FolderPath: normalizedFolderPath,
							Relations: ComponentRelations{
								IsParentOf: []string{},
							},
						}
					}
				}
			}
		}
	}

	// Convert map to sorted slice
	components := make([]Component, 0, len(componentsMap))
	for _, comp := range componentsMap {
		components = append(components, comp)
	}

	sort.Slice(components, func(i, j int) bool {
		return components[i].Name < components[j].Name
	})

	return components, nil
}

// inferRelations determines parent-child relationships based on naming conventions
func inferRelations(components []Component) {
	// Build a name index
	nameIndex := make(map[string]int)
	for i, comp := range components {
		nameIndex[comp.Name] = i
	}

	// For each component, determine if it is a child of another component
	for i := range components {
		comp := &components[i]

		// Check if this component's name starts with another component's name
		// A component B is a child of component A if B's name starts with A's name
		// and A exists as a component
		for parentName, parentIdx := range nameIndex {
			if parentIdx == i {
				continue // skip self
			}

			// Check if comp.Name starts with parentName
			if strings.HasPrefix(comp.Name, parentName) && len(comp.Name) > len(parentName) {
				// Additional check: the character after the prefix should indicate continuation
				// This ensures "NavbarItem" is child of "Navbar", but "Button" is not child of "Box"
				// We accept any case where the name starts with parent name
				comp.Relations.IsChildOf = parentName

				// Add this component to the parent's IsParentOf list
				components[parentIdx].Relations.IsParentOf = append(
					components[parentIdx].Relations.IsParentOf,
					comp.Name,
				)

				break // A component can only have one parent
			}
		}
	}

	// Sort IsParentOf lists
	for i := range components {
		sort.Strings(components[i].Relations.IsParentOf)
	}
}
