---
title: "Panel"
description: ""
category: "Components"
atomicLevel: "ORGANISM"
bulmaClass: ".panel"
---

# Panel

**Atomic Level:** `ORGANISM`

## Description

Panel is used to group related content and actions.

It defines structural component and does not manage state.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.panel` container. |

## Usage

```go
@components.Panel(
    components.PanelProps{},
    components.Items{components.Html("...")},
)
```

## Bulma Reference

[Bulma Panel Documentation](https://bulma.io/documentation/components/panel/)

## Source

**File:** `components/panel.templ`
