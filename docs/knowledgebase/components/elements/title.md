---
title: "Title"
description: "Bulma typographic title."
category: "Elements"
atomicLevel: "ATOM"
bulmaClass: ".title"
---

# Title

**Atomic Level:** `ATOM`

Bulma typographic title.

## Description

Title defines a visual typographic component that applies

the Bulma `.title` styles.

It does NOT define HTML semantics or document structure.

The semantic HTML element must be chosen at the

page or layout level.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Size` | `string` | Size defines the visual size modifier (is-1, is-2, is-3, is-4, is-5, is-6). |
| `Spaced` | `bool` | Spaced applies the `is-spaced` modifier. |
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the wrapper element. |

## Usage

```go
@elements.Title(
    elements.TitleProps{},
    elements.Items{elements.Html("...")},
)
```

## Bulma Reference

[Bulma Title Documentation](https://bulma.io/documentation/elements/title/)

## Source

**File:** `elements/title.templ`
