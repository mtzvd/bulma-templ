---
title: "MenuItem"
description: ""
category: "Components"
atomicLevel: "ATOM"
bulmaClass: ".menu-item"
---

# MenuItem

**Atomic Level:** `ATOM`

## Description

MenuItem is always rendered as an <a> element.

Active state is controlled externally.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Href` | `string` | Href defines the link URL. If empty, the item is treated as an action item. |
| `Active` | `bool` | Active applies the `is-active` modifier. |
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the <a> element. |

## Usage

```go
@components.MenuItem(
    components.MenuItemProps{},
    components.Items{components.Html("...")},
)
```

## Bulma Reference

[Bulma MenuItem Documentation](https://bulma.io/documentation/components/menu-item/)

## Source

**File:** `components/menu.templ`
