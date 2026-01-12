# Contributing to Bulma-Templ

Thank you for your interest in contributing to **Bulma-Templ**.

This document defines **what kinds of contributions are welcome** and, just as importantly, **what is intentionally out of scope**.

**Bulma-Templ is intentionally conservative.** Please read this document carefully before opening an issue or pull request.

**Quick Links:**

- [DESIGN_SYSTEM.md](docs/DESIGN_SYSTEM.md) — Normative architectural decisions
- [COMMENT_STYLE.md](docs/COMMENT_STYLE.md) — Comment formatting rules
- [LLM_INSTRUCTIONS.md](docs/LLM_INSTRUCTIONS.md) — Non-negotiable rules for AI assistants
- [Bulma Documentation](https://bulma.io/documentation/) — CSS framework reference

---

## Project Philosophy (CRITICAL)

Bulma-Templ follows the same philosophy as **Bulma** itself:

- Provides **canonical mappings**, not abstractions
- Shows how things are intended to be used
- Does **not** restrict how consumers compose their HTML
- Avoids opinionated extensions and framework behavior
- Maintains **1:1 correspondence** with Bulma documentation

**Core Principles** (detailed in [CANONICAL_PROJECT_CONTEXT.md](docs/CANONICAL_PROJECT_CONTEXT.md)):


- **Bulma-first**: Components reflect Bulma patterns directly, no re-invention
- **templ-first**: All components are written in templ
- **SSR-first**: No client-side rendering assumptions
- **Minimal JS**: Only Alpine.js, only where Bulma requires interactivity
- **Zero internal state**: Components render HTML, nothing more

**Atomic Design (STRICT)**:

- Only three levels: ATOM, MOLECULE, ORGANISM
- No custom levels (LAYOUT-*, COMPONENT-PART, etc.)

Bulma-Templ is both:


- A **design system** (clear structure, fixed contracts)
- A **library** (meant to be used, not extended arbitrarily)

**The architecture is intentionally stable and frozen (v1.0).**

---

## What Contributions Are Welcome

### ✅ Bug Fixes

**Definition:** Code that produces incorrect output compared to Bulma documentation.

**Examples:**

- Incorrect Bulma class mappings
- Incorrect HTML structure or attribute serialization
- Missing or incorrectly applied modifiers
- Regressions covered by existing tests

**Requirements:**


- Preserve existing public APIs
- Include or update tests where appropriate
- Reference Bulma documentation for correctness

---

### ✅ Documentation Improvements

**Definition:** Changes that improve clarity without changing meaning.

**Examples:**

- Clarifying component-level comments
- Improving wording in existing documentation
- Fixing typos, grammar, or inconsistencies
- Adding examples that demonstrate **intended usage**
- Correcting cross-references between docs

**Requirements:**


- Follow [COMMENT_STYLE.md](docs/COMMENT_STYLE.md) exactly
- Maintain English-only comments and documentation
- Do not introduce new architectural concepts

---

### ✅ Tests

**Definition:** Additions or improvements to test coverage.

**Examples:**

- New unit tests for existing components
- Regression tests for reported bugs
- Improvements to test coverage without changing behavior
- Better assertions or test organization

**Requirements:**


- Follow existing testing conventions
- Use `render_test.go` patterns for rendering tests
- Do not change component behavior to accommodate tests

---

### ✅ Bulma-Aligned Feature Additions

**Definition:** Features corresponding to **new Bulma releases** only.

**When acceptable:**

- New features introduced in Bulma CSS framework
- Changes in Bulma documentation requiring updates
- New Bulma components or modifiers

**Requirements:**


- Implementation MUST follow existing design patterns
- No additional abstractions beyond 1:1 mapping
- Changes must remain **Bulma-first** and **Templ-first**
- Reference specific Bulma version and documentation
- Match Bulma HTML structure exactly

**Verification steps:**


1. Find the component in [Bulma documentation](https://bulma.io/documentation/)
2. Compare HTML structure character-by-character
3. Verify all CSS classes match exactly
4. Ensure all modifiers are supported
5. Test with Bulma CSS to confirm visual output

**Example:** If Bulma 1.1.0 introduces a new `.is-ghost` modifier, adding support for that modifier is welcome.

---

## What is explicitly out of scope

The following will **not** be accepted:

### ❌ New abstractions or convenience APIs

**Examples of what will be rejected:**

- Generic layout helpers (e.g., `Grid()`, `Flex()`)
- Semantic inference (e.g., automatic role/ARIA detection)
- Automatic tag selection (e.g., choosing between `<button>` and `<a>`)
- Alternate composition models (e.g., slots, named children, builders)
- Helpers that "simplify" Bulma usage (e.g., `QuickForm()`, `StandardCard()`)
- Wrapper components that combine multiple Bulma components
- Builder patterns (e.g., `Button().Primary().Large()`)
- DSL layers (e.g., custom enum types for Bulma classes)

**Why rejected:**


- Bulma-Templ provides building blocks, not convenience layers
- Applications have different needs for abstraction
- Every abstraction is an opinion we don't want to enforce
- These belong in application-specific libraries

Bulma-Templ intentionally avoids these.

---

### ❌ Framework-Style Features

**Why rejected:** Bulma-Templ is a component library, not a framework.

**Examples:**

- Routing or navigation logic
- State management (stores, signals, context)
- JavaScript framework integration
- Lifecycle hooks or event systems
- Client-side rendering logic
- Form validation or data binding

**Principle:** Components render HTML. Period.

---

### ❌ Opinionated Extensions

**Why rejected:** Applications are opinionated. Libraries are not.

**Examples:**

- Enforcing semantic HTML (`<nav>` must contain `<a>`)
- Restricting allowed markup patterns
- Introducing "best practices" not defined by Bulma
- Changing defaults for convenience
- Validation of prop combinations

**Principle:** Caller control is absolute. Components never enforce opinion.

---

### ❌ Infrastructure Exposure

**Why rejected:** Internal primitives are implementation details.

**Examples:**

- Exposing `BaseElement` directly to users
- Making `Wrap` a public API
- Creating new composition primitives
- Extending the `Items` model

**Principle:** Users interact with components, not infrastructure.

---

## About HTML and Composition

Using **plain HTML** alongside Bulma-Templ components is:

- Fully supported
- Expected at the application level
- Demonstrated in Kitchen Sink examples

**Bulma-Templ does not replace HTML** and does not enforce semantic structure.

**Example (fully valid):**

```go
bulma.Card(bulma.CardProps{}, templ.Raw(`
    <header class="card-header">
        <p class="card-header-title">Custom HTML</p>
    </header>
`))
```

Applications may choose to be semantic. Bulma-Templ does not impose that choice.

---

## Code Style and Conventions

### Language and Formatting

- All comments and documentation must be **English-only**
- Follow [COMMENT_STYLE.md](docs/COMMENT_STYLE.md) exactly:
  - Use `//` only, never `/* */`
  - No blank comment lines
  - All field comments end with a period
  - Component title format: `// ComponentName — Brief description.`
  - Include atomic level: `// Atomic level: ATOM|MOLECULE|ORGANISM`

### Architecture Rules


- **Do not introduce new abstractions** or convenience layers
- **Respect the `Items` composition model** — the only composition pattern
- **Do not expose infrastructure primitives** to user code
- **Follow existing design patterns** — consistency over innovation
- **No logic or state management** — components render HTML only

### Component API Requirements

Every component must:

- Accept a `Props` struct for configuration
- Accept `Items` for content (not `templ.Component`)
- Include `Attr templ.Attributes` in Props for escape hatch
- Use explicit content model, no positional parameters
- Have zero internal state

### Testing Requirements


- Use `render_test.go` helper pattern for rendering components
- Test default states, modifiers, and edge cases
- Verify HTML structure matches Bulma documentation
- Follow existing test naming conventions (`TestComponentName_Behavior`)
- No tests that require changing component behavior

If in doubt, prefer **clarity over convenience**.

---

## Development Setup

### Prerequisites

- Go 1.25.1 or later
- [Templ CLI](https://templ.guide/) installed
- [Task](https://taskfile.dev/) (optional but recommended)

### Setup Instructions

1. Clone the repository:

   ```bash
   git clone https://github.com/mtzvd/bulma-templ.git
   cd bulma-templ
   ```

2. Install dependencies:

   ```bash
   go mod download
   ```

3. Generate templ files:

   ```bash
   task templ
   # or manually:
   templ generate
   ```

### Available Commands

Using Task (recommended):

```bash
task clean      # Remove all generated *_templ.go files
task templ      # Generate templ files (runs clean first)
task test       # Generate and run all tests
task run        # Run dev server
task dump       # Collect all code into all_code.txt
```

Without Task:

```bash
templ generate              # Generate templ files
go test ./...              # Run tests
go run ./cmd/dev           # Run dev server
```

### Development Workflow

1. **Make changes to `.templ` files** (not `*_templ.go` generated files)
2. **Regenerate:** `task templ` or `templ generate`
3. **Test:** `task test` or `go test ./...`
4. **Verify:** Check Kitchen Sink demo with `task run`
5. **Commit:** Commit both `.templ` files AND regenerated `*_templ.go` files

**Important:** Generated `*_templ.go` files ARE tracked in git for pkg.go.dev compatibility. Always regenerate before committing changes to `.templ` files.

---

### Project Structure

```text
bulma-templ/
├── elements/          # ATOM level components (button, box, icon, etc.)
├── components/        # MOLECULE/ORGANISM components (card, navbar, etc.)
├── layout/            # Layout components (container, section, hero, etc.)
├── form/              # Form components (input, checkbox, select, etc.)
├── columns/           # Grid system (columns, column)
├── grid/              # CSS Grid components
├── cmd/dev/           # Development server (Kitchen Sink demo)
├── examples/          # Usage examples
│   ├── kitchensink/   # Comprehensive demo of all components
│   └── starter/       # Minimal starter example
├── docs/              # Documentation (you are here)
│   ├── CANONICAL_PROJECT_CONTEXT.md
│   ├── COMMENT_STYLE.md
│   ├── CONTRIBUTING.md
│   ├── DESIGN_SYSTEM.md
│   └── LLM_INSTRUCTIONS.md
└── tools/             # Build tools (code collector, etc.)
```

**Key files:**

- `*.templ` — Source template files (tracked in git)
- `*_templ.go` — Generated files (tracked in git for pkg.go.dev, regenerated by `task templ`)
- `*_test.go` — Unit tests
- `render_test.go` — Test helper for component rendering (in each package)

---

## Testing Your Changes

### Running Tests

```bash
# Run all tests
task test

# Run tests for specific package
go test ./elements
go test ./components

# Run with verbose output
go test -v ./...
```

### Writing Tests

Example test structure:

```go
package elements

import (
    "strings"
    "testing"
)

func TestButton_DefaultState(t *testing.T) {
    html := render(t,
        Button(
            ButtonProps{},
            Items{Html("Click")},
        ),
    )

    if !strings.Contains(html, `class="button"`) {
        t.Fatalf("expected button class, got: %s", html)
    }
}
```

Key testing patterns:

- Use `render(t, component)` helper to render components
- Verify HTML structure with `strings.Contains()` or `strings.HasPrefix()`
- Test default states (no modifiers applied)
- Test modifier application (color, size, state)
- Test edge cases and omitted fields

---

## Before Submitting a Pull Request

Please ensure that:

1. **Tests pass**: `task test` completes without errors
2. **Generated files regenerate correctly**: Run `task templ` to verify (do NOT commit `*_templ.go` files)
3. **No architectural contracts are changed**: Review [DESIGN_SYSTEM.md](docs/DESIGN_SYSTEM.md)
4. **Comments follow style guide**: See [COMMENT_STYLE.md](docs/COMMENT_STYLE.md)
5. **Changes are minimal and focused**: One logical change per PR
6. **Commit messages are clear**: Describe what and why, not how

### PR Description Template

```markdown
## Description
Brief description of the change

## Type of Change
- [ ] Bug fix (incorrect Bulma mapping)
- [ ] Documentation improvement
- [ ] Test addition/improvement
- [ ] Bulma-aligned feature (specify Bulma version)

## Bulma Reference
Link to relevant Bulma documentation

## Testing
How was this tested?

## Checklist
- [ ] Tests pass (`task test`)
- [ ] Comments follow COMMENT_STYLE.md
- [ ] No new abstractions introduced
- [ ] Generated files committed
```

### Discussion Before Implementation

Large or conceptual changes should be discussed in an issue first:

- New components (even if Bulma-aligned)
- Changes to existing component APIs
- Documentation restructuring
- Testing infrastructure changes

---

## Common Questions

### Q: Can I add a helper function to simplify common patterns?

**A:** No. Bulma-Templ provides 1:1 mappings only. Helper abstractions belong in application code.

**Example — What to do instead:**

```go
// ❌ Don't add to Bulma-Templ
func PrimaryButton(text string) templ.Component { ... }

// ✅ Do in your application
// my-app/ui/helpers.go
func PrimaryButton(text string) templ.Component {
    return elements.Button(
        elements.ButtonProps{Color: "is-primary"},
        elements.Items{elements.Html(text)},
    )
}
```

---

### Q: Can I add validation or business logic to a component?

**A:** No. Components render HTML only. All logic belongs in the caller.

**Example — What to do instead:**

```go
// ❌ Don't add to Bulma-Templ
func EmailInput(props InputProps) templ.Component {
    if !isValidEmail(props.Value) {
        props.Color = "is-danger"
    }
    // ...
}

// ✅ Do in your application
func EmailField(email string) templ.Component {
    isValid := validateEmail(email)
    color := ""
    if !isValid {
        color = "is-danger"
    }
    return form.Input(
        form.InputProps{Type: "email", Value: email, Color: color},
    )
}
```

---

### Q: Can I change how `Items` works or add slots/named children?

**A:** No. `Items` is the only composition model and is frozen.

**Rationale:** `Items` provides explicit, ordered composition that matches Bulma's structural patterns exactly. Any alternative model would introduce abstraction.

---

### Q: Can I add TypeScript/JSX/other template syntax?

**A:** No. Bulma-Templ is templ-only.

**Rationale:** The project is specifically a Go/templ implementation. Other template systems would be separate projects.

---

### Q: Can I add a component for a custom Bulma extension or theme?

**A:** No. Only official Bulma components and modifiers are supported.

**What to do instead:**

- Fork Bulma-Templ for your custom theme
- Create a separate library that extends Bulma-Templ
- Use `Attr` to add custom classes in your application

---

### Q: Can I contribute translations or i18n support?

**A:** No. Content is caller-provided. Components don't handle text or localization.

**Rationale:** Components accept `Items` (content) as input. The caller controls all text, including localization.

---

### Q: Can I add accessibility (ARIA) attributes automatically?

**A:** No. Callers control all attributes via `Attr`. Components don't infer accessibility needs.

**Rationale:** Accessibility requirements vary by application context. Automatic ARIA would be opinionated and potentially incorrect.

**What to do instead:**

```go
// Application code
elements.Button(
    elements.ButtonProps{
        Attr: templ.Attributes{
            "aria-label":    "Close dialog",
            "aria-pressed":  "false",
            "aria-disabled": "false",
        },
    },
    elements.Items{elements.Html("×")},
)
```

---

### Q: I found a missing Bulma modifier. Can I add it?

**A:** Yes! This is a bug fix. Please:

1. Reference the [Bulma documentation](https://bulma.io/documentation/)
2. Show the missing modifier in the docs
3. Submit a PR with the fix and a test

---

### Q: Can I improve performance by adding memoization or caching?

**A:** No. Templ and Go handle performance. Optimizations are application-specific.

**Rationale:** Premature optimization would add complexity without proven benefit. Applications can optimize at their level.

---

### Q: The current API is verbose. Can I make it more concise?

**A:** No. Verbosity trades off for explicitness and transparency.

**Philosophy:** We optimize for clarity and correctness, not for minimal typing. Applications can build concise wrappers if needed.

---

## Examples of Good Contributions

### Example 1: Bug Fix (Missing Bulma Modifier)

**Issue:** Button component doesn't support `is-ghost` modifier from Bulma 1.0.2

**Good PR:**

```go
// In elements/button.templ

type ButtonProps struct {
    // ... existing fields
    IsGhost bool // Bulma ghost modifier (transparent background).
}

templ Button(props ButtonProps, content Items) {
    {{
        classes := []string{"button"}
        
        // ... existing modifiers
        
        if props.IsGhost {
            classes = append(classes, "is-ghost")
        }
    }}
    // ... rest of component
}
```

**Test:**

```go
func TestButton_Ghost(t *testing.T) {
    html := render(t,
        Button(
            ButtonProps{IsGhost: true},
            Items{Html("Ghost")},
        ),
    )
    
    if !strings.Contains(html, "is-ghost") {
        t.Fatalf("expected is-ghost class")
    }
}
```

**PR Description:**

> Adds support for `is-ghost` modifier introduced in Bulma 1.0.2
>
> Reference: <https://bulma.io/documentation/elements/button/#ghost-button>
>
> - Adds `IsGhost bool` field to `ButtonProps`
> - Appends `is-ghost` class when true
> - Includes test for ghost button rendering

---

### Example 2: Documentation Improvement

**Issue:** Comment for `Pagination` component unclear about ordering

**Good PR:**

```go
// Pagination — Bulma pagination component.
// Atomic level: ORGANISM
//
// Pagination renders a navigation component for multi-page content.
// The order of child components (PaginationPrev, PaginationNext, PaginationList)
// is preserved exactly as provided in Content.Items, allowing flexible layouts.
//
// This component does NOT manage page state or navigation logic.
// Callers must handle current page tracking and URL generation.
```

**PR Description:**

> Clarifies pagination component behavior
>
> - Explains that child ordering is preserved (not rearranged)
> - Clarifies that pagination doesn't manage state
> - Maintains existing comment style conventions

---

### Example 3: Test Coverage Improvement

**Issue:** No test for `Dropdown` with multiple items

**Good PR:**

```go
func TestDropdown_MultipleItems(t *testing.T) {
    html := render(t,
        Dropdown(
            DropdownProps{},
            Items{
                DropdownTrigger(...),
                DropdownMenu(
                    DropdownMenuProps{},
                    Items{
                        DropdownItem(...),
                        DropdownItem(...),
                        DropdownItem(...),
                    },
                ),
            },
        ),
    )
    
    // Verify structure
    if !strings.Contains(html, "dropdown-menu") {
        t.Fatalf("missing dropdown-menu")
    }
    
    // Verify multiple items rendered
    itemCount := strings.Count(html, "dropdown-item")
    if itemCount != 3 {
        t.Fatalf("expected 3 items, got %d", itemCount)
    }
}
```

**PR Description:**

> Adds test for dropdown with multiple items
>
> - Tests that multiple DropdownItem components render correctly
> - Verifies dropdown-menu container is present
> - Follows existing test patterns

---

## Examples of Rejected Contributions

### Example 1: Convenience Helper (❌ Rejected)

```go
// ❌ This would be rejected
func PrimaryButton(text string, onClick string) templ.Component {
    return Button(
        ButtonProps{
            Color: "is-primary",
            Attr: templ.Attributes{"@click": onClick},
        },
        Items{Html(text)},
    )
}
```

**Why rejected:** This is a convenience abstraction, not a Bulma feature. Belongs in application code.

---

### Example 2: Opinionated Validation (❌ Rejected)

```go
// ❌ This would be rejected
templ Input(props InputProps) {
    {{
        // Automatically add is-danger if email is invalid
        if props.Type == "email" && !isValidEmail(props.Value) {
            classes = append(classes, "is-danger")
        }
    }}
}
```

**Why rejected:** Components don't validate data. Validation is application logic.

---

### Example 3: Alternative Composition (❌ Rejected)

```go
// ❌ This would be rejected
type CardProps struct {
    Header templ.Component  // Named slot
    Body   templ.Component  // Named slot
    Footer templ.Component  // Named slot
}
```

**Why rejected:** This introduces a different composition model than `Items`. The architecture is frozen.

---

## Getting Help

- **Documentation**: Read [docs/](.) for comprehensive architecture details
- **Examples**: See [examples/kitchensink](../examples/kitchensink) for complete usage patterns
- **Issues**: [Open an issue](https://github.com/mtzvd/bulma-templ/issues) for bugs or questions
- **Discussions**: Use [GitHub Discussions](https://github.com/mtzvd/bulma-templ/discussions) for general questions

**Before asking:**

1. Read [DESIGN_SYSTEM.md](docs/DESIGN_SYSTEM.md) for architectural decisions
2. Check [LLM_INSTRUCTIONS.md](docs/LLM_INSTRUCTIONS.md) for common patterns and examples
3. Review the Kitchen Sink demo in [examples/kitchensink](../examples/kitchensink)
4. Search [existing issues](https://github.com/mtzvd/bulma-templ/issues) for similar questions
5. Consult [Bulma documentation](https://bulma.io/documentation/) for CSS reference

---

## Final Note

**Bulma-Templ is intentionally conservative by design.**

### What This Project Is

- **A faithful mapping** of Bulma CSS to Go/templ components
- **A stable foundation** for building Go web applications
- **A reference implementation** of Bulma patterns in templ
- **A design system** with clear, frozen architecture

### What This Project Is Not

- A convenience layer or helper library
- A full-stack framework
- An opinionated UI system
- A place for experimentation

### If You Want More

If you're looking to:


- Experiment with new patterns or abstractions
- Build higher-level convenience APIs
- Integrate JavaScript frameworks (React, Vue, Svelte)
- Create opinionated UI systems or themes
- Add validation, state management, or business logic
- Build pre-composed layouts or page templates

**Those ideas are valid and encouraged — but belong outside this repository.**

### Consider Instead

1. **Fork the project** for your specific needs and customizations
2. **Build a wrapper library** on top of Bulma-Templ with your conventions
3. **Create application-specific component libraries** in your own codebase
4. **Use Bulma-Templ as primitives** and compose your own abstractions

### Why This Matters


By staying **minimal and unopinionated**, Bulma-Templ:

- Remains universally applicable
- Avoids forcing architectural decisions
- Stays maintainable and predictable
- Serves as a reliable foundation for diverse use cases

The goal is not to provide every possible feature, but to provide **correct, predictable, and transparent** Bulma components that you can trust.

This intentional conservatism is a feature, not a limitation.

Thank you for respecting the scope and philosophy of the project.

---

## Related Documentation

- **[DESIGN_SYSTEM.md](docs/DESIGN_SYSTEM.md)** — Normative architectural decisions
- **[COMMENT_STYLE.md](docs/COMMENT_STYLE.md)** — Comment formatting rules
- **[LLM_INSTRUCTIONS.md](docs/LLM_INSTRUCTIONS.md)** — Non-negotiable rules for AI assistants
- **[CANONICAL_PROJECT_CONTEXT.md](docs/CANONICAL_PROJECT_CONTEXT.md)** — Comprehensive project context
- **[Bulma Documentation](https://bulma.io/documentation/)** — CSS framework reference

---

## License

By contributing, you agree that your contributions will be licensed under the same license as the project.
