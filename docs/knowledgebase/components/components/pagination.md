---
title: "Pagination"
description: "Bulma pagination container."
category: "Components"
atomicLevel: "ORGANISM"
bulmaClass: ".pagination"
---

# Pagination

**Atomic Level:** `ORGANISM`

Bulma pagination container.

## Description

Pagination renders the root `.pagination` element.

This component is a pure structural container and does NOT:

- calculate page ranges

- decide which blocks are rendered

- enforce ordering of child blocks

All pagination blocks MUST be provided explicitly

via the Content field and rendered in the exact

order defined by the caller.

This design matches Bulma documentation, where

`.pagination-previous`, `.pagination-next`, and

`.pagination-list` may appear in any order.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Content` | `elements.Items` | Content defines ordered pagination blocks.  Supported blocks include: - PaginationPrev - PaginationNext - PaginationList  Order is significant and MUST be defined explicitly. |
| `Size` | `string` | Size defines the pagination size modifier (is-small, is-medium, is-large). |
| `Alignment` | `string` | Alignment defines pagination alignment modifier (is-centered, is-right). |
| `Rounded` | `bool` | Rounded applies the `is-rounded` modifier. |
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the root `.pagination` container. |

## Usage

```go
@components.Pagination(
    components.PaginationProps{},
)
```

## Bulma Reference

[Bulma Pagination Documentation](https://bulma.io/documentation/components/pagination/)

## Source

**File:** `components/pagination.templ`
