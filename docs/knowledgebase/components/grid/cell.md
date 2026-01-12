---
title: "Cell"
description: "Bulma Smart Grid cell."
category: "Grid"
atomicLevel: "MOLECULE"
bulmaClass: ".cell"
---

# Cell

**Atomic Level:** `MOLECULE`

Bulma Smart Grid cell.

## Description

Cell renders the Bulma `.cell` element,

which defines a single grid item

inside a Grid container.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.cell` element.  Use this for: - cell span modifiers - alignment helpers - responsive behavior |

## Usage

```go
@grid.Cell(
    grid.CellProps{},
    grid.Items{grid.Html("...")},
)
```

## Bulma Reference

[Bulma Cell Documentation](https://bulma.io/documentation/grid/cell/)

## Source

**File:** `grid/grid.templ`
