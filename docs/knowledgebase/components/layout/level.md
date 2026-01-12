---
title: "Level"
description: "Bulma level layout component."
category: "Layout"
atomicLevel: "ORGANISM"
bulmaClass: ".level"
---

# Level

**Atomic Level:** `ORGANISM`

Bulma level layout component.

## Description

Level is used to horizontally align elements

to the left, right, or center using the Bulma `.level` pattern.

This component does NOT manage content behavior or state.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Mobile` | `bool` | Mobile applies the `is-mobile` modifier, preventing collapsing on mobile devices. |
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.level` container. |

## Usage

```go
@layout.Level(
    layout.LevelProps{},
    layout.Items{layout.Html("...")},
)
```

## Bulma Reference

[Bulma Level Documentation](https://bulma.io/documentation/layout/level/)

## Source

**File:** `layout/level.templ`
