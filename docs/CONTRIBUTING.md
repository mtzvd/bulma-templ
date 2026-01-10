# Contributing to Bulma-Templ

Thank you for your interest in contributing to **Bulma-Templ**.

This document defines **what kinds of contributions are welcome**,  
and just as importantly, **what is intentionally out of scope**.

Bulma-Templ is intentionally conservative. Please read this carefully
before opening an issue or pull request.

---

## Project philosophy (important)

Bulma-Templ follows the same philosophy as **Bulma** itself:

- It provides **canonical mappings**, not abstractions
- It shows how things are intended to be used
- It does **not** restrict how consumers compose their HTML
- It avoids opinionated extensions and framework behavior

Bulma-Templ is both:

- a **design system** (clear structure, fixed contracts)
- a **library** (meant to be used, not extended arbitrarily)

The architecture is intentionally **stable and frozen**.

---

## What contributions are welcome

### ✅ Bug fixes

- Incorrect Bulma class mappings
- Incorrect HTML output
- Attribute loss or incorrect serialization
- Regressions covered by existing tests

Bug fixes should:

- preserve existing public APIs
- include or update tests where appropriate

---

### ✅ Documentation improvements

- Clarifying comments
- Improving wording in existing docs
- Fixing typos or inconsistencies
- Adding examples that demonstrate intended usage

Documentation changes are always welcome.

---

### ✅ Tests

- New unit tests for existing components
- Regression tests for reported bugs
- Improvements to test coverage without changing behavior

Tests must follow existing testing conventions.

---

### ✅ Bulma-aligned feature additions

Contributions are welcome **only** when they correspond to:

- new features introduced in Bulma itself
- changes in Bulma documentation that require updates

In such cases:

- the implementation must follow existing design patterns
- no additional abstractions should be introduced
- changes must remain Bulma-first and Templ-first

---

## What is explicitly out of scope

The following will **not** be accepted:

### ❌ New abstractions or convenience APIs

Examples:

- generic layout helpers
- semantic inference
- automatic tag selection
- alternate composition models
- helpers that “simplify” Bulma usage

Bulma-Templ intentionally avoids these.

---

### ❌ Framework-style features

Examples:

- routing
- state management
- JavaScript frameworks
- lifecycle hooks
- client-side rendering logic

Bulma-Templ is not a framework.

---

### ❌ Opinionated extensions

Examples:

- enforcing semantic HTML
- restricting allowed markup
- introducing “best practices” not defined by Bulma
- changing defaults for convenience

Applications are free to be opinionated.  
Bulma-Templ is not.

---

## About HTML and composition

Using plain HTML alongside Bulma-Templ components is:

- fully supported
- expected at the application level
- demonstrated in examples

Bulma-Templ does **not** replace HTML and does **not** enforce semantic structure.

---

## Code style and conventions

- All comments and documentation must be **English-only**
- Follow existing comment style exactly (`COMMENT_STYLE.md`)
- Do not introduce new architectural concepts
- Respect the `Items` composition model
- Do not expose infrastructure primitives to user code

If in doubt, prefer **clarity over convenience**.

---

## Before submitting a pull request

Please ensure that:

- `task test` passes
- No architectural contracts are changed
- Changes are minimal and focused
- Commit messages are clear and descriptive

Large or conceptual changes should be discussed in an issue first.

---

## Final note

Bulma-Templ is intentionally conservative.

If you are looking to:

- experiment
- build higher-level abstractions
- integrate JavaScript frameworks
- create opinionated UI systems

Those ideas are encouraged — but **outside** this repository.

Thank you for respecting the scope and philosophy of the project.
