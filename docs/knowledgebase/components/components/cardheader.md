---
title: "CardHeader"
description: "header section of Card."
category: "Components"
atomicLevel: "MOLECULE"
bulmaClass: ".card-header"
---

# CardHeader

**Atomic Level:** `MOLECULE`

header section of Card.

## Description

CardHeader defines the header container of a Card.

It does NOT manage layout, actions, or state.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Attr` | `templ.Attributes` |  |

## Usage

```go
@components.CardHeader(
    components.CardHeaderProps{},
    components.Items{components.Html("...")},
)
```

## Bulma Reference

[Bulma CardHeader Documentation](https://bulma.io/documentation/components/card-header/)

## Source

**File:** `components/card.templ`
