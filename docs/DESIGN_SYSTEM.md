# Bulma-Templ Design System for templ (Go) — v1.0

## Status

Architecture: **FINAL**  
Components: **Complete**  
Documentation: **Synchronized with Canonical Context**  
OSS readiness: **80%**

---

## Philosophy

This design system is a **faithful, type-safe implementation of Bulma**
for Go projects using **templ**.

The system prioritizes:

- correctness over convenience
- transparency over abstraction
- explicit composition over magic

---

## Atomic Design

Only three atomic levels exist and are strictly enforced:

- **ATOM** — visual and infrastructure primitives
- **MOLECULE** — structural composition
- **ORGANISM** — complex UI and page-level blocks

Atomic level **MUST** be declared explicitly in code comments.
No other levels are allowed.

---

## Component Contract

Every component:

- is written in **templ**
- exposes a **Props struct**
- accepts explicit content as `Items`
- never manages application state
- never hides or reinterprets Bulma behavior

There are:

- no implicit children
- no slots
- no positional parameters
- no internal state

---

## Content Model

This design system uses an explicit multi-child composition model.

```go
type Items []templ.Component
```

- `Items` represents an ordered list of sibling components.
- This model exists to:
  - avoid excessive `templ.Join` usage
  - better reflect Bulma’s structural patterns
  - enable readable, linear composition
- Single-child composition is treated as a special case of `Items`.

All content-based components accept `Items`, not `templ.Component`.

---

## Infrastructure Primitives

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

## Bulma Coverage

Implemented sections:

- Elements (button, icon, tag, table, skeletons, etc.)
- Components (navbar, tabs, dropdown, pagination, card, etc.)
- Forms (input, select, textarea, checkbox, radio, file)
- Layout (container, section, hero, level, media, footer)
- Grid (columns and grid helpers)

Explicitly excluded:

- dark mode logic
- CSS abstraction layers
- JavaScript frameworks beyond Alpine
- application state management

---

## JavaScript Policy

- Alpine.js only
- Only where Bulma explicitly requires interactivity
- No inline business logic
- No hidden or internal state inside components

---

## Semantics vs Presentation

- HTML semantics (`h1`, `p`, `section`, etc.) are handled at
  **page or layout level**
- Visual components (e.g. `Title`) do **not** define semantics
- Mixing semantics and visual components inside a single component
  is explicitly forbidden

---

## Package Structure

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

This structure mirrors Bulma documentation directly.
Atomic level is declared in code comments, never inferred from folders.

---

## Quality Bar

A component is considered **canonical** if:

- atomic level is valid and explicit
- Props naming is explicit and consistent
- content uses `Items`
- Bulma documentation can be mapped 1:1 to the component
- no implicit behavior or hidden structure exists

---

## Versioning

- **v1.0** — architecture freeze
- **v1.x** — additive improvements only
- **v2.0** — breaking API or architectural changes

---

## Architectural Notes (v1.0)

- The `Items` + `RenderItems` model is a deliberate deviation
  from single-slot templ composition
- This choice preserves:
  - Bulma fidelity
  - explicit structure
  - readability at scale
- No slots, no magic, no implicit composition

---

## Next Steps

- Kitchen Sink page
- usage examples
- tests
- final README polish
