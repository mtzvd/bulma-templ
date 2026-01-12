---
title: "Footer"
description: "Bulma footer layout component."
category: "Layout"
atomicLevel: "ORGANISM"
bulmaClass: ".footer"
---

# Footer

**Atomic Level:** `ORGANISM`

Bulma footer layout component.

## Description

Footer defines the bottom section of a page

using the Bulma `.footer` class.

This component is purely structural and does NOT

impose any internal layout or content structure.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.footer` element. |

## Usage

```go
@layout.Footer(
    layout.FooterProps{},
    layout.Items{layout.Html("...")},
)
```

## Bulma Reference

[Bulma Footer Documentation](https://bulma.io/documentation/layout/footer/)

## Source

**File:** `layout/footer.templ`
