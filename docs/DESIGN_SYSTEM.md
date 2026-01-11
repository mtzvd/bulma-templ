# Design System — Architectural Decisions (Canonical)

This document defines the key architectural decisions of the **Bulma-Templ Design System**.

This document is **normative**: all future components and refactors MUST comply with it.

**Reference:** [Bulma CSS Framework Documentation](https://bulma.io/documentation/)

---

## 1. Default State Rule (CRITICAL)

**Rule:**

> The default state of any component MUST represent the most common, active, and usable state.

### Implications

- Boolean props MUST NOT require explicit `false`
- Modifiers (`is-disabled`, `is-active`, etc.) are added ONLY when the value is `true`
- HTML attributes (`disabled`, `multiple`, etc.) are rendered ONLY when `true`

### Example (Button)

```go
if props.Disabled {
    classes = append(classes, "is-disabled")
    attrs["disabled"] = "disabled"
}
```

### ❌ Forbidden

- `templ.KV(...)` for conditional attributes
- disabled-by-default components
- hidden implicit states

### ✅ Correct Pattern

```go
// Default is enabled and visible
type ButtonProps struct {
    Disabled bool  // false by default
}

// NOT:
type ButtonProps struct {
    Enabled bool  // requires explicit true
}
```

---

## 2. Items — Canonical Composition Model

```go
type Items []templ.Component
```

**What Items accepts:**

- `templ.Component` instances
- Raw HTML strings (via `templ.Raw()`)
- Any value implementing the `templ.Component` interface

Items is the ONLY composition model in Bulma-Templ:

- NO slots
- NO named children
- NO positional parameters

### Render Method

```go
templ (i Items) Render() {
    for _, item := range i {
        @item
    }
}
```

### Why Items

- Explicit structure
- Predictable order
- Matches Bulma documentation 1:1
- Zero magic
- Caller-controlled composition

### Usage Examples

```go
// Composing components
@Card(CardProps{}, elements.Items{
    CardHeader(...),
    CardContent(...),
    CardFooter(...),
})

// Mixing components and HTML
@Button(ButtonProps{}, elements.Items{
    Icon(...),
    elements.Html("Save"),
})
```

---

## 3. Atomic Design Levels (STRICT)

Every component MUST be classified as one of three atomic levels:

### ATOM

**Definition:** Basic building blocks that cannot be decomposed further.

**Characteristics:**

- Single, standalone element
- No internal composition
- Renders one Bulma class (e.g., `.button`, `.tag`, `.icon`)

**Examples:** Button, Input, Icon, Tag, Label, Help

### MOLECULE

**Definition:** Simple combinations of atoms into functional groups.

**Characteristics:**

- Groups 2-3 atoms
- Single responsibility
- Part of a larger organism

**Examples:** Field (label + control + help), CardHeader, LevelItem, BreadcrumbItem

### ORGANISM

**Definition:** Complex UI components composed of molecules and atoms.

**Characteristics:**

- Multiple nested components
- Complete UI section
- Self-contained functionality

**Examples:** Navbar, Card, Modal, Pagination, Tabs, Hero

### Classification Rule

> When in doubt, classify as the LOWEST viable level. Prefer ATOM over MOLECULE, MOLECULE over ORGANISM.

---

## 4. Props Design Principles (MANDATORY)

### Value Types Only

All Props structs MUST use value types, NEVER pointers:

```go
✅ type ButtonProps struct {
    Disabled bool
    Color    string
}

❌ type ButtonProps struct {
    Disabled *bool
    Color    *string
}
```

**Rationale:** Zero values are meaningful defaults. Pointers create nil-checking complexity.

### Plain Strings for CSS Modifiers

Bulma modifiers (color, size, alignment) MUST be plain strings:

```go
✅ type ButtonProps struct {
    Color string  // "is-primary", "is-danger", etc.
    Size  string  // "is-small", "is-large", etc.
}

❌ type ButtonProps struct {
    Color ButtonColor  // enum
    Size  ButtonSize   // enum
}
```

**Rationale:** Bulma classes are Bulma's API. We map 1:1, not abstract.

### Attr for HTML Pass-Through

Every component MUST expose `Attr templ.Attributes`:

```go
type ComponentProps struct {
    // ... component-specific fields
    
    // Attr contains additional HTML attributes for the `.component` element.
    Attr templ.Attributes
}
```

### WrapperAttr Pattern

Components with wrapper + native element MAY expose separate attributes:

```go
type SelectProps struct {
    // Attr for native <select>
    Attr templ.Attributes
    
    // WrapperAttr for .select wrapper
    WrapperAttr templ.Attributes
}
```

**Rule:** Only use WrapperAttr when there is a semantic distinction between wrapper and control.

---

## 5. Sentinel Values Pattern

For optional numeric attributes, use package-level constants:

```go
// ProgressUnset indicates that a progress value is not set.
// Use this constant when you want to omit the attribute.
const ProgressUnset = -1

type ProgressProps struct {
    Max     int  // Use ProgressUnset to omit
    Current int  // Use ProgressUnset to omit
}

// In component
if props.Max != ProgressUnset {
    attrs["max"] = props.Max
}
```

### ❌ Do NOT use pointers

```go
❌ type ProgressProps struct {
    Max     *int
    Current *int
}
```

---

## 6. Wrap — Infrastructure Primitive

`Wrap` is a low-level rendering primitive used internally.

### Responsibilities

- Render arbitrary HTML tag
- Merge component classes with user classes
- Extract `class` from Attr
- Prevent duplicated attributes

### Rules

- MUST NOT be used by application code
- MAY be used by BaseElement and internal infrastructure only
- NOT part of public API

---

## 7. BaseElement — Canonical HTML Container

```go
templ BaseElement(props BaseElementProps, content Items)
```

### Purpose

BaseElement is the single abstraction for most Bulma containers:

- Eliminates duplicated tag/class/attr logic
- Improves readability and consistency
- Enforces correct class merging

### Rule

> If a component differs only by tag name and base classes, it MUST be implemented via BaseElement.

### Real Example

```go
// Box component using BaseElement
templ Box(props BoxProps, content Items) {
    @BaseElement(
        BaseElementProps{
            Tag:         "div",
            BaseClasses: []string{"box"},
            Attr:        props.Attr,
        },
        content,
    )
}

// Notification component using BaseElement
templ Notification(props NotificationProps, content Items) {
    @BaseElement(
        BaseElementProps{
            Tag:         "div",
            BaseClasses: []string{"notification", props.Color},
            Attr:        props.Attr,
        },
        content,
    )
}
```

### When NOT to use BaseElement

- Component has complex internal structure
- Multiple nested elements required
- Conditional rendering logic beyond classes

---

## 8. Class Handling (STRICT)

### Rules

1. `class` attribute is extracted from `Attr`
2. Merged with component-defined classes
3. NOT re-expanded via `{ Attr... }`

This guarantees:

- No duplicated `class` attributes
- Correct class merging
- Predictable HTML output

### ✅ Correct

```go
classes := []string{"button"}
if props.Color != "" {
    classes = append(classes, props.Color)
}

attrs := templ.Attributes{}
for k, v := range props.Attr {
    attrs[k] = v
}
```

### ❌ Wrong

```go
// DO NOT spread Attr without extracting class
<div class="button" { props.Attr... }>
```

---

## 9. CSS Helpers Policy (MANDATORY)

Bulma CSS helper classes (spacing, typography, colors) MUST remain plain strings:

### Rules

- NEVER create enums for helpers
- NEVER create wrapper functions for helpers
- NEVER create constants for helpers
- Users pass them directly as strings

### Rationale

CSS helpers are part of Bulma's language, not Bulma-Templ's API. We expose Bulma 1:1.

### ✅ Correct

```go
@Box(BoxProps{
    Attr: templ.Attributes{
        "class": "has-text-centered p-5 m-3",
    },
}, ...)
```

### ❌ Wrong

```go
// DO NOT create helper enums
type Spacing int
const (
    Spacing1 Spacing = 1
    Spacing2 Spacing = 2
)

// DO NOT create helper functions
func HasTextCentered() string { return "has-text-centered" }
```

---

## 10. Flow Control & Order Freedom

Bulma allows flexible block order (e.g., pagination: prev / next / list in any order).

### Rules

- Components MUST NOT hardcode block order
- Structural blocks are separate templ templates
- User controls final composition

### Example

```go
// ✅ Correct: User controls order
@Pagination(props, elements.Items{
    PaginationPrev(...),
    PaginationList(...),
    PaginationNext(...),
})

// ❌ Wrong: Component hardcodes order
templ Pagination(props) {
    <div class="pagination">
        @PaginationPrev()  // forced order
        @PaginationList()
        @PaginationNext()
    </div>
}
```

---

## 11. Kitchen Sink as Authority

The file `examples/kitchensink/kitchensink.templ` serves as:

- **Authoritative usage examples** for all components
- **Inventory** of the complete component library
- **Reference implementation** for composition patterns

### Rules

- When in doubt about usage, check Kitchen Sink first
- Kitchen Sink examples are normative
- All components MUST have Kitchen Sink examples

---

## 12. HTML is Idiomatic

Direct HTML in templates is allowed and encouraged:

### Rules

- Not everything needs to be a component
- Inline `<div>`, `<span>`, `<p>` are idiomatic
- Bulma-Templ does NOT replace HTML
- Components are for Bulma-specific patterns only

### ✅ Correct

```go
templ MyPage() {
    @Section(props, elements.Items{
        elements.Html(`
            <div class="content">
                <h1>Welcome</h1>
                <p>Direct HTML is fine here.</p>
            </div>
        `),
    })
}
```

---

## 13. Canonical Goals

This design system:

### ❌ Does NOT

- Abstract Bulma
- Hide HTML
- Infer behavior or state
- Enforce app structure
- Provide framework features

### ✅ Does

- Map Bulma 1:1
- Stay explicit and predictable
- Enable composition
- Maintain OSS-grade quality
- Serve as design system reference

---

## 14. Related Documentation

This document is part of a larger documentation set:

- **[LLM_INSTRUCTIONS.md](LLM_INSTRUCTIONS.md)** — Primary guidance for LLMs working with the codebase
- **[COMMENT_STYLE.md](COMMENT_STYLE.md)** — Comment formatting rules and standards
- **[CANONICAL_PROJECT_CONTEXT.md](CANONICAL_PROJECT_CONTEXT.md)** — Project structure and patterns
- **[CONTRIBUTING.md](CONTRIBUTING.md)** — Contribution guidelines and philosophy

All documentation files are normative and MUST be followed.

---

## Status

All architectural decisions in this document are **FINAL**:

- Architecture
- Composition model
- Class handling
- Default state semantics
- Atomic design levels
- Props patterns
- CSS helpers policy

Any deviation from this document is an **architectural error** and MUST be corrected.
