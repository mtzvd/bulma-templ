---
title: "DropdownItem"
description: ""
category: "Components"
atomicLevel: "ATOM"
bulmaClass: ".dropdown-item"
---

# DropdownItem

**Atomic Level:** `ATOM`

## Description

DropdownItem is always rendered as an <a> element.

If Href is empty, the item is treated as an action item

and can be handled via JS/Alpine.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Href` | `string` | Href defines the link URL. If empty, the item is treated as an action item. |
| `Active` | `bool` | Active applies the `is-active` modifier. |
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the item element. |

## Usage

```go
@components.DropdownItem(
    components.DropdownItemProps{},
    components.Items{components.Html("...")},
)
```

## Bulma Reference

[Bulma DropdownItem Documentation](https://bulma.io/documentation/components/dropdown-item/)

## Source

**File:** `components/dropdown.templ`
