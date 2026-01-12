---
title: "PaginationLink"
description: "page link."
category: "Components"
atomicLevel: "ATOM"
bulmaClass: ".pagination-link"
---

# PaginationLink

**Atomic Level:** `ATOM`

page link.

## Description


PaginationLink — page link.
Atomic level: ATOM

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Active` | `bool` | Active indicates whether the link is active. |
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the link element. |

## Usage

```go
@components.PaginationLink(
    components.PaginationLinkProps{},
    components.Items{components.Html("...")},
)
```

## Bulma Reference

[Bulma PaginationLink Documentation](https://bulma.io/documentation/components/pagination-link/)

## Source

**File:** `components/pagination.templ`
