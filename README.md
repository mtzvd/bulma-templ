# Bulma-Templ

[![Go Version](https://img.shields.io/badge/Go-1.25%2B-00ADD8?style=flat&logo=go)](https://go.dev)
[![Tests](https://github.com/mtzvd/bulma-templ/actions/workflows/test.yml/badge.svg)](https://github.com/mtzvd/bulma-templ/actions/workflows/test.yml)
[![Lint](https://github.com/mtzvd/bulma-templ/actions/workflows/lint.yml/badge.svg)](https://github.com/mtzvd/bulma-templ/actions/workflows/lint.yml)
[![Build](https://github.com/mtzvd/bulma-templ/actions/workflows/build.yml/badge.svg)](https://github.com/mtzvd/bulma-templ/actions/workflows/build.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Bulma](https://img.shields.io/badge/Bulma-1.0-00D1B2?style=flat&logo=bulma)](https://bulma.io)
[![Templ](https://img.shields.io/badge/Templ-latest-7C3AED?style=flat)](https://templ.guide)

A **type-safe**, **SSR-first** Bulma component library for Go using [Templ](https://templ.guide).

Provides 1:1 mappings of [Bulma CSS](https://bulma.io) components with explicit composition and zero internal state.

```go
@Button(
    ButtonProps{Color: "is-primary", Size: "is-large"},
    Items{Html("Get Started")},
)
```

## Features

- ✅ **Complete Bulma coverage** — All elements, components, forms, and layouts
- ✅ **Type-safe** — Props structs with IDE autocomplete
- ✅ **SSR-first** — No JavaScript required for rendering
- ✅ **Zero state** — Pure rendering functions
- ✅ **Explicit composition** — Multi-child `Items` pattern
- ✅ **Escape hatch** — `Attr` for custom attributes and Alpine.js
- ✅ **Stable v1.0** — Frozen architecture, no breaking changes

## Quick Start

### Installation

```bash
go get github.com/mtzvd/bulma-templ
```

That's it! Generated templ files are included in the repository.

### Requirements

- Go 1.25+
- Bulma CSS (via CDN or bundled)

### Basic Example

```go
package main

import "github.com/mtzvd/bulma-templ/elements"

templ Page() {
    <!DOCTYPE html>
    <html>
        <head>
            <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bulma@1.0.0/css/bulma.min.css">
        </head>
        <body>
            @elements.Button(
                elements.ButtonProps{Color: "is-primary"},
                elements.Items{elements.Html("Click Me")},
            )
        </body>
    </html>
}
```

## Examples

### Buttons

```go
// Primary button
@elements.Button(
    elements.ButtonProps{Color: "is-primary", Size: "is-large"},
    elements.Items{elements.Html("Primary")},
)

// With loading state
@elements.Button(
    elements.ButtonProps{Color: "is-info", IsLoading: true},
    elements.Items{elements.Html("Loading...")},
)
```

### Card

```go
@components.Card(
    components.CardProps{},
    components.Items{
        components.CardHeader(
            components.CardHeaderProps{},
            components.Items{
                components.CardHeaderTitle(
                    components.Items{elements.Html("Card Title")},
                ),
            },
        ),
        components.CardContent(
            components.Items{
                elements.Content(
                    elements.ContentProps{},
                    elements.Items{elements.Html("<p>Card content</p>")},
                ),
            },
        ),
    },
)
```

### Form

```go
@form.Field(
    form.FieldProps{},
    form.Items{
        form.Label(
            form.LabelProps{},
            form.Items{elements.Html("Email")},
        ),
        form.Control(
            form.ControlProps{},
            form.Items{
                form.Input(
                    form.InputProps{
                        Type: "email",
                        Placeholder: "user@example.com",
                    },
                ),
            },
        ),
    },
)
```

### Alpine.js Integration

```go
@elements.Button(
    elements.ButtonProps{
        Color: "is-primary",
        Attr: templ.Attributes{
            "x-on:click": "alert('Clicked!')",
        },
    },
    elements.Items{elements.Html("Interactive")},
)
```

## Philosophy

**Bulma-Templ is intentionally conservative.**

- **1:1 Bulma mapping** — No abstractions, no re-interpretation
- **Explicit composition** — Multi-child `Items` pattern
- **Zero state** — Components render HTML only
- **Type-safe** — Compile-time validation
- **Stable** — v1.0 architecture frozen

**What this is:**

- Type-safe Bulma components for Go/Templ
- Predictable HTML output
- Foundation for server-rendered UIs

**What this is NOT:**

- UI framework with routing/state
- JavaScript component library
- Theme engine or CSS abstraction

## Component API

Every component follows this pattern:

```go
templ ComponentName(props ComponentNameProps, content Items)
```

**Props struct** — Configuration via typed struct with Bulma modifiers  
**Items** — Multi-child composition model (`[]templ.Component`)  
**Attr** — Escape hatch for custom attributes (Alpine.js, ARIA, etc.)

Example Props:

```go
type ButtonProps struct {
    Color    string            // is-primary, is-link, etc.
    Size     string            // is-small, is-medium, is-large
    Disabled bool              // Disabled state
    Attr     templ.Attributes  // Escape hatch
}
```

## Available Components

### Elements

`block`, `box`, `button`, `content`, `delete`, `icon`, `image`, `notification`, `progress`, `skeleton`, `table`, `tag`, `title`

### Components

`breadcrumb`, `card`, `dropdown`, `menu`, `message`, `modal`, `navbar`, `pagination`, `panel`, `tabs`

### Form

`checkbox`, `control`, `field`, `file`, `help`, `input`, `label`, `radio`, `select`, `textarea`

### Layout

`container`, `footer`, `hero`, `level`, `media`, `section`

### Columns

`columns`, `column`

### Grid

`grid`

## Documentation

- [CANONICAL_PROJECT_CONTEXT.md](docs/CANONICAL_PROJECT_CONTEXT.md) — Architecture and design decisions
- [DESIGN_SYSTEM.md](docs/DESIGN_SYSTEM.md) — Component patterns and usage guidelines
- [CONTRIBUTING.md](CONTRIBUTING.md) — How to contribute
- [COMMENT_STYLE.md](docs/COMMENT_STYLE.md) — Code documentation standards
- [LLM_INSTRUCTIONS.md](docs/LLM_INSTRUCTIONS.md) — Guidelines for LLM-assisted development

## Development

### Quick Start

**Note:** `*_templ.go` files are generated and tracked in git for pkg.go.dev compatibility.

**For contributors:** If you modify `.templ` files, regenerate with `task templ` before committing.

```bash
task templ   # Generate Templ files (after modifying .templ files)
task run     # Run dev server
task test    # Run tests
```

### Testing

The project has **183 tests** with **80.9% coverage** for core packages, covering:

- **Infrastructure primitives** (Wrap, BaseElement, Items)
- **Components with conditional logic** (Progress, Navbar, Modal, Dropdown, Form inputs)
- **Smoke tests** for thin wrapper components

```bash
# Run all tests
go test ./...

# Run tests with race detection
go test ./... -race

# Run tests for specific package
go test ./elements -v
go test ./components -v
go test ./form -v
```

### CI/CD

All pull requests automatically run:

- ✅ Full test suite with race detection and coverage reporting
- ✅ `templ generate` verification (ensures generated files build correctly)
- ✅ Kitchen Sink and example compilation checks
- ✅ Multi-OS build verification (Linux, Windows, macOS)
- ✅ `go vet` and `gofmt` formatting checks
- ✅ golangci-lint static analysis

## Project Status

**Version 1.0** — Architecture is frozen. No breaking changes will be introduced.

All core Bulma components are implemented and stable. Future work focuses on documentation, examples, and community support.

## Contributing

Contributions are welcome! Please read [CONTRIBUTING.md](CONTRIBUTING.md) before submitting PRs.

**What's welcome:**

- Bug fixes
- Documentation improvements
- Test coverage
- Examples
- Accessibility enhancements

**What's out of scope:**

- New components (beyond Bulma)
- Framework-specific integrations
- State management
- Routing

## License

MIT © 2026 mtzvd

## Acknowledgments

- [Bulma](https://bulma.io) by Jeremy Thomas
- [Templ](https://templ.guide) by Joe Davidson
- Inspired by the Go and HTMX communities
