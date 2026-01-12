---
title: "Navbar"
description: "Bulma navbar layout component."
category: "Components"
atomicLevel: "ORGANISM"
bulmaClass: ".navbar"
---

# Navbar

**Atomic Level:** `ORGANISM`

Bulma navbar layout component.

## Description

Navbar provides the main site navigation structure.

It does NOT manage state, burger behavior, or dropdown logic.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Color` | `string` | Color defines the Bulma color modifier (is-primary, is-dark, is-light, etc.). |
| `Transparent` | `bool` | Transparent applies the `is-transparent` modifier. |
| `Fixed` | `string` | Fixed defines fixed positioning (is-fixed-top or is-fixed-bottom). |
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.navbar` container. |

## Usage

```go
@components.Navbar(
    components.NavbarProps{},
    components.Items{components.Html("...")},
)
```

## Bulma Reference

[Bulma Navbar Documentation](https://bulma.io/documentation/components/navbar/)

## Source

**File:** `components/navbar.templ`
