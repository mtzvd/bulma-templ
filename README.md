# Bulma-Templ Design System

A canonical Bulma-based design system for Go + templ.

SSR-first. Bulma-first. Type-safe. Minimal JavaScript.

---

## Status

- Architecture: FINAL
- Components: Complete
- Comment style: Frozen
- OSS readiness: v1.0

This project is considered stable and publishable.

---

## Philosophy

This design system is a faithful, explicit implementation of Bulma
for Go projects using templ.

Core values:

- Correctness over convenience
- Transparency over abstraction
- Composition over magic
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

- not a UI framework
- not a component abstraction layer
- not a JavaScript-driven UI system
- not a state management solution
- not a theme engine (dark mode is app-level)

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
- Bulma documentation maps 1:1 to components
- No re-interpretation or re-design

### Atomic Design (STRICT)

Only three atomic levels exist:

- ATOM — visual primitives
- MOLECULE — structural composition
- ORGANISM — complex UI blocks

No other levels are allowed.

Each component declares its atomic level explicitly in code comments.

### templ-first / SSR-first

- All components are written in templ
- No client-side rendering assumptions
- Server-side rendering is the default

### Minimal JavaScript

- Alpine.js only
- No inline business logic
- No hidden state inside components
- All state is external

---

## Component contract

Every component:

- uses a Props struct
- optionally accepts content as templ.Component
- exposes no positional parameters
- never manages application state
- never hides Bulma behavior
- provides Attr (templ.Attributes) as an escape hatch

---

## Attr (templ.Attributes)

Attr exists on every component and is used for:

- extra CSS classes
- Alpine (x-*)
- aria-*
- data-*

templ handles attribute merging correctly (no duplication).

---

## Package structure

The design system is located in a single top-level package:

- bulma-templ/
  - elements/
  - components/
  - form/
  - grid/
  - columns/
  - layout/

The structure mirrors Bulma documentation, not atomic levels.

Atomic level is declared explicitly in code comments,
and is never inferred from directory names.

---

## Bulma coverage

Implemented sections:

- Elements (button, icon, tag, table, skeletons, etc.)
- Components (navbar, tabs, dropdown, pagination, modal, etc.)
- Forms (input, select, textarea, checkbox, radio, file)
- Layout (container, section, hero, level, media, footer)
- Grid (columns + smart grid)

Explicitly excluded:

- dark mode logic
- Bulma CDN management
- CSS abstraction layers
- application state

---

## Comment style

Comment style is strictly defined and frozen.

See COMMENT_STYLE.md

Any deviation from the defined wording is considered an error.

---

## Intended usage

This design system is intended to be:

- vendored or imported into Go projects
- used as a shared UI foundation across services
- extended via composition, not modification

Project-specific components belong in a separate custom package.

---

## Extending the system

To extend safely:

- never modify core components
- compose existing components
- follow the same atomic rules
- follow COMMENT_STYLE.md

---

## Stability and versioning

- v1.0 — architecture frozen
- v1.x — additive, backward-compatible changes only
- v2.0 — breaking API changes

---

## License

MIT

---

## Final note

This project intentionally avoids:

- smart components
- implicit behavior
- hidden defaults

If you are looking for abstraction or automation,
this design system is not the right tool.

If you value clarity, predictability, and Bulma fidelity,
you are in the right place.
