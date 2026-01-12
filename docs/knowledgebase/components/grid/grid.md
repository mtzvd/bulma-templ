---
title: "Grid"
description: "Bulma Smart Grid container."
category: "Grid"
atomicLevel: "ORGANISM"
bulmaClass: ".grid"
---

# Grid

**Atomic Level:** `ORGANISM`

Bulma Smart Grid container.

## Description

Grid renders the Bulma `.grid` container,

which is based on CSS Grid and provides

automatic responsive behavior.

Grid does not define column sizes explicitly.

All layout configuration is done via Bulma

utility classes passed through Attr.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.grid` container.  Use this for: - grid modifiers (is-col-min-*, is-gap-*) - responsive classes - layout helpers |

## Usage

```go
@grid.Grid(
    grid.GridProps{},
    grid.Items{grid.Html("...")},
)
```

## Bulma Reference

[Bulma Grid Documentation](https://bulma.io/documentation/grid/grid/)

## Source

**File:** `grid/grid.templ`
