---
title: "BreadcrumbItem"
description: "breadcrumb list item."
category: "Components"
atomicLevel: "MOLECULE"
bulmaClass: ".breadcrumb-item"
---

# BreadcrumbItem

**Atomic Level:** `MOLECULE`

breadcrumb list item.

## Description

BreadcrumbItem renders a `<li>` element.

Active state MUST be provided explicitly by the caller.

This component does not infer current page.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Active` | `bool` | Active applies the `is-active` modifier. |
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `<li>` element. |

## Usage

```go
@components.BreadcrumbItem(
    components.BreadcrumbItemProps{},
    components.Items{components.Html("...")},
)
```

## Bulma Reference

[Bulma BreadcrumbItem Documentation](https://bulma.io/documentation/components/breadcrumb-item/)

## Source

**File:** `components/breadcrumb.templ`
