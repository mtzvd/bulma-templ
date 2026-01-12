---
title: "Icon"
description: "Bulma icon element."
category: "Elements"
atomicLevel: "ATOM"
bulmaClass: ".icon"
---

# Icon

**Atomic Level:** `ATOM`

Bulma icon element.

## Description

Icon renders a Bulma `.icon` wrapper with a nested <i> element.

This component is intentionally minimal and does not manage text

or layout composition.

For icon + text combinations, use explicit composition with

the Bulma `.icon-text` pattern.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `IconClass` | `string` | IconClass defines the icon CSS classes (e.g. "fas fa-home"). |
| `Size` | `string` | Size defines the icon size modifier (is-small, is-medium, is-large). |
| `Color` | `string` | Color defines the text color helper (has-text-primary, has-text-danger, etc.). |
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.icon` wrapper. |

## Usage

```go
@elements.Icon(
    elements.IconProps{},
)
```

## Bulma Reference

[Bulma Icon Documentation](https://bulma.io/documentation/elements/icon/)

## Source

**File:** `elements/icon.templ`
