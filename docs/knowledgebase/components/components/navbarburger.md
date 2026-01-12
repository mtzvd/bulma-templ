---
title: "NavbarBurger"
description: "burger button for mobile navigation."
category: "Components"
atomicLevel: "MOLECULE"
bulmaClass: ".navbar-burger"
---

# NavbarBurger

**Atomic Level:** `MOLECULE`

burger button for mobile navigation.

## Description

NavbarBurger controls the visibility of NavbarMenu.

Behavior must be implemented externally.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Active` | `bool` | Active applies the `is-active` modifier. |
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.navbar-burger` element. |

## Usage

```go
@components.NavbarBurger(
    components.NavbarBurgerProps{},
)
```

## Bulma Reference

[Bulma NavbarBurger Documentation](https://bulma.io/documentation/components/navbar-burger/)

## Source

**File:** `components/navbar.templ`
