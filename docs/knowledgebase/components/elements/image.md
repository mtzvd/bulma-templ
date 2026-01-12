---
title: "Image"
description: "Bulma image component."
category: "Elements"
atomicLevel: "ATOM"
bulmaClass: ".image"
---

# Image

**Atomic Level:** `ATOM`

Bulma image component.

## Description

Image renders a Bulma `.image` wrapper with a nested `<img>` element.

It supports fixed sizes, aspect ratios, and rounded images.

This component is layout-oriented and always renders a <figure>

element as required by Bulma.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Src` | `string` | Src defines the image source URL. This field is required. |
| `Alt` | `string` | Alt defines the alternative text for accessibility. This field is required. |
| `Ratio` | `string` | Ratio defines the aspect ratio modifier (is-square, is-16by9, is-4by3, etc.). |
| `Size` | `string` | Size defines the fixed image size (is-16x16, is-32x32, is-64x64, etc.). |
| `Rounded` | `bool` | Rounded applies the `is-rounded` modifier to the <img> element. |
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the <img> element. |
| `FigureAttr` | `templ.Attributes` | FigureAttr contains additional HTML attributes for the <figure> wrapper. |

## Usage

```go
@elements.Image(
    elements.ImageProps{},
)
```

## Bulma Reference

[Bulma Image Documentation](https://bulma.io/documentation/elements/image/)

## Source

**File:** `elements/image.templ`
