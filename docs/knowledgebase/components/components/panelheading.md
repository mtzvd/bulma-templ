---
title: "PanelHeading"
description: ""
category: "Components"
atomicLevel: "MOLECULE"
bulmaClass: ".panel-heading"
---

# PanelHeading

**Atomic Level:** `MOLECULE`

## Description

PanelHeading is used to display the panel title.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.panel-heading` element. |

## Usage

```go
@components.PanelHeading(
    components.PanelHeadingProps{},
    components.Items{components.Html("...")},
)
```

## Bulma Reference

[Bulma PanelHeading Documentation](https://bulma.io/documentation/components/panel-heading/)

## Source

**File:** `components/panel.templ`
