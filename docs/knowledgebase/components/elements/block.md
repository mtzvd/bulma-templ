---
title: "Block"
description: "Bulma block element."
category: "Elements"
atomicLevel: "ATOM"
bulmaClass: ".block"
---

# Block

**Atomic Level:** `ATOM`

Bulma block element.

## Description

Block defines a simple spacing container that adds

vertical margin between elements using the

Bulma `.block` class.

This component does not manage layout or content

semantics and serves as a low-level UI primitive.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.block` container. |

## Usage

```go
@elements.Block(
    elements.BlockProps{},
    elements.Items{elements.Html("...")},
)
```

## Bulma Reference

[Bulma Block Documentation](https://bulma.io/documentation/elements/block/)

## Source

**File:** `elements/block.templ`
