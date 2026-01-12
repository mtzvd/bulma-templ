---
title: "SkeletonBlock"
description: ""
category: "Elements"
atomicLevel: "ATOM"
bulmaClass: ".skeleton-block"
---

# SkeletonBlock

**Atomic Level:** `ATOM`

## Description

SkeletonBlock defines visual wrapper that applies

a skeleton loading style to its content.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.skeleton-block` element. |

## Usage

```go
@elements.SkeletonBlock(
    elements.SkeletonBlockProps{},
    elements.Items{elements.Html("...")},
)
```

## Bulma Reference

[Bulma SkeletonBlock Documentation](https://bulma.io/documentation/elements/skeleton-block/)

## Source

**File:** `elements/skeleton.templ`
