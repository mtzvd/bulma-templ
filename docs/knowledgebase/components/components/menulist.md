---
title: "MenuList"
description: ""
category: "Components"
atomicLevel: "MOLECULE"
bulmaClass: ".menu-list"
---

# MenuList

**Atomic Level:** `MOLECULE`

## Description

MenuList wraps menu items and supports

nested lists via standard <ul> nesting.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.menu-list` container. |

## Usage

```go
@components.MenuList(
    components.MenuListProps{},
    components.Items{components.Html("...")},
)
```

## Bulma Reference

[Bulma MenuList Documentation](https://bulma.io/documentation/components/menu-list/)

## Source

**File:** `components/menu.templ`
