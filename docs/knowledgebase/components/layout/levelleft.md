---
title: "LevelLeft"
description: "left-aligned section of Level."
category: "Layout"
atomicLevel: "MOLECULE"
bulmaClass: ".level-left"
---

# LevelLeft

**Atomic Level:** `MOLECULE`

left-aligned section of Level.

## Description

LevelLeft groups items aligned to the left side

of the Level container.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.level-left` container. |

## Usage

```go
@layout.LevelLeft(
    layout.LevelLeftProps{},
    layout.Items{layout.Html("...")},
)
```

## Bulma Reference

[Bulma LevelLeft Documentation](https://bulma.io/documentation/layout/level-left/)

## Source

**File:** `layout/level.templ`
