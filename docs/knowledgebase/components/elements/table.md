---
title: "Table"
description: ""
category: "Elements"
atomicLevel: "ATOM"
bulmaClass: ".table"
---

# Table

**Atomic Level:** `ATOM`

## Description

Table defines structural and typographic element.

It does not manage data, rows, or cells.

All table markup is provided via content.

Supported modifiers:

is-bordered, is-striped, is-narrow,

is-hoverable, is-fullwidth

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Modifiers` | `string` | Modifiers contains Bulma table modifiers. Example: "is-striped is-hoverable is-fullwidth" |
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the <table> element. |

## Usage

```go
@elements.Table(
    elements.TableProps{},
    elements.Items{elements.Html("...")},
)
```

## Bulma Reference

[Bulma Table Documentation](https://bulma.io/documentation/elements/table/)

## Source

**File:** `elements/table.templ`
