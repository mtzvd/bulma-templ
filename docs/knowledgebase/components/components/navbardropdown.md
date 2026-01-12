---
title: "NavbarDropdown"
description: "dropdown menu inside Navbar."
category: "Components"
atomicLevel: "MOLECULE"
bulmaClass: ".navbar-dropdown"
---

# NavbarDropdown

**Atomic Level:** `MOLECULE`

dropdown menu inside Navbar.

## Description

NavbarDropdown — dropdown menu inside Navbar.
Atomic level: MOLECULE

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Right` | `bool` | Right applies the `is-right` modifier. |
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.navbar-dropdown` container. |

## Usage

```go
@components.NavbarDropdown(
    components.NavbarDropdownProps{},
    components.Items{components.Html("...")},
)
```

## Bulma Reference

[Bulma NavbarDropdown Documentation](https://bulma.io/documentation/components/navbar-dropdown/)

## Source

**File:** `components/navbar.templ`
