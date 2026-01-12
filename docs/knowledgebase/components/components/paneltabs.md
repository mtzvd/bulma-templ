---
title: "PanelTabs"
description: ""
category: "Components"
atomicLevel: "MOLECULE"
bulmaClass: ".panel-tabs"
---

# PanelTabs

**Atomic Level:** `MOLECULE`

## Description

PanelTabs is used to display tab navigation inside Panel.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.panel-tabs` container. |

## Usage

```go
@components.PanelTabs(
    components.PanelTabsProps{},
    components.Items{components.Html("...")},
)
```

## Bulma Reference

[Bulma PanelTabs Documentation](https://bulma.io/documentation/components/panel-tabs/)

## Source

**File:** `components/panel.templ`
