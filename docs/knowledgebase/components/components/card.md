---
title: "Card"
description: "Bulma card component."
category: "Components"
atomicLevel: "ORGANISM"
bulmaClass: ".card"
---

# Card

**Atomic Level:** `ORGANISM`

Bulma card component.

## Description

Card defines composite UI component used to display

grouped content such as articles, previews, or actions.

Card does not impose internal structure and must be

composed explicitly using its subcomponents.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.card` container. |

## Usage

```go
@components.Card(
    components.CardProps{},
    components.Items{components.Html("...")},
)
```

## Bulma Reference

[Bulma Card Documentation](https://bulma.io/documentation/components/card/)

## Source

**File:** `components/card.templ`
