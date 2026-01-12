---
title: "CardFooter"
description: "footer section of Card."
category: "Components"
atomicLevel: "MOLECULE"
bulmaClass: ".card-footer"
---

# CardFooter

**Atomic Level:** `MOLECULE`

footer section of Card.

## Description

CardFooter displays actions or links at the bottom of Card.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.card-footer` container. |

## Usage

```go
@components.CardFooter(
    components.CardFooterProps{},
    components.Items{components.Html("...")},
)
```

## Bulma Reference

[Bulma CardFooter Documentation](https://bulma.io/documentation/components/card-footer/)

## Source

**File:** `components/card.templ`
