---
title: "NavbarLink"
description: "clickable dropdown trigger inside navbar-item."
category: "Components"
atomicLevel: "ATOM"
bulmaClass: ".navbar-link"
---

# NavbarLink

**Atomic Level:** `ATOM`

clickable dropdown trigger inside navbar-item.

## Description

NavbarLink represents a `.navbar-link` element.

Used as dropdown trigger with is-arrowless modifier support.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Arrowless` | `bool` | Arrowless removes the dropdown arrow indicator. |
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.navbar-link` element. |

## Usage

```go
@components.NavbarLink(
    components.NavbarLinkProps{},
    components.Items{components.Html("...")},
)
```

## Bulma Reference

[Bulma NavbarLink Documentation](https://bulma.io/documentation/components/navbar-link/)

## Source

**File:** `components/navbar.templ`
