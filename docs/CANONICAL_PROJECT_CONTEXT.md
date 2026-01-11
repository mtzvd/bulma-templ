# Canonical Project Context — Bulma-Templ Design System (templ / Go)

**Version:** 1.0  
**Status:** Stable and frozen  
**License:** MIT

## Project Overview

A standalone **Bulma-based design system**, implemented with **templ (Go)**.

**Core characteristics:**
- SSR-first
- Type-safe
- Bulma-first
- Minimal JavaScript
- Zero internal state

**Project goal:**  
Create a complete, canonical, OSS-ready Bulma design system for Go/templ,  
suitable for reuse across multiple projects and for public distribution.

**What makes this unique:**
- 1:1 mapping to Bulma CSS framework
- Explicit, predictable composition model
- No abstractions or convenience layers
- Transparent HTML output

---

## Core Principles (NON-NEGOTIABLE)

These principles are frozen for v1.0 and define the project's identity.

### 1. Bulma-First

**Principle:** Components reflect Bulma patterns directly. No re-invention.

**This means:**
- HTML structure matches Bulma documentation exactly
- CSS classes map 1:1 to component props
- No additional abstractions or modifications
- Bulma documentation is the single source of truth

**Example:**
```go
// ✅ Correct: Direct Bulma mapping
Button(ButtonProps{Color: "is-primary", Size: "is-large"}, Items{...})

// ❌ Wrong: Custom abstraction
PrimaryButton(size: Large, ...)
```

### 2. Atomic Design (STRICT)

**Allowed levels:**
- **ATOM** — Basic UI elements (button, icon, input)
- **MOLECULE** — Simple component groups (field with label)
- **ORGANISM** — Complex, self-contained components (navbar, card)

**Forbidden:**
- ❌ LAYOUT-* (e.g., LAYOUT-HEADER)
- ❌ COMPONENT-PART (e.g., COMPONENT-PART-TITLE)
- ❌ UI-LEVEL-*
- ❌ Any custom atomic classifications

**Rationale:** Atomic levels are documentation metadata, not architectural layers.

### 3. Templ-First

**Principle:** All components are written in templ.

**This means:**
- No HTML templates or other template engines
- No string concatenation for HTML
- Type-safe composition via Go types
- Compile-time template validation

### 4. SSR-First

**Principle:** No client-side rendering assumptions.

**This means:**
- Components render complete HTML on the server
- No hydration requirements
- No mandatory JavaScript execution
- Works with JavaScript disabled

### 5. Minimal JavaScript

**Principle:** Only Alpine.js, only where Bulma requires interactivity.

**This means:**
- Interactivity via Alpine.js `x-*` attributes
- No custom JavaScript
- No framework-specific code (React, Vue, etc.)
- State is always external to components

**Examples where Alpine is used:**
- Dropdown toggle
- Navbar burger menu
- Modal show/hide
- Tabs switching

**Components never:**
- Manage state internally
- Include inline scripts
- Require JavaScript to render

### 6. Zero Internal State

**Principle:** Components render HTML only. No state management.

**This means:**
- No component-level state variables
- No event handlers in component code
- No lifecycle hooks
- All state is caller-provided via props

---

## Component API Rules

Every component must follow this exact pattern:

### Standard Component Signature

```go
templ ComponentName(props ComponentNameProps, content Items) {
    // implementation
}
```

**Rules:**

1. **Always accept Props struct**
   - First parameter is always a typed Props struct
   - Never use primitive types directly
   - Props contain configuration, not content

2. **Always accept Items for content**
   - Second parameter is always `Items`
   - Never `templ.Component` (single child)
   - Never variadic `...templ.Component`

3. **No positional parameters**
   - Bad: `Button("Click me", "is-primary", true)`
   - Good: `Button(ButtonProps{Color: "is-primary", Disabled: true}, Items{Html("Click me")})`

4. **No implicit behavior**
   - No automatic defaults that change based on context
   - No hidden state or side effects
   - No magic inference

5. **No internal state**
   - Components are pure rendering functions
   - No variables that persist across renders
   - No closures capturing state

### Props Struct Pattern

```go
type ComponentNameProps struct {
    // Bulma modifiers
    Color    string // Bulma color modifier (is-primary, is-link, etc.).
    Size     string // Bulma size modifier (is-small, is-medium, is-large).
    
    // Boolean modifiers
    IsRounded bool // Applies is-rounded modifier.
    IsOutlined bool // Applies is-outlined modifier.
    
    // HTML attributes
    ID   string // HTML id attribute.
    Href string // HTML href attribute (for links).
    
    // Escape hatch
    Attr templ.Attributes // Additional HTML attributes.
}
```

**Field naming conventions:**
- Bulma modifiers: Use exact Bulma class name as field name (e.g., `Color` for `is-*` colors)
- Boolean modifiers: Use `Is*` prefix (e.g., `IsActive`, `IsLoading`)
- HTML attributes: Use exact attribute name (e.g., `ID`, `Href`, `Type`)
- Always include `Attr` as escape hatch

### Content Model (IMPORTANT)

This design system uses an explicit multi-child composition model.

```go
type Items []templ.Component
```

**Definition:**
- `Items` represents an ordered list of sibling components
- Preserves exact ordering for rendering
- Can contain zero, one, or multiple children

**Why Items exists:**

1. **Avoids excessive `templ.Join` usage**
   - Bulma components often have multiple children
   - `templ.Join` is verbose and error-prone at scale

2. **Matches Bulma's structural patterns**
   - Bulma components have predictable child structures
   - Items makes this explicit in the API

3. **Enables readable composition**
   - Linear, ordered composition
   - No magic or hidden behavior
   - Caller controls exact structure

**Usage examples:**

```go
// Multiple children
Card(CardProps{}, Items{
    CardHeader(CardHeaderProps{}, Items{Html("Title")}),
    CardContent(Items{Html("<p>Body content</p>")}),
    CardFooter(Items{
        CardFooterItem(Items{Html("Action 1")}),
        CardFooterItem(Items{Html("Action 2")}),
    }),
})

// Single child (special case of Items)
Button(ButtonProps{}, Items{Html("Click Me")})

// Empty children
Section(SectionProps{}, Items{})

// Mixed content types
Box(BoxProps{}, Items{
    Title(TitleProps{Level: "2"}, Items{Html("Heading")}),
    Content(ContentProps{}, Items{Html(`
        <p>Paragraph 1</p>
        <p>Paragraph 2</p>
    `)}),
    Button(ButtonProps{Color: "is-primary"}, Items{Html("Submit")}),
})
```

**Important:**
- Single-child composition is NOT special-cased in the API
- Always use `Items{...}` even for single child
- This keeps the API consistent and predictable

All content-based components accept `Items`, never `templ.Component` directly.

---

### Attr (escape hatch)

- `Attr templ.Attributes` exists everywhere
- Used for:
  - extra classes
  - Alpine (`x-*`)
  - `aria-*`, `data-*`
- templ merges class attributes correctly (never duplicates)

---

## Infrastructure Primitives (NON-BULMA)

Some primitives exist to enable practical composition but are **not Bulma components**.  
These are implementation details, not user-facing features.

### Items

**Location:** `elements/elements.templ`

**Definition:**
```go
type Items []templ.Component
```

**Purpose:**
- Canonical multi-child content container
- Used across all packages (`elements`, `components`, `layout`, `form`, etc.)
- Enables explicit, ordered composition

**Render method:**
```go
templ (i Items) Render() {
    for _, item := range i {
        @item
    }
}
```

**Why not exposed directly:**
- Implementation detail
- Users compose via component APIs, not Items directly
- Keeps the surface area minimal

### RenderItems

**Location:** `elements/elements.templ`

**Definition:**
```go
templ RenderItems(content Items) {
    @content.Render()
}
```

**Purpose:**
- Centralized rendering helper for `Items`
- Prevents duplication of render loops
- Ensures consistent rendering semantics
- Makes internal rendering explicit in component code

**Usage in components:**
```go
templ Card(props CardProps, content Items) {
    <div class="card">
        @RenderItems(content)
    </div>
}
```

### Html

**Location:** `elements/html.templ`

**Definition:**
```go
templ Html(content string) {
    @templ.Raw(content)
}
```

**Purpose:**
- Raw HTML rendering helper
- Enables embedding of arbitrary HTML
- Wrapper around `templ.Raw` for consistency

**Security warning:**
- ⚠️ MUST be used only with trusted, pre-sanitized content
- Does NOT perform escaping or validation
- User-provided content MUST be escaped before passing to Html

**When to use:**
- Page-level HTML that's impractical in templ DSL
- Pre-rendered markdown or other safe HTML
- Static, trusted content

**When NOT to use:**
- User-provided content (use templ components instead)
- Dynamic content that needs escaping
- Content that could contain XSS vectors

**Example:**
```go
// ✅ Good: Trusted, static content
Html("<p>Welcome to <strong>Bulma-Templ</strong></p>")

// ✅ Good: Pre-sanitized markdown output
Html(sanitizedMarkdownOutput)

// ❌ Bad: User input (XSS risk)
Html(userProvidedHTML)

// ✅ Alternative: Use templ components for user content
Content(ContentProps{}, Items{
    Html("<p>" + template.HTMLEscapeString(userText) + "</p>"),
})
```

### BaseElement (Internal)

**Location:** `elements/base_element.templ`

**Purpose:**
- Internal helper for rendering HTML elements with classes
- Not exposed to users
- Handles class merging and attribute application

**Pattern:**
```go
templ BaseElement(props BaseElementProps, content Items) {
    // Dynamic tag selection, class building, attribute merging
}
```

**Why internal:**
- Implementation detail of component rendering
- Users should never call BaseElement directly
- Keeps API surface minimal and predictable

---

## Package Structure (CURRENT ARCHITECTURE)

```text
bulma-templ/
├── elements/          # ATOM level components
│   ├── block.templ
│   ├── box.templ
│   ├── button.templ
│   ├── content.templ
│   ├── delete.templ
│   ├── icon.templ
│   ├── image.templ
│   ├── notification.templ
│   ├── progress.templ
│   ├── table.templ
│   ├── tag.templ
│   ├── title.templ
│   ├── skeleton.templ
│   ├── elements.templ  # Items, RenderItems
│   └── html.templ      # Html
│
├── components/        # MOLECULE/ORGANISM level components
│   ├── breadcrumb.templ
│   ├── card.templ
│   ├── dropdown.templ
│   ├── menu.templ
│   ├── message.templ
│   ├── modal.templ
│   ├── navbar.templ
│   ├── pagination.templ
│   ├── panel.templ
│   └── tabs.templ
│
├── form/              # Form-specific components
│   ├── checkbox.templ
│   ├── control.templ
│   ├── field.templ
│   ├── file.templ
│   ├── help.templ
│   ├── input.templ
│   ├── label.templ
│   ├── radio.templ
│   ├── select.templ
│   └── textarea.templ
│
├── layout/            # Layout components
│   ├── container.templ
│   ├── footer.templ
│   ├── hero.templ
│   ├── level.templ
│   ├── media.templ
│   └── section.templ
│
├── columns/           # Bulma columns system
│   └── columns.templ
│
├── grid/              # CSS Grid components
│   └── grid.templ
│
├── docs/              # Documentation
│   ├── CANONICAL_PROJECT_CONTEXT.md  # This file
│   ├── COMMENT_STYLE.md
│   ├── DESIGN_SYSTEM.md
│   └── LLM_INSTRUCTIONS.md
│
├── CONTRIBUTING.md    # Contribution guidelines
│
├── examples/          # Usage examples
│   ├── kitchensink/   # Comprehensive demo
│   └── starter/       # Minimal starter
│
├── cmd/dev/           # Development server
│   └── main.go
│
└── tools/             # Build tools
    └── collector/     # Code collector for LLM context
```

**Key principles:**

1. **Structure mirrors Bulma documentation**
   - Package names match Bulma section names exactly
   - Easy to find components by consulting Bulma docs

2. **Atomic level is explicit in comments**
   - Never inferred from directory names
   - Each component declares its level: `// Atomic level: ATOM`
   - Directory organization is for convenience, not semantics

3. **Flat package structure**
   - No nested subdirectories within component packages
   - All components in a package are at the same level
   - Reduces import complexity

4. **Separation of concerns**
   - `elements/` — Basic UI building blocks
   - `components/` — Complex, self-contained components
   - `form/` — Form-specific components
   - `layout/` — Page layout components
   - `columns/`, `grid/` — Grid systems

5. **Documentation co-located with code**
   - `docs/` at root level for easy access
   - Comprehensive architectural documentation
   - Living documentation that evolves with code

---

## Documentation Rules (STRICT)

See [COMMENT_STYLE.md](COMMENT_STYLE.md) for comprehensive style guide.

### Component-Level Comments

Every component MUST have:

```go
// ComponentName — Brief description.
// Atomic level: ATOM|MOLECULE|ORGANISM
//
// Detailed description of component purpose and behavior.
// What it does and what it does NOT do.
templ ComponentName(props ComponentNameProps, content Items) {
    // ...
}
```

**Required elements:**

1. **Title line**
   - Format: `// ComponentName — Brief description.`
   - Use em dash (—), not hyphen (-)
   - End with period
   - Keep under 10 words

2. **Atomic level line**
   - Exact format: `// Atomic level: ATOM` (or MOLECULE, ORGANISM)
   - Must be on second line
   - No variations allowed

3. **Blank comment line**
   - Third line must be blank: `//`

4. **Detailed description**
   - Explain what component does
   - Explain what it does NOT do (when relevant)
   - Mention key Bulma patterns

### Props Struct Comments

```go
// ComponentNameProps defines configuration for ComponentName.
//
// All props are optional unless otherwise noted.
// Bulma modifiers map directly to CSS classes.
type ComponentNameProps struct {
    Color string // Bulma color modifier (is-primary, is-link, etc.).
    Size  string // Bulma size modifier (is-small, is-medium, is-large).
    
    IsActive bool // Applies is-active modifier.
    
    Attr templ.Attributes // Additional HTML attributes.
}
```

**Rules:**

1. **Struct comment**
   - Start with: `// ComponentNameProps defines configuration for ComponentName.`
   - Use "defines", not "is" or "represents"

2. **Field comments**
   - End with period
   - For Bulma modifiers, list common values in parentheses
   - Align comments for readability (optional but encouraged)

3. **Language**
   - **English only** — no exceptions
   - Clear, concise, technical writing
   - Avoid subjective terms ("nice", "easy", "simple")

### Package-Level Documentation

```go
// Package elements provides ATOM level Bulma components.
//
// This package implements basic UI building blocks from Bulma CSS framework.
// Components in this package have minimal complexity and compose into larger structures.
package elements
```

### Test Comments

```go
// TestButton_DefaultState verifies button renders with correct default classes.
func TestButton_DefaultState(t *testing.T) {
    // ...
}
```

**Rules:**
- Describe what the test verifies
- Use "verifies" or "ensures"
- Be specific about expected behavior

---

## Design Decisions (FIXED FOR v1.0)

These decisions are frozen and define project scope.

### What Has Components

**Tables**
- ✅ HAVE components: `Table`, `TableContainer`
- Maps to Bulma's table structure
- Provides type-safe table composition

**Skeletons**
- ✅ HAVE components: `SkeletonBlock`, `SkeletonLines`, `Skeleton`
- Support content children for flexible loading states
- Follow Bulma's skeleton patterns

**Starter Template**
- ✅ EXISTS as example in `examples/starter/`
- Demonstrates minimal working setup
- Not a component, but a reference implementation

### What Does NOT Have Components

**Dark Mode**
- ❌ NOT a component
- **Rationale:** Dark mode is an application-level concern
- Implementation varies widely (CSS vars, class toggle, etc.)
- Applications handle theme switching

**Bulma CDN**
- ❌ NOT a component
- **Rationale:** CSS loading is a layout/head concern
- Different apps have different needs (CDN vs bundled)
- Applications control asset loading

**HTML Semantics**
- ❌ NOT enforced in typographic components
- **Rationale:** Semantic structure is application/page-level concern
- Examples:
  - `Title` component doesn't enforce `<h1>` vs `<h2>`
  - `Content` doesn't enforce `<article>` vs `<div>`
  - Caller decides semantic structure

**JavaScript Frameworks**
- ❌ NO framework integrations (React, Vue, Svelte)
- **Rationale:** Bulma-Templ is templ-only
- Framework bindings would be separate projects

**Routing**
- ❌ NO routing logic
- **Rationale:** Routing is application-level concern
- Go has many routing libraries
- Bulma-Templ stays UI-only

**Form Validation**
- ❌ NO validation logic
- **Rationale:** Validation is business logic
- Components render HTML, callers validate
- Applications choose validation strategy

### Default State Decisions

**Button disabled state**
- Default: Enabled
- Explicit `Disabled: true` required to disable
- Matches Bulma's default behavior

**Modal visibility**
- Default: Hidden (no `is-active` class)
- Explicit `IsActive: true` required to show
- Caller controls visibility via prop

**Dropdown state**
- Default: Closed
- Alpine.js handles toggle via `x-data`
- No component-level state

### Composition Decisions

**Why Items instead of single Component**
- Decision: Use `Items []templ.Component` everywhere
- Rationale: Bulma components often have multiple children
- Consistency: Same pattern across all components
- Explicit: Caller sees exact structure

**Why Props struct instead of many parameters**
- Decision: Single Props struct per component
- Rationale: Type-safe, extensible, self-documenting
- Avoids: Function signatures with 10+ parameters
- Enables: Easy addition of new props without breaking changes

**Why Attr escape hatch**
- Decision: Every Props has `Attr templ.Attributes`
- Rationale: Enables any attribute without changing Props
- Use cases: Alpine.js, ARIA, data-*, custom classes
- Flexibility: Callers can extend without forking

---

## Current State

- All major Bulma sections implemented:
  - elements
  - components
  - form
  - layout
  - grid
  - columns
- Unified content model based on `Items`
- High architectural consistency

Remaining work:

- usage examples
- Kitchen Sink page
- tests
- final README polish

---

## Architectural Notes (v1.0)

### Deliberate Deviations from Typical Templ Patterns

**Items composition model**

- `Items` + `RenderItems` is a **deliberate deviation** from single-child composition
- Typical templ: `templ Component(content templ.Component)`
- Bulma-Templ: `templ Component(props Props, content Items)`

**Why this deviation:**

1. **Bulma fidelity**
   - Bulma components have predictable multi-child structures
   - Single-child model doesn't match Bulma's patterns
   - Items makes structure explicit

2. **Explicit composition**
   - Caller sees exact child structure
   - No hidden wrapping or transformation
   - Order preservation guaranteed

3. **Readability at scale**
   - Clear, linear composition
   - Easy to see component hierarchy
   - Reduces cognitive load

4. **Avoids templ.Join everywhere**
   - `templ.Join` is verbose for multi-child composition
   - Items provides cleaner syntax
   - Better error messages

**Trade-off:**
- Slight verbosity: `Items{...}` wrapper required
- **Benefit:** Consistency, predictability, clarity

### Props-First Design

**Pattern:**
```go
templ Component(props ComponentProps, content Items)
```

**Why Props first:**
- Configuration before content (logical order)
- Matches function signature conventions
- Makes component behavior clear upfront

**Why typed Props:**
- Compile-time validation
- Self-documenting API
- IDE autocomplete support
- Easy to extend without breaking changes

### Attr as Universal Escape Hatch

**Every Props struct has:**
```go
Attr templ.Attributes // Additional HTML attributes.
```

**This enables:**
- Alpine.js integration (`x-*` attributes)
- ARIA attributes for accessibility
- Data attributes for JavaScript
- Custom classes without changing Props
- Any HTML attribute

**Philosophy:**
- Components provide structure
- Attr provides flexibility
- Balance between rigidity and freedom

### No Internal State (Fundamental)

**Components are pure functions:**
- Input: Props + Items
- Output: HTML
- No side effects, no state, no lifecycle

**All interactivity via:**
- Alpine.js (`x-*` attributes in `Attr`)
- Server-side rendering
- Forms and links (standard HTML)

**Benefits:**
- Predictable output
- Easy to test
- No hidden behavior
- SSR-friendly

### Frozen Architecture

**v1.0 is frozen:**
- No breaking changes to API patterns
- No changes to composition model
- No changes to naming conventions
- No changes to architectural principles

**What can change:**
- New components (matching Bulma releases)
- Bug fixes
- Documentation improvements
- Performance optimizations (if non-breaking)

**What will never change:**
- Items composition model
- Props-first signatures
- Atomic design levels
- Bulma-first philosophy

---

## Summary

Bulma-Templ is a **conservative, predictable, transparent** implementation of Bulma CSS for Go/templ.

**Core values:**
- Correctness over convenience
- Explicitness over magic
- Bulma fidelity over innovation
- Stability over features

**What makes it unique:**
- 1:1 Bulma mapping (no abstractions)
- Explicit composition (Items pattern)
- Zero internal state (pure rendering)
- Frozen architecture (v1.0 stability)

**Who it's for:**
- Go developers using templ
- Teams wanting predictable UI components
- Projects valuing correctness and transparency
- Applications needing stable, long-term foundation

**Who it's NOT for:**
- Teams wanting convenience layers
- Projects needing custom abstractions
- Applications requiring framework integration
- Developers preferring magic over explicitness

**Related documentation:**
- [CONTRIBUTING.md](../CONTRIBUTING.md) — How to contribute
- [COMMENT_STYLE.md](COMMENT_STYLE.md) — Code documentation style
- [DESIGN_SYSTEM.md](DESIGN_SYSTEM.md) — Detailed design decisions
- [LLM_INSTRUCTIONS.md](LLM_INSTRUCTIONS.md) — Guidelines for AI assistants

---

**This document is the single source of truth for Bulma-Templ's architecture.**

**Last updated:** January 11, 2026  
**Version:** 1.0  
**Status:** Frozen
