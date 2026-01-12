---
title: "DropdownTrigger"
description: ""
category: "Components"
atomicLevel: "MOLECULE"
bulmaClass: ".dropdown-trigger"
---

# DropdownTrigger

**Atomic Level:** `MOLECULE`

## Description

DropdownTrigger is used to render the trigger element

(button, link, etc.) that controls the dropdown state.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.dropdown-trigger` container. |

## Usage

```go
@components.DropdownTrigger(
    components.DropdownTriggerProps{},
    components.Items{components.Html("...")},
)
```

## Bulma Reference

[Bulma DropdownTrigger Documentation](https://bulma.io/documentation/components/dropdown-trigger/)

## Source

**File:** `components/dropdown.templ`
