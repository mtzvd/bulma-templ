---
title: "Section"
description: "Bulma section layout wrapper."
category: "Layout"
atomicLevel: "ORGANISM"
bulmaClass: ".section"
---

# Section

**Atomic Level:** `ORGANISM`

Bulma section layout wrapper.

## Description

Section provides vertical spacing between page sections

using the Bulma `.section` class.

This component does NOT constrain content width.

To limit horizontal width, compose it with Container explicitly.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Size` | `string` | Size defines the section size modifier (is-medium, is-large). |
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.section` element. |

## Usage

```go
@layout.Section(
    layout.SectionProps{},
    layout.Items{layout.Html("...")},
)
```

## Bulma Reference

[Bulma Section Documentation](https://bulma.io/documentation/layout/section/)

## Source

**File:** `layout/section.templ`
