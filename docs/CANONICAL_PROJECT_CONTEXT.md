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
  - Optionally accepts `content templ.Component`
- No positional params
- No implicit behavior

### Attr (escape hatch)

- `Attr templ.Attributes` exists everywhere
- Used for:
  - extra classes
  - Alpine (`x-*`)
  - `aria-*`, `data-*`
- templ merges class attributes correctly (never duplicates)

---

## Package Structure (CURRENT ARCHITECTURE)

``` text
/
├── bulma-templ/
│ ├── elements/
│ ├── components/
│ ├── form/
│ ├── grid/
│ ├── columns/
│ └── layout/
```

This structure reflects Bulma documentation directly.
Atomic level is declared **explicitly in code comments**, not inferred from folders.

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

---

## Current State

- All major Bulma sections implemented:
  elements, components, forms, layout, grid
- High architectural consistency
- Remaining work:
  - comment/style unification
  - usage examples
  - tests

  