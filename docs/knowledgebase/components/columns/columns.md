---
title: "Columns"
description: "Bulma columns layout container."
category: "Columns"
atomicLevel: "ORGANISM"
bulmaClass: ".columns"
---

# Columns

**Atomic Level:** `ORGANISM`

Bulma columns layout container.

## Description

Columns defines a flexbox-based layout system

for arranging content in horizontal columns.

It does not impose any content structure and

serves purely as a layout mechanism.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Multiline` | `bool` | Multiline allows columns to wrap onto multiple lines. |
| `Centered` | `bool` | Centered centers the columns horizontally. |
| `Gapless` | `bool` | Gapless removes the gap between columns. |
| `Variable` | `bool` | Variable enables variable gap sizes. |
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.columns` container. |

## Usage

```go
@columns.Columns(
    columns.ColumnsProps{},
    columns.Items{columns.Html("...")},
)
```

## Bulma Reference

[Bulma Columns Documentation](https://bulma.io/documentation/columns/columns/)

## Source

**File:** `columns/columns.templ`
