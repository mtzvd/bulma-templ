---
title: "NavbarMenu"
description: "collapsible menu section."
category: "Components"
atomicLevel: "MOLECULE"
bulmaClass: ".navbar-menu"
---

# NavbarMenu

**Atomic Level:** `MOLECULE`

collapsible menu section.

## Description

NavbarMenu defines the structural container for navbar items.

It does NOT manage visibility, animation, or interaction logic.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Active` | `bool` | Active applies the `is-active` modifier. |
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.navbar-menu` container. |

## Usage

```go
@components.NavbarMenu(
    components.NavbarMenuProps{},
    components.Items{components.Html("...")},
)
```

## Bulma Reference

[Bulma NavbarMenu Documentation](https://bulma.io/documentation/components/navbar-menu/)

## Source

**File:** `components/navbar.templ`
