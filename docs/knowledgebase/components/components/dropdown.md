---
title: "Dropdown"
description: ""
category: "Components"
atomicLevel: "ORGANISM"
bulmaClass: ".dropdown"
---

# Dropdown

**Atomic Level:** `ORGANISM`

## Description

Dropdown defines composite UI component used to display

contextual menus. It does not manage state or behavior;

state must be controlled externally (e.g. via Alpine.js).

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Active` | `bool` | Active applies the `is-active` modifier. |
| `Hoverable` | `bool` | Hoverable enables hover-based activation (`is-hoverable`). |
| `Right` | `bool` | Right aligns the dropdown menu to the right (`is-right`). |
| `Up` | `bool` | Up opens the dropdown upwards (`is-up`). |
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.dropdown` container. |

## Usage

```go
@components.Dropdown(
    components.DropdownProps{},
    components.Items{components.Html("...")},
)
```

## Bulma Reference

[Bulma Dropdown Documentation](https://bulma.io/documentation/components/dropdown/)

## Source

**File:** `components/dropdown.templ`
