---
title: "HeroFoot"
description: "bottom section of Hero."
category: "Layout"
atomicLevel: "MOLECULE"
bulmaClass: ".hero-foot"
---

# HeroFoot

**Atomic Level:** `MOLECULE`

bottom section of Hero.

## Description

HeroFoot is used for secondary navigation

or additional actions below the Hero body.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.hero-foot` container. |

## Usage

```go
@layout.HeroFoot(
    layout.HeroFootProps{},
    layout.Items{layout.Html("...")},
)
```

## Bulma Reference

[Bulma HeroFoot Documentation](https://bulma.io/documentation/layout/hero-foot/)

## Source

**File:** `layout/hero.templ`
