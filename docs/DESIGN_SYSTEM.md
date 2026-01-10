# Design System – Architectural Decisions (Canonical)

This document fixes the key architectural decisions
of the **Bulma-Templ Design System**.

This document is **normative**: all future components
and refactors MUST comply with it.

---

## 1. Default State Rule (CRITICAL)

**Rule:**
> The default state of any component MUST represent
> the most common, active, and usable state.

### Implications

- Boolean props MUST NOT require explicit `false`
- Modifiers (`is-disabled`, `is-active`, etc.)
  are added ONLY when the value is `true`
- HTML attributes (`disabled`, `multiple`, etc.)
  are rendered ONLY when `true`

### Example (Button)

```go
if props.Disabled {
    classes = append(classes, "is-disabled")
    attrs["disabled"] = "disabled"
}
```

❌ Forbidden:
- `templ.KV(...)`
- disabled-by-default components
- hidden implicit states

---

## 2. Items — Canonical Composition Model

```go
type Items []templ.Component
```

- Used everywhere a component accepts multiple children
- Explicit rendering via `Items.Render()`

### Render

```go
templ (i Items) Render() {
    for _, item := range i {
        @item
    }
}
```

### Why

- Explicit structure
- Predictable order
- Matches Bulma documentation
- Zero magic

---

## 3. Wrap — Infrastructure Primitive

`Wrap` is a low-level rendering primitive.

Responsibilities:
- render arbitrary HTML tag
- merge component classes with user classes
- extract `class` from Attr
- prevent duplicated attributes

Rules:
- MUST NOT be used by application code
- MAY be used by BaseElement and internals only

---

## 4. BaseElement — Canonical HTML Container

```go
templ BaseElement(props BaseElementProps, content Items)
```

### Purpose

- Single abstraction for most Bulma containers
- Eliminates duplicated tag/class/attr logic
- Improves readability and consistency

### Rule

> If a component differs only by tag name and classes,
> it MUST be implemented via BaseElement.

---

## 5. Class Handling (STRICT)

- `class` is extracted from `Attr`
- merged with component-defined classes
- NOT re-expanded via `{ Attr... }`

This guarantees:

- no duplicated `class` attributes
- correct class merging
- predictable output

---

## 6. Attr vs WrapperAttr

Final rule:

- Single-element component → `Attr`
- Multi-element (wrapper + control) → separate attrs MAY exist

But:
> Do NOT introduce `WrapperAttr` unless there is
> a real semantic distinction.

---

## 7. Flow Control & Order Freedom

Bulma allows flexible block order
(e.g. pagination: prev / next / list in any order).

Therefore:

- Components MUST NOT hardcode block order
- Structural blocks are separate templ templates
- User controls final composition

---

## 8. Canonical Goal

This design system:

- ❌ does not abstract Bulma
- ❌ does not hide HTML
- ❌ does not infer behavior

It:

- ✅ maps Bulma 1:1
- ✅ is explicit and predictable
- ✅ is composable
- ✅ is OSS-grade

---

## Status

- Architecture: FINAL
- Composition model: FINAL
- Class handling: FINAL
- Default state semantics: FINAL

Any deviation is an architectural error.
