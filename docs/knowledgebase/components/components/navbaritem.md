---
title: "NavbarItem"
description: "individual item in navbar (link, div, or dropdown trigger)."
category: "Components"
atomicLevel: "ATOM"
bulmaClass: ".navbar-item"
---

# NavbarItem

**Atomic Level:** `ATOM`

individual item in navbar (link, div, or dropdown trigger).

## Description

NavbarItem represents a `.navbar-item` element.

Can be a link (<a>), div (<div>), or dropdown trigger.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Href` | `string` | Href defines the link target. If empty, renders a <div>; if set, renders an <a>. |
| `Tab` | `bool` | Tab applies the `is-tab` modifier for tab-like styling. |
| `Active` | `bool` | Active applies the `is-active` modifier. |
| `Hoverable` | `bool` | Hoverable applies the `is-hoverable` modifier. |
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.navbar-item` element. |

## Usage

```go
@components.NavbarItem(
    components.NavbarItemProps{},
    components.Items{components.Html("...")},
)
```

## Bulma Reference

[Bulma NavbarItem Documentation](https://bulma.io/documentation/components/navbar-item/)

## Source

**File:** `components/navbar.templ`
