---
title: "CardContent"
description: "main content area of Card."
category: "Components"
atomicLevel: "MOLECULE"
bulmaClass: ".card-content"
---

# CardContent

**Atomic Level:** `MOLECULE`

main content area of Card.

## Description

CardContent is used to display the main content of Card.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.card-content` container. |

## Usage

```go
@components.CardContent(
    components.CardContentProps{},
    components.Items{components.Html("...")},
)
```

## Bulma Reference

[Bulma CardContent Documentation](https://bulma.io/documentation/components/card-content/)

## Source

**File:** `components/card.templ`
