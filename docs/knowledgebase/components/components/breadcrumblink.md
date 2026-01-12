---
title: "BreadcrumbLink"
description: "breadcrumb link element."
category: "Components"
atomicLevel: "ATOM"
bulmaClass: ".breadcrumb-link"
---

# BreadcrumbLink

**Atomic Level:** `ATOM`

breadcrumb link element.

## Description

BreadcrumbLink renders an `<a>` element inside BreadcrumbItem.

This component does not assume navigation semantics.

If `Href` is empty, the link is rendered without `href`.

Accessibility attributes (aria-current) MUST be provided

explicitly via props.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Href` | `string` | Href defines the link URL. When empty, no `href` attribute is rendered. |
| `Current` | `bool` | Current applies `aria-current="page"`. |
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `<a>` element. |

## Usage

```go
@components.BreadcrumbLink(
    components.BreadcrumbLinkProps{},
    components.Items{components.Html("...")},
)
```

## Bulma Reference

[Bulma BreadcrumbLink Documentation](https://bulma.io/documentation/components/breadcrumb-link/)

## Source

**File:** `components/breadcrumb.templ`
