---
title: "LevelItem"
description: "individual item inside Level."
category: "Layout"
atomicLevel: "MOLECULE"
bulmaClass: ".level-item"
---

# LevelItem

**Atomic Level:** `MOLECULE`

individual item inside Level.

## Description

LevelItem wraps a single element inside Level

and applies proper alignment and spacing.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.level-item` wrapper. |

## Usage

```go
@layout.LevelItem(
    layout.LevelItemProps{},
    layout.Items{layout.Html("...")},
)
```

## Bulma Reference

[Bulma LevelItem Documentation](https://bulma.io/documentation/layout/level-item/)

## Source

**File:** `layout/level.templ`
