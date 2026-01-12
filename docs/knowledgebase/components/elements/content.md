---
title: "Content"
description: "Bulma typographic container."
category: "Elements"
atomicLevel: "ATOM"
bulmaClass: ".content"
---

# Content

**Atomic Level:** `ATOM`

Bulma typographic container.

## Description

Content applies Bulma typography styles to all nested

HTML elements (headings, lists, tables, code blocks, etc.).

This component is used for articles, documentation,

FAQ sections, and SEO-oriented content.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Size` | `string` | Size defines the typography size modifier (is-small, is-medium, is-large). |
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.content` container. |

## Usage

```go
@elements.Content(
    elements.ContentProps{},
    elements.Items{elements.Html("...")},
)
```

## Bulma Reference

[Bulma Content Documentation](https://bulma.io/documentation/elements/content/)

## Source

**File:** `elements/content.templ`
