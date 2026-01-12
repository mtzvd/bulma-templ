---
title: "Container"
description: "Bulma container layout wrapper."
category: "Layout"
atomicLevel: "ATOM"
bulmaClass: ".container"
---

# Container

**Atomic Level:** `ATOM`

Bulma container layout wrapper.

## Description

Container constrains the horizontal width of its content

according to Bulma container rules.

This component does NOT manage grid, columns, or spacing.

It is intended to be used as a structural layout wrapper.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Fluid` | `bool` | Fluid applies the `is-fluid` modifier, making the container full-width. |
| `Widescreen` | `bool` | Widescreen applies the `is-widescreen` modifier. |
| `FullHD` | `bool` | FullHD applies the `is-fullhd` modifier. |
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.container` element. |

## Usage

```go
@layout.Container(
    layout.ContainerProps{},
    layout.Items{layout.Html("...")},
)
```

## Bulma Reference

[Bulma Container Documentation](https://bulma.io/documentation/layout/container/)

## Source

**File:** `layout/container.templ`
