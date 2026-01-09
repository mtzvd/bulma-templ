# Canonical Project Context — Bulma-Templ Design System (templ / Go)

## Project

A standalone **Bulma-based design system**, implemented with **templ (Go)**.

SSR-first, type-safe, Bulma-first, minimal JavaScript.

Project goal:  
Create a complete, canonical, OSS-ready Bulma design system for Go/templ,  
suitable for reuse across multiple projects and for public distribution.

---

## Core Principles (NON-NEGOTIABLE)

- **Bulma-first**  
  Components reflect Bulma patterns directly. No re-invention.

- **Atomic Design (STRICT)**  
  Only these atomic levels are allowed:
  - ATOM
  - MOLECULE
  - ORGANISM

  ❌ No LAYOUT-*, COMPONENT-PART, UI-LEVEL, etc.

- **templ-first**  
  All components are written in templ.

- **SSR-first**  
  No client-side rendering assumptions.

- **Minimal JS**  
  Only Alpine.js, only where Bulma requires interactivity.  
  State is always external to components.

---

## Component API Rules

- Every component:
  - Accepts a **Props struct**
  - Accepts explicit content as `Items`
- No positional params
- No implicit behavior
- No internal state

### Content model (IMPORTANT)

This design system uses an explicit multi-child composition model.

```go
type Items []templ.Component
```

- `Items` represents an ordered list of sibling components.
- This model exists to:
  - avoid excessive `templ.Join` usage
  - better match Bulma’s structural patterns
  - enable readable, linear composition
- Single-child composition is treated as a special case of `Items`.

All content-based components accept `Items`, not `templ.Component`.

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

### Items

- Defined in `elements`
- Canonical multi-child content container
- Used across all packages (`elements`, `components`, `layout`, `form`, etc.)

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
- Exists to enable page-level and low-level composition where templ DSL becomes impractical
- Does NOT perform escaping or validation

---

## Package Structure (CURRENT ARCHITECTURE)

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

This structure reflects Bulma documentation directly.

Atomic level is declared **explicitly in code comments**,  
never inferred from directory names.

---

## Documentation Rules (STRICT)

- Every component has:
  - A component-level comment
  - An explicit atomic level line:
    `// Atomic level: ATOM | MOLECULE | ORGANISM`
- Comment language: **English only**
- Unified wording:
  - Props descriptions use **"defines"**, not "is"

---

## Design Decisions (FIXED)

- Tables: HAVE components (`Table`, `TableContainer`)
- Skeletons: HAVE components (`SkeletonBlock`, `SkeletonLines`) and support content
- Dark mode: NOT a component (application concern)
- Bulma CDN: NOT a component (layout/head concern)
- Starter Template: EXISTS as an example / quick start reference
- HTML semantics (`h1`, `p`, `section`, etc.):  
  **Handled at page or layout level**, not inside typographic components

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

- `Items` + `RenderItems` is a **deliberate deviation** from single-slot templ composition
- This choice is made to preserve:
  - Bulma fidelity
  - explicit composition
  - readability at scale
- No implicit slots, no magic, no hidden structure
