---
title: "Label"
description: "Bulma form label element."
category: "Form"
atomicLevel: "ATOM"
bulmaClass: ".label"
---

# Label

**Atomic Level:** `ATOM`

Bulma form label element.

## Description

Label renders a textual label associated with

a form control. It does not manage layout,

accessibility relations or form state.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Text` | `string` | Text defines the label content. |
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the <label> element. |

## Usage

```go
@form.Label(
    form.LabelProps{},
    form.Items{form.Html("...")},
)
```

## Bulma Reference

[Bulma Label Documentation](https://bulma.io/documentation/form/label/)

## Source

**File:** `form/label.templ`
