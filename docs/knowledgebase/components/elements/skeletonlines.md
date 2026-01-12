---
title: "SkeletonLines"
description: ""
category: "Elements"
atomicLevel: "ATOM"
bulmaClass: ".skeleton-lines"
---

# SkeletonLines

**Atomic Level:** `ATOM`

## Description

SkeletonLines defines wrapper used to apply

skeleton loading styles to text-like content.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.skeleton-lines` container. |

## Usage

```go
@elements.SkeletonLines(
    elements.SkeletonLinesProps{},
    elements.Items{elements.Html("...")},
)
```

## Bulma Reference

[Bulma SkeletonLines Documentation](https://bulma.io/documentation/elements/skeleton-lines/)

## Source

**File:** `elements/skeleton.templ`
