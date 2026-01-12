---
title: "DropdownMenu"
description: ""
category: "Components"
atomicLevel: "MOLECULE"
bulmaClass: ".dropdown-menu"
---

# DropdownMenu

**Atomic Level:** `MOLECULE`

## Description

DropdownMenu wraps dropdown items and does not manage state.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.dropdown-menu` container. |

## Usage

```go
@components.DropdownMenu(
    components.DropdownMenuProps{},
    components.Items{components.Html("...")},
)
```

## Bulma Reference

[Bulma DropdownMenu Documentation](https://bulma.io/documentation/components/dropdown-menu/)

## Source

**File:** `components/dropdown.templ`
