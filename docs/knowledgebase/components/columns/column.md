---
title: "Column"
description: "Bulma column element."
category: "Columns"
atomicLevel: "MOLECULE"
bulmaClass: ".column"
---

# Column

**Atomic Level:** `MOLECULE`

Bulma column element.

## Description

Column defines a single column inside a Columns container.

It controls width, offset and responsiveness.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Size` | `string` | Size defines the column width (e.g. is-half, is-one-third). |
| `Narrow` | `bool` | Narrow makes the column only as wide as its content. |
| `Offset` | `string` | Offset defines column offset (e.g. is-offset-one-quarter). |
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.column` element. |

## Usage

```go
@columns.Column(
    columns.ColumnProps{},
    columns.Items{columns.Html("...")},
)
```

## Bulma Reference

[Bulma Column Documentation](https://bulma.io/documentation/columns/column/)

## Source

**File:** `columns/columns.templ`
