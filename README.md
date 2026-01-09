# Bulma-Templ Design System

A canonical Bulma-based design system for Go + templ.

SSR-first. Bulma-first. Type-safe. Explicit composition. Minimal JavaScript.

---

## Status

- Architecture: **FINAL**
- Components: **Complete**
- Comment style: **Frozen**
- OSS readiness: **v1.0**

This project is considered **stable and publishable**.

---

## Philosophy

This design system is a **faithful, explicit implementation of Bulma**
for Go projects using **templ**.

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

- A standalone design system
- A direct mapping of Bulma concepts to templ components
- A reference-quality implementation suitable for reuse and OSS

---

## What this is NOT

- a UI framework
- a component abstraction layer
- a JavaScript-driven UI system
- a state management solution
- a theme engine (dark mode is app-level)

---

## Technology stack

- Go
- templ (SSR-first)
- Bulma CSS
- Alpine.js (optional, minimal, external)

---

## Core principles (NON-NEGOTIABLE)

### Bulma-first

- Bulma classes are used directly
- Bulma documentation maps **1:1** to components
- No reinterpretation or redesign

### Atomic Design (STRICT)

Only three atomic levels exist:

- **ATOM** — visual or infrastructure primitives
- **MOLECULE** — structural composition
- **ORGANISM** — complex UI or page-level blocks

No other levels are allowed.

Each component declares its atomic level explicitly in code comments.

### templ-first / SSR-first

- All components are written in templ
- No client-side rendering assumptions
- Server-side rendering is the default

### Minimal JavaScript

- Alpine.js only
- Only where Bulma explicitly requires interactivity
- No inline business logic
- No hidden or internal state

---

## Component contract

Every component:

- uses a **Props struct**
- accepts explicit content as `Items`
- exposes no positional parameters
- never manages application state
- never hides or reinterprets Bulma behavior
- provides `Attr` (`templ.Attributes`) as an escape hatch

There are:
- no implicit children
- no slots
- no magic composition

---

## Content model

This design system uses an explicit multi-child composition model.

```go
type Items []templ.Component
```

- `Items` represents an ordered list of sibling components
- It exists to:
  - avoid excessive `templ.Join` usage
  - better match Bulma’s structural patterns
  - enable readable, linear composition
- Single-child composition is treated as a special case of `Items`

All content-based components accept `Items`, not `templ.Component`.

---

## Infrastructure primitives (non-Bulma)

Some primitives exist to enable practical composition but are **not Bulma components**.

### Items

- Defined in `elements`
- Canonical multi-child content container
- Used consistently across all packages

### RenderItems

```go
templ RenderItems(content Items)
```

- Centralized rendering helper for `Items`
- Prevents duplication of render loops
- Ensures consistent rendering semantics

### Html

```go
templ Html(content string)
```

- Raw HTML rendering helper
- Uses `templ.Raw`
- MUST be used only with trusted, pre-sanitized content
- Intended for page-level and low-level composition
- Performs no escaping or validation

---

## Attr (templ.Attributes)

`Attr` exists on every component and is used for:

- extra CSS classes
- Alpine (`x-*`)
- `aria-*`
- `data-*`

templ handles attribute merging correctly (no duplication).

---

## Package structure

The design system is located in a single top-level module:

```text
/
├── elements/
├── components/
├── form/
├── grid/
├── columns/
├── layout/
├── docs/
├── examples (optional)
```

The structure mirrors Bulma documentation, not atomic levels.

Atomic level is declared explicitly in code comments  
and is never inferred from directory names.

---

## Bulma coverage

Implemented sections:

- Elements (button, icon, tag, table, skeletons, etc.)
- Components (navbar, tabs, dropdown, pagination, modal, card, etc.)
- Forms (input, select, textarea, checkbox, radio, file)
- Layout (container, section, hero, level, media, footer)
- Grid (columns and grid helpers)

Explicitly excluded:

- dark mode logic
- Bulma CDN management
- CSS abstraction layers
- application state management

---

## Comment style

Comment style is **strictly defined and frozen**.

See `COMMENT_STYLE.md`.

Any deviation from the defined wording is considered an error.

---

## Intended usage

This design system is intended to be:

- imported or vendored into Go projects
- used as a shared UI foundation across services
- extended via composition, not modification

Project-specific components belong in a separate custom package.

---

## Extending the system

To extend safely:

- never modify core components
- compose existing components
- follow the same atomic rules
- follow `COMMENT_STYLE.md`
- respect the `Items` content model

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

- smart components
- implicit behavior
- hidden defaults
- framework-style abstractions

If you are looking for automation or opinionated magic,  
this design system is not the right tool.

If you value clarity, predictability, and Bulma fidelity,  
you are in the right place.
