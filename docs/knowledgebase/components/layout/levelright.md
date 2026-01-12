---
title: "LevelRight"
description: "right-aligned section of Level."
category: "Layout"
atomicLevel: "MOLECULE"
bulmaClass: ".level-right"
---

# LevelRight

**Atomic Level:** `MOLECULE`

right-aligned section of Level.

## Description

LevelRight groups items aligned to the right side

of the Level container.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.level-right` container. |

## Usage

```go
@layout.LevelRight(
    layout.LevelRightProps{},
    layout.Items{layout.Html("...")},
)
```

## Bulma Reference

[Bulma LevelRight Documentation](https://bulma.io/documentation/layout/level-right/)

## Source

**File:** `layout/level.templ`
