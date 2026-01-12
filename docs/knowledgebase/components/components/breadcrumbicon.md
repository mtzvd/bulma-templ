---
title: "BreadcrumbIcon"
description: "icon wrapper for breadcrumb items."
category: "Components"
atomicLevel: "ATOM"
bulmaClass: ".breadcrumb-icon"
---

# BreadcrumbIcon

**Atomic Level:** `ATOM`

icon wrapper for breadcrumb items.

## Description

BreadcrumbIcon renders an icon inside a breadcrumb link.

This component is a thin wrapper around the Icon element

and exists solely for semantic clarity.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Icon` | `elements.IconProps` |  |

## Usage

```go
@components.BreadcrumbIcon(
    components.BreadcrumbIconProps{},
    components.Items{components.Html("...")},
)
```

## Bulma Reference

[Bulma BreadcrumbIcon Documentation](https://bulma.io/documentation/components/breadcrumb-icon/)

## Source

**File:** `components/breadcrumb.templ`
