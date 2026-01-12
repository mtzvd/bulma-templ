---
title: "TableContainer"
description: ""
category: "Elements"
atomicLevel: "ATOM"
bulmaClass: ".table-container"
---

# TableContainer

**Atomic Level:** `ATOM`

## Description

TableContainer is used to wrap tables

to enable horizontal scrolling.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.table-container` element. |

## Usage

```go
@elements.TableContainer(
    elements.TableContainerProps{},
    elements.Items{elements.Html("...")},
)
```

## Bulma Reference

[Bulma TableContainer Documentation](https://bulma.io/documentation/elements/table-container/)

## Source

**File:** `elements/table.templ`
