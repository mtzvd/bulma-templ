---
title: "HeroBody"
description: "main content section of Hero."
category: "Layout"
atomicLevel: "MOLECULE"
bulmaClass: ".hero-body"
---

# HeroBody

**Atomic Level:** `MOLECULE`

main content section of Hero.

## Description

HeroBody defines the primary content container of Hero.

It does NOT impose layout rules or manage content behavior.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.hero-body` container. |

## Usage

```go
@layout.HeroBody(
    layout.HeroBodyProps{},
    layout.Items{layout.Html("...")},
)
```

## Bulma Reference

[Bulma HeroBody Documentation](https://bulma.io/documentation/layout/hero-body/)

## Source

**File:** `layout/hero.templ`
