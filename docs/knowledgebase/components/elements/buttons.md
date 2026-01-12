---
title: "Buttons"
description: "Bulma buttons container."
category: "Elements"
atomicLevel: "MOLECULE"
bulmaClass: ".buttons"
---

# Buttons

**Atomic Level:** `MOLECULE`

Bulma buttons container.

## Description

Buttons renders a `.buttons` wrapper used

to group multiple Button components.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `HasAddons` | `bool` | HasAddons applies the `has-addons` modifier. |
| `Align` | `string` | Align defines alignment modifier (is-centered, is-right). |
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes. |

## Usage

```go
@elements.Buttons(
    elements.ButtonsProps{},
    elements.Items{elements.Html("...")},
)
```

## Bulma Reference

[Bulma Buttons Documentation](https://bulma.io/documentation/elements/buttons/)

## Source

**File:** `elements/button.templ`
