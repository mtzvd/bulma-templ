---
title: "Menu"
description: ""
category: "Components"
atomicLevel: "ORGANISM"
bulmaClass: ".menu"
---

# Menu

**Atomic Level:** `ORGANISM`

## Description

Menu is used for vertical navigation and grouped links.

It does not manage state or impose internal structure.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.menu` container. |

## Usage

```go
@components.Menu(
    components.MenuProps{},
    components.Items{components.Html("...")},
)
```

## Bulma Reference

[Bulma Menu Documentation](https://bulma.io/documentation/components/menu/)

## Source

**File:** `components/menu.templ`
