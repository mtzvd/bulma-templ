---
title: "MediaRight"
description: "right section of Media."
category: "Layout"
atomicLevel: "MOLECULE"
bulmaClass: ".media-right"
---

# MediaRight

**Atomic Level:** `MOLECULE`

right section of Media.

## Description

MediaRight is used for actions such as buttons,

icons, or delete controls aligned to the right.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.media-right` container. |

## Usage

```go
@layout.MediaRight(
    layout.MediaRightProps{},
    layout.Items{layout.Html("...")},
)
```

## Bulma Reference

[Bulma MediaRight Documentation](https://bulma.io/documentation/layout/media-right/)

## Source

**File:** `layout/media.templ`
