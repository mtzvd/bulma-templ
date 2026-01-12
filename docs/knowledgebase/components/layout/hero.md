---
title: "Hero"
description: "Bulma hero layout component."
category: "Layout"
atomicLevel: "ORGANISM"
bulmaClass: ".hero"
---

# Hero

**Atomic Level:** `ORGANISM`

Bulma hero layout component.

## Description

Hero defines a large layout wrapper used for prominent

page sections such as headers and promotional blocks.

This component is structural only and does NOT manage

content, behavior, or internal composition.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Color` | `string` | Color defines the Bulma color modifier (is-primary, is-link, is-dark, etc.). |
| `Size` | `string` | Size defines the Bulma size modifier (is-small, is-medium, is-large, is-fullheight). |
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.hero` container. |

## Usage

```go
@layout.Hero(
    layout.HeroProps{},
    layout.Items{layout.Html("...")},
)
```

## Bulma Reference

[Bulma Hero Documentation](https://bulma.io/documentation/layout/hero/)

## Source

**File:** `layout/hero.templ`
