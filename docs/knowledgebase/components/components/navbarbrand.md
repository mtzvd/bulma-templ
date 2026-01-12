---
title: "NavbarBrand"
description: "brand section of Navbar."
category: "Components"
atomicLevel: "MOLECULE"
bulmaClass: ".navbar-brand"
---

# NavbarBrand

**Atomic Level:** `MOLECULE`

brand section of Navbar.

## Description

NavbarBrand typically contains the logo and burger button.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.navbar-brand` container. |

## Usage

```go
@components.NavbarBrand(
    components.NavbarBrandProps{},
    components.Items{components.Html("...")},
)
```

## Bulma Reference

[Bulma NavbarBrand Documentation](https://bulma.io/documentation/components/navbar-brand/)

## Source

**File:** `components/navbar.templ`
