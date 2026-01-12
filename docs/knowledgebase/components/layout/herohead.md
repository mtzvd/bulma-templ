---
title: "HeroHead"
description: "top section of Hero."
category: "Layout"
atomicLevel: "MOLECULE"
bulmaClass: ".hero-head"
---

# HeroHead

**Atomic Level:** `MOLECULE`

top section of Hero.

## Description

HeroHead is typically used to place a Navbar

or auxiliary content above the main Hero body.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.hero-head` container. |

## Usage

```go
@layout.HeroHead(
    layout.HeroHeadProps{},
    layout.Items{layout.Html("...")},
)
```

## Bulma Reference

[Bulma HeroHead Documentation](https://bulma.io/documentation/layout/hero-head/)

## Source

**File:** `layout/hero.templ`
