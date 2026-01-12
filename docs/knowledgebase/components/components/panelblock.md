---
title: "PanelBlock"
description: ""
category: "Components"
atomicLevel: "MOLECULE"
bulmaClass: ".panel-block"
---

# PanelBlock

**Atomic Level:** `MOLECULE`

## Description

PanelBlock can be rendered as a link (<a>) or

as a container (<div>) depending on the presence of Href.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Href` | `string` | Href defines the link URL. If empty, the block is rendered as a <div>. |
| `Active` | `bool` | Active applies the `is-active` modifier. |
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.panel-block` element. |

## Usage

```go
@components.PanelBlock(
    components.PanelBlockProps{},
    components.Items{components.Html("...")},
)
```

## Bulma Reference

[Bulma PanelBlock Documentation](https://bulma.io/documentation/components/panel-block/)

## Source

**File:** `components/panel.templ`
