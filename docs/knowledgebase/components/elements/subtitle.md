---
title: "Subtitle"
description: "Bulma typographic subtitle."
category: "Elements"
atomicLevel: "ATOM"
bulmaClass: ".subtitle"
---

# Subtitle

**Atomic Level:** `ATOM`

Bulma typographic subtitle.

## Description

Subtitle applies Bulma `.subtitle` styles and does NOT

define HTML semantics.

Semantic HTML tags must be selected at the page

or layout level.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Size` | `string` | Size defines the visual size modifier (is-1, is-2, is-3, is-4, is-5, is-6). |
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the wrapper element. |

## Usage

```go
@elements.Subtitle(
    elements.SubtitleProps{},
    elements.Items{elements.Html("...")},
)
```

## Bulma Reference

[Bulma Subtitle Documentation](https://bulma.io/documentation/elements/subtitle/)

## Source

**File:** `elements/title.templ`
