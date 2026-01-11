# Comment Style Guide — Bulma-Templ Design System (templ / Go)

This document defines the canonical comment style for all components in the Bulma-Templ Design System.

This style is frozen for v1.0.

---

## Core Rules

### 1. Use `//` only (NEVER `/* */`)

```go
// Correct
// Button — Bulma button component.

// WRONG
/* Button — Bulma button component. */
```

### 2. No blank comment lines

```go
// Correct
// Button renders a Bulma button.
templ Button(props ButtonProps, content Items) {

// WRONG
// Button renders a Bulma button.

templ Button(props ButtonProps, content Items) {
```

### 3. All field comments end with a period

```go
// Correct
// Size defines the Bulma size modifier.
Size string

// WRONG
// Size defines the Bulma size modifier
Size string
```

---

## Component-level Comments

Every component MUST include the following metadata at the top of its declaration:

1. Component name and brief description
2. Explicit atomic level
3. Detailed responsibility definition
4. Explicit non-responsibility (when relevant)

### Canonical Form

```go
// Button — Bulma button component.
// Atomic level: ATOM
//
// Button defines a clickable UI element styled with Bulma.
// It does NOT manage state, routing, or business logic.
```

### Title Line Format

**Pattern**: `// ComponentName — Brief description.`

**Rules**:

- Use em dash (—) not hyphen (-)
- End with a period
- Keep description under 10 words
- Use "Bulma" when referring to Bulma-specific features

**Examples**:

```go
✅ // Button — Bulma button component.
✅ // Navbar — Bulma navigation bar component.
✅ // CardHeaderIcon — icon/action inside CardHeader.

❌ // Button - Bulma button component    (hyphen instead of em dash)
❌ // Button — Bulma button              (no period)
❌ // Button — this is a button          (not Bulma-specific)
```

### Atomic Level Line

**MUST be exact**:

```go
// Atomic level: ATOM
// Atomic level: MOLECULE  
// Atomic level: ORGANISM
```

**No variants allowed**:

```go
❌ // Atomic: ATOM
❌ // Level: ATOM
❌ // Atomic Level: ATOM (wrong capitalization)
```

### Description Section

**Rules**:

- Start with a verb (defines, renders, wraps, applies, groups, composes)
- Use article "a" or "an" for singular nouns
- Multi-sentence descriptions are allowed
- Explain what the component DOES, not what it IS

**Allowed verbs** (STRICT):

```text
defines
renders
wraps
applies
groups
composes
```

**Forbidden verbs**:

```text
is a
represents
responsible for
manages (except in negation)
handles
abstracts
simplifies
```

**Article usage**:

```go
✅ // Box defines a visual container that provides padding.
✅ // Hero defines a large layout wrapper used for prominent sections.

❌ // Box defines visual container that provides padding.
❌ // Hero defines large layout wrapper used for sections.
```

**Examples**:

```go
// Simple ATOM
// Icon — Bulma icon wrapper.
// Atomic level: ATOM
//
// Icon renders a Bulma `.icon` container for icon fonts.

// Complex ATOM with negation
// Button — Bulma button component.
// Atomic level: ATOM
//
// Button defines a clickable UI element styled with Bulma.
// It does NOT manage state, routing, or business logic.

// MOLECULE with context
// Field — Bulma form field container.
// Atomic level: MOLECULE
//
// Field groups label, control, and help text into
// a single form field structure.
//
// This component does NOT validate input or manage form state.

// ORGANISM with multiple sentences
// Navbar — Bulma navigation bar component.
// Atomic level: ORGANISM
//
// Navbar renders a responsive navigation bar with brand,
// menu toggle, and navigation items. It does NOT manage
// routing, authentication, or mobile menu state.
```

---

## Props Struct Comments

### Structure

Every Props struct MUST have:

1. A summary line before the struct
2. Individual field comments
3. Aligned formatting

### Summary Line

Describe the purpose of the Props struct in one sentence:

```go
✅ // ButtonProps defines configuration for the Button component.

❌ // Button properties
❌ type ButtonProps struct {  // (no comment)
```

### Field Comments

**Rules**:

- Every field MUST have a comment
- Start with a verb: defines, controls, applies, contains, enables, disables
- End with a period
- Align multi-line comments

**Standard verbs for common fields**:

```go
// Color defines the Bulma color modifier...
// Size defines the Bulma size modifier...
// Attr contains additional HTML attributes...
// Disabled disables the component...
// Active applies the active state...
```

### Attr Field Pattern (Canonical)

```go
// Attr contains additional HTML attributes for the `.component-class` element.
Attr templ.Attributes
```

**Examples by component type**:

```go
// For container components
// Attr contains additional HTML attributes for the `.box` container.

// For wrapper components  
// Attr contains additional HTML attributes for the `.card-header` element.

// For native elements
// Attr contains additional HTML attributes for the <input> element.
```

### Complete Examples

**Simple ATOM**:

```go
// TagProps defines configuration for the Tag component.
type TagProps struct {
    // Color defines the Bulma color modifier
    // (is-primary, is-info, is-success, is-warning, is-danger).
    Color string

    // Size defines the Bulma size modifier
    // (is-small, is-medium, is-large).
    Size string

    // Rounded applies the `is-rounded` modifier.
    Rounded bool

    // Attr contains additional HTML attributes for the `.tag` element.
    Attr templ.Attributes
}
```

**Complex MOLECULE**:

```go
// FieldProps defines configuration for the Field component.
type FieldProps struct {
    // Grouped applies the `is-grouped` modifier for horizontal layout.
    Grouped bool

    // GroupedMultiline applies the `is-grouped-multiline` modifier.
    GroupedMultiline bool

    // Addons applies the `has-addons` modifier for connected controls.
    Addons bool

    // Attr contains additional HTML attributes for the `.field` container.
    Attr templ.Attributes
}
```

**Component with multiple Attr fields**:

```go
// SelectProps defines configuration for the Select component.
type SelectProps struct {
    // Color defines the select color modifier
    // (e.g. is-primary, is-danger).
    Color string

    // Disabled disables the native <select> element.
    Disabled bool

    // Attr contains additional HTML attributes
    // for the native <select> element.
    //
    // Use this for:
    //   - name
    //   - value
    //   - data-*
    //   - x-model
    //   - aria-*
    Attr templ.Attributes

    // WrapperAttr contains additional HTML attributes
    // for the Bulma `.select` wrapper.
    //
    // Use this for:
    //   - layout-related classes
    //   - visibility helpers
    //   - Alpine.js state (x-data, x-show, etc.)
    WrapperAttr templ.Attributes
}
```

### Sub-component Pattern

Sub-components (parts of ORGANISM) may have minimal descriptions:

```go
// CardHeaderTitle — title inside CardHeader.
// Atomic level: MOLECULE
//
// CardHeaderTitle displays the title text in CardHeader.
type CardHeaderTitleProps struct {
    // Attr contains additional HTML attributes for the `.card-header-title` element.
    Attr templ.Attributes
}
```

---

## Templ Function Comments

Every templ function MUST have a comment:

```go
// ComponentName renders [brief description].
templ ComponentName(props ComponentNameProps, content Items) {
```

**Examples**:

```go
// Button renders a Bulma button.
templ Button(props ButtonProps, content Items) {

// Field renders a Bulma form field container.
templ Field(props FieldProps, content Items) {

// CardHeader renders a Bulma `.card-header`.
templ CardHeader(props CardHeaderProps, content Items) {
```

---

## Special Cases

### Components Without Items

```go
// Image renders a Bulma image component.
templ Image(props ImageProps) {
```

### Components With Context

When a component requires explanation:

```go
// File renders the Bulma `.file` container.
//
// FileLabel must be provided via content.
templ File(props FileProps, content Items) {
```

### Constants

Package-level constants for special values:

```go
// ProgressUnset indicates that a progress value is not set.
// Use this constant when you want to omit the attribute.
const ProgressUnset = -1
```

---

## Alignment & Formatting

### Field Comment Alignment

When field comments are short, align them:

```go
type SimpleProps struct {
    Color string           // Bulma color modifier (is-primary, is-link, etc.).
    Size  string           // Bulma size modifier (is-small, is-medium, is-large).
    Attr  templ.Attributes // Additional HTML attributes.
}
```

When comments are long, use separate lines:

```go
type ComplexProps struct {
    // Color defines the Bulma color modifier
    // (is-primary, is-link, is-info, is-success, etc.).
    Color string

    // Size defines the Bulma size modifier
    // (is-small, is-medium, is-large).
    Size string

    // Attr contains additional HTML attributes for the component.
    Attr templ.Attributes
}
```

### Multi-line Comment Indentation

```go
// Attr contains additional HTML attributes
// for the `.card-header-icon` element.
Attr templ.Attributes
```

**Align continuation lines with the start of the first line**.

---

## Examples: Before and After

### ❌ Before (Incorrect)

```go
/* Card component */
// Atomic Level: ORGANISM
type CardProps struct {
    Attr templ.Attributes  // attrs
}

func Card(props CardProps, content Items) {
```

### ✅ After (Correct)

```go
// Card — Bulma card component.
// Atomic level: ORGANISM
//
// Card defines a composite UI component used to display
// grouped content such as articles, previews, or actions.
//
// Card does not impose internal structure and must be
// composed explicitly using its subcomponents.
type CardProps struct {
    // Attr contains additional HTML attributes for the `.card` container.
    Attr templ.Attributes
}

// Card renders a Bulma `.card` container.
templ Card(props CardProps, content Items) {
```

---

## Checklist for New Components

Before submitting a new component, verify:

- [ ] Title line uses em dash (—) and ends with period
- [ ] Atomic level is explicitly stated (ATOM/MOLECULE/ORGANISM)
- [ ] Description uses allowed verbs (defines, renders, etc.)
- [ ] Article "a"/"an" is present for singular nouns
- [ ] Props struct has a summary comment
- [ ] All fields have comments ending with periods
- [ ] Attr comment follows canonical pattern
- [ ] Templ function has a brief comment
- [ ] No blank comment lines before templ
- [ ] Multi-line comments are properly aligned
- [ ] `//` is used consistently (never `/* */`)

---

## Freeze Policy

From v1.0 onward:

- **New components MUST follow this guide**
- **Existing comments are only changed for**:
  - Correctness (wrong atomic level, missing descriptions)
  - Clarity (ambiguous wording)
  - Documentation expansion (adding missing details)

**Style consistency changes** (spacing, alignment) should be batched and not done ad-hoc.
