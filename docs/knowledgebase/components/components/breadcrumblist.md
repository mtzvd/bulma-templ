---
title: "BreadcrumbList"
description: "breadcrumb list container."
category: "Components"
atomicLevel: "MOLECULE"
bulmaClass: ".breadcrumb-list"
---

# BreadcrumbList

**Atomic Level:** `MOLECULE`

breadcrumb list container.

## Description

BreadcrumbList renders the `<ul>` element inside Breadcrumb.

This component does not validate children and allows

arbitrary ordering and composition of breadcrumb items.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `<ul>` element. |

## Usage

```go
@components.BreadcrumbList(
    components.BreadcrumbListProps{},
    components.Items{components.Html("...")},
)
```

## Bulma Reference

[Bulma BreadcrumbList Documentation](https://bulma.io/documentation/components/breadcrumb-list/)

## Source

**File:** `components/breadcrumb.templ`
