---
title: "MenuLabel"
description: ""
category: "Components"
atomicLevel: "MOLECULE"
bulmaClass: ".menu-label"
---

# MenuLabel

**Atomic Level:** `MOLECULE`

## Description

MenuLabel is used to display section headers inside Menu.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.menu-label` element. |

## Usage

```go
@components.MenuLabel(
    components.MenuLabelProps{},
    components.Items{components.Html("...")},
)
```

## Bulma Reference

[Bulma MenuLabel Documentation](https://bulma.io/documentation/components/menu-label/)

## Source

**File:** `components/menu.templ`
