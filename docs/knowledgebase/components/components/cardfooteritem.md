---
title: "CardFooterItem"
description: "item inside CardFooter."
category: "Components"
atomicLevel: "MOLECULE"
bulmaClass: ".card-footer-item"
---

# CardFooterItem

**Atomic Level:** `MOLECULE`

item inside CardFooter.

## Description

CardFooterItem defines a single action or link inside CardFooter.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.card-footer-item` element. |

## Usage

```go
@components.CardFooterItem(
    components.CardFooterItemProps{},
    components.Items{components.Html("...")},
)
```

## Bulma Reference

[Bulma CardFooterItem Documentation](https://bulma.io/documentation/components/card-footer-item/)

## Source

**File:** `components/card.templ`
