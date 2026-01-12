---
title: "Tag"
description: "Bulma tag element."
category: "Elements"
atomicLevel: "ATOM"
bulmaClass: ".tag"
---

# Tag

**Atomic Level:** `ATOM`

Bulma tag element.

## Description

Tag defines a content-based component that renders

a Bulma `.tag` element. It does not manage delete

buttons, icons, or links internally.

Any additional elements must be composed explicitly.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Color` | `string` | Color defines the tag color modifier (is-primary, is-info, is-success, is-warning, is-danger). |
| `Size` | `string` | Size defines the tag size modifier (is-small, is-medium, is-large). |
| `Rounded` | `bool` | Rounded applies the `is-rounded` modifier. |
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.tag` element. |

## Usage

```go
@elements.Tag(
    elements.TagProps{},
    elements.Items{elements.Html("...")},
)
```

## Bulma Reference

[Bulma Tag Documentation](https://bulma.io/documentation/elements/tag/)

## Source

**File:** `elements/tag.templ`
