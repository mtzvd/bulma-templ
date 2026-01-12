---
title: "Breadcrumb"
description: "Bulma breadcrumb container."
category: "Components"
atomicLevel: "ORGANISM"
bulmaClass: ".breadcrumb"
---

# Breadcrumb

**Atomic Level:** `ORGANISM`

Bulma breadcrumb container.

## Description

Breadcrumb renders the root `.breadcrumb` navigation container.

This component is a pure structural container and does NOT:

- infer active items

- calculate navigation hierarchy

- manage routing or application state

- impose ordering of child elements

All structure, ordering, and active state MUST be

composed explicitly by the caller using subcomponents.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Separator` | `string` | Separator defines the Bulma separator modifier.  Examples: - has-arrow-separator - has-bullet-separator - has-dot-separator - has-succeeds-separator |
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `<nav>` element. |

## Usage

```go
@components.Breadcrumb(
    components.BreadcrumbProps{},
    components.Items{components.Html("...")},
)
```

## Bulma Reference

[Bulma Breadcrumb Documentation](https://bulma.io/documentation/components/breadcrumb/)

## Source

**File:** `components/breadcrumb.templ`
