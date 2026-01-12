---
title: "Tabs"
description: ""
category: "Components"
atomicLevel: "ORGANISM"
bulmaClass: ".tabs"
---

# Tabs

**Atomic Level:** `ORGANISM`

## Description

Tabs is used for navigation between sections.

It defines structural component and does not manage active state.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Centered` | `bool` | Centered applies the `is-centered` modifier. |
| `Right` | `bool` | Right applies the `is-right` modifier. |
| `Boxed` | `bool` | Boxed applies the `is-boxed` modifier. |
| `Toggle` | `bool` | Toggle applies the `is-toggle` modifier. |
| `ToggleRounded` | `bool` | ToggleRounded applies the `is-toggle-rounded` modifier. |
| `Fullwidth` | `bool` | Fullwidth applies the `is-fullwidth` modifier. |
| `Size` | `string` | Size defines the tabs size (is-small, is-medium, is-large). |
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.tabs` container. |

## Usage

```go
@components.Tabs(
    components.TabsProps{},
    components.Items{components.Html("...")},
)
```

## Bulma Reference

[Bulma Tabs Documentation](https://bulma.io/documentation/components/tabs/)

## Source

**File:** `components/tabs.templ`
