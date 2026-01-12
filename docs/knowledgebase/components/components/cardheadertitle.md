---
title: "CardHeaderTitle"
description: "title inside CardHeader."
category: "Components"
atomicLevel: "MOLECULE"
bulmaClass: ".card-header-title"
---

# CardHeaderTitle

**Atomic Level:** `MOLECULE`

title inside CardHeader.

## Description

CardHeaderTitle displays the title text in CardHeader.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.card-header-title` element. |

## Usage

```go
@components.CardHeaderTitle(
    components.CardHeaderTitleProps{},
    components.Items{components.Html("...")},
)
```

## Bulma Reference

[Bulma CardHeaderTitle Documentation](https://bulma.io/documentation/components/card-header-title/)

## Source

**File:** `components/card.templ`
