---
title: "MediaContent"
description: "main content section of Media."
category: "Layout"
atomicLevel: "MOLECULE"
bulmaClass: ".media-content"
---

# MediaContent

**Atomic Level:** `MOLECULE`

main content section of Media.

## Description

MediaContent holds the primary textual or structured

content associated with the media object.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.media-content` container. |

## Usage

```go
@layout.MediaContent(
    layout.MediaContentProps{},
    layout.Items{layout.Html("...")},
)
```

## Bulma Reference

[Bulma MediaContent Documentation](https://bulma.io/documentation/layout/media-content/)

## Source

**File:** `layout/media.templ`
