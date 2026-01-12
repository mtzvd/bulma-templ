---
title: "Media"
description: "Bulma media object layout component."
category: "Layout"
atomicLevel: "ORGANISM"
bulmaClass: ".media"
---

# Media

**Atomic Level:** `ORGANISM`

Bulma media object layout component.

## Description

Media defines a structural layout pattern used to align

images, icons, or avatars next to content.

This component is purely structural and does not

manage content behavior or state.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.media` container. |

## Usage

```go
@layout.Media(
    layout.MediaProps{},
    layout.Items{layout.Html("...")},
)
```

## Bulma Reference

[Bulma Media Documentation](https://bulma.io/documentation/layout/media/)

## Source

**File:** `layout/media.templ`
