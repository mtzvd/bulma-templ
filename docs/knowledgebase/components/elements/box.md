---
title: "Box"
description: "Bulma box element."
category: "Elements"
atomicLevel: "ATOM"
bulmaClass: ".box"
---

# Box

**Atomic Level:** `ATOM`

Bulma box element.

## Description

Box defines a visual container that provides padding,

border radius, and shadow using the Bulma `.box` class.

This component is content-based and does not manage

layout, behavior, or semantics.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.box` container. |

## Usage

```go
@elements.Box(
    elements.BoxProps{},
    elements.Items{elements.Html("...")},
)
```

## Bulma Reference

[Bulma Box Documentation](https://bulma.io/documentation/elements/box/)

## Source

**File:** `elements/box.templ`
