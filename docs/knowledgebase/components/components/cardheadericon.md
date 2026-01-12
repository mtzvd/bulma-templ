---
title: "CardHeaderIcon"
description: "icon/action inside CardHeader."
category: "Components"
atomicLevel: "MOLECULE"
bulmaClass: ".card-header-icon"
---

# CardHeaderIcon

**Atomic Level:** `MOLECULE`

icon/action inside CardHeader.

## Description

CardHeaderIcon is used for action buttons or icons in CardHeader.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.card-header-icon` element. |

## Usage

```go
@components.CardHeaderIcon(
    components.CardHeaderIconProps{},
    components.Items{components.Html("...")},
)
```

## Bulma Reference

[Bulma CardHeaderIcon Documentation](https://bulma.io/documentation/components/card-header-icon/)

## Source

**File:** `components/card.templ`
