---
title: "Field"
description: "Bulma form field wrapper."
category: "Form"
atomicLevel: "MOLECULE"
bulmaClass: ".field"
---

# Field

**Atomic Level:** `MOLECULE`

Bulma form field wrapper.

## Description

Field defines grouping related form elements

such as Label, Control and Help into a single logical block.

It defines the base structure for vertical and horizontal

form layouts, but does not manage form state or validation.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Grouped` | `bool` | Grouped adds the is-grouped modifier. |
| `Horizontal` | `bool` | Horizontal adds the is-horizontal modifier. |
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.field` container. |

## Usage

```go
@form.Field(
    form.FieldProps{},
    form.Items{form.Html("...")},
)
```

## Bulma Reference

[Bulma Field Documentation](https://bulma.io/documentation/form/field/)

## Source

**File:** `form/field.templ`
