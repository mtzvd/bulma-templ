---
title: "Input"
description: "Bulma text input element."
category: "Form"
atomicLevel: "ATOM"
bulmaClass: ".input"
---

# Input

**Atomic Level:** `ATOM`

Bulma text input element.

## Description

Input renders a Bulma `.input` form control.

It is responsible only for visual appearance

and HTML attributes, but does not manage layout,

validation logic or value binding.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Type` | `string` | Type defines the input type attribute (e.g. text, email, password). |
| `Value` | `string` | Value defines the input value. |
| `Placeholder` | `string` | Placeholder defines the input placeholder text. |
| `Color` | `string` | Color defines the input color modifier (e.g. is-primary, is-danger). |
| `Size` | `string` | Size defines the input size modifier (is-small, is-medium, is-large). |
| `Disabled` | `bool` | Disabled disables the input. |
| `ReadOnly` | `bool` | ReadOnly makes the input read-only. |
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the <input> element. |

## Usage

```go
@form.Input(
    form.InputProps{},
)
```

## Bulma Reference

[Bulma Input Documentation](https://bulma.io/documentation/form/input/)

## Source

**File:** `form/input.templ`
