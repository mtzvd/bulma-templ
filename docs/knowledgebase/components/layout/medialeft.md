---
title: "MediaLeft"
description: "left section of Media."
category: "Layout"
atomicLevel: "MOLECULE"
bulmaClass: ".media-left"
---

# MediaLeft

**Atomic Level:** `MOLECULE`

left section of Media.

## Description

MediaLeft is typically used for images, icons,

or avatars displayed to the left of the content.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.media-left` container. |

## Usage

```go
@layout.MediaLeft(
    layout.MediaLeftProps{},
    layout.Items{layout.Html("...")},
)
```

## Bulma Reference

[Bulma MediaLeft Documentation](https://bulma.io/documentation/layout/media-left/)

## Source

**File:** `layout/media.templ`
