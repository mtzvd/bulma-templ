# Bulma-Templ Design System

A canonical Bulma-based design system for Go projects using Templ.

SSR-first. Bulma-first. Explicit composition. Zero internal state.

---

## Status

- Architecture: **FINAL**
- Components: **Complete**
- Comment style: **Frozen**
- OSS readiness: **v1.0**

This project is considered **stable and publishable**.

---

## Why

- You build server-rendered applications with Go
- You use `Templ`
- You want Bulma **as-is**, not reimagined
- You want predictable, inspectable HTML
- You value correctness and transparency over abstraction

---

## Philosophy

This design system is a **faithful, explicit implementation of Bulma**
for Go projects using **Templ**.

Core values:

- Correctness over convenience
- Transparency over abstraction
- Explicit composition over magic
- Bulma-first, not “Bulma-inspired”

Nothing is hidden.  
Nothing is inferred.  
Nothing manages application state.

---

## What this is

- A 1:1 mapping of Bulma components to Templ components
- Server-side rendered by default
- Explicit multi-child composition via `Items`
- A reference-quality, OSS-ready design system

---

## What this is NOT

- A UI framework
- A JavaScript component library
- A state management solution
- A theme engine (dark mode is app-level)
- A “Bulma-inspired” abstraction

---

## Technology stack

- Go
- Templ (SSR-first)
- Bulma CSS

## JavaScript

This design system is JavaScript-agnostic.

It does not ship with or depend on any JavaScript framework.
Client-side behavior (if needed) is entirely application-level.

---

## Core principles (NON-NEGOTIABLE)

### Bulma-first

- Bulma classes are used directly
- Bulma documentation maps **1:1** to components
- No reinterpretation or redesign

### Templ-first / SSR-first

- All components are written in Templ
- No client-side rendering assumptions
- Server-side rendering is the default

### Atomic Design (STRICT)

Only three atomic levels exist:

- **ATOM**
- **MOLECULE**
- **ORGANISM**

No other levels are allowed.  
Each component declares its atomic level explicitly in code comments.

### Explicit composition

- No slots
- No implicit children
- No positional parameters

### No state

Components never manage application state.

---

## Component contract

Every component:

- Uses a dedicated **Props struct**
- Accepts explicit content as `Items`
- Exposes no positional parameters
- Never manages application state
- Never hides or reinterprets Bulma behavior
- Provides `Attr` (`Templ.Attributes`) as an escape hatch

There are:

- No implicit children
- No slots
- No magic composition

---

## Content model

All content-based components use an explicit multi-child model:

```
type Items []Templ.Component
```

- Ordered, explicit composition
- No hidden rendering logic
- Matches Bulma’s structural patterns
- Single-child composition is a special case of `Items`

---

## Infrastructure primitives (non-Bulma)

Some primitives exist to support composition but are **not Bulma components**.

### Items

- Canonical multi-child content container
- Used consistently across all packages

### Html

```
Templ Html(content string)
```

- Renders raw HTML via `Templ.Raw`
- MUST be used only with trusted, pre-sanitized content
- Intended for page-level and low-level composition
- Performs no escaping or validation

---

## Example

```
@elements.Button(
  elements.ButtonProps{
    Color: "is-primary",
  },
  elements.Items{
    elements.Html("Click me"),
  },
)
```

---

## Package structure

```
/
├── elements/
├── components/
├── form/
├── layout/
├── grid/
├── columns/
├── docs/
├── examples/
```

The structure mirrors Bulma documentation.  
Atomic level is declared explicitly in code comments and is never inferred.

---

## Bulma coverage

Implemented sections:

- Elements
- Components
- Forms
- Layout
- Grid and Columns

Explicitly excluded:

- Dark mode logic
- Bulma CDN management
- CSS abstraction layers
- Application state management

---

## Documentation

- `docs/DESIGN_SYSTEM.md` — architectural contract (FINAL)
- `docs/COMMENT_STYLE.md` — frozen comment style
- `docs/CANONICAL_PROJECT_CONTEXT.md` — project intent and scope

These documents are normative.

---

## Testing

Minimal but sufficient unit test coverage is provided for v1.0:

- Wrap
- BaseElement
- Items
- Button
- Pagination

Focus:

- Structural correctness
- Attribute handling
- Regression safety

---

## Intended usage

This design system is intended to be:

- Imported or vendored into Go projects
- Used as a shared UI foundation across services
- Extended via composition, not modification

Project-specific components belong in a separate package.

---

## Stability and versioning

- **v1.0** — architecture frozen
- **v1.x** — additive, backward-compatible changes only
- **v2.0** — breaking API or architectural changes

---

## License

MIT

---

## Final note

This project intentionally avoids:

- Implicit behavior
- Hidden defaults
- Smart components
- Framework-style abstractions

If you value clarity, predictability, and Bulma fidelity —  
this design system is the right tool.

Bulma-Templ follows the same philosophy as Bulma itself.

It provides canonical component mappings and structural guidance,
but does not restrict how consumers compose their markup.
Using plain HTML alongside Bulma-Templ components is fully supported
and often expected at the application level.
