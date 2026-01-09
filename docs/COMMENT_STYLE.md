# Comment Style Guide — Bulma-Templ Design System (templ / Go)

This document defines the canonical comment style for all components in the Bulma-Templ Design System.

This style is frozen for v1.0.

---

## Component-level comments

Every component MUST include the following metadata at the top of its declaration:

1. Short Bulma-aligned title
2. Explicit atomic level
3. Responsibility definition
4. Explicit non-responsibility

### Canonical form

```go
// Button — Bulma button component.
// Atomic level: ATOM
//
// Button defines a clickable UI element styled with Bulma.
// It does NOT manage state, routing, or business logic.
```

### Allowed verbs (STRICT)

```text
defines

renders

wraps

applies

groups

composes

```

### Forbidden verbs

```text
is a

represents

responsible for

manages

handles

abstracts

simplifies
```

## Props comments

### Rules

- Every field has a comment

- Starts with a verb

    Uses/ defines / applies / controls

Example

```go

// Size defines the Bulma size modifier.
Size string
```

- Atomic level

Must be present and exact:

```go

// Atomic level: ATOM | MOLECULE | ORGANISM

```

No variants allowed.

## Freeze policy

From v1.0 onward:
New components MUST follow this guide
Existing comments are only changed for:

- correctness

- clarity

- documentation expansion
