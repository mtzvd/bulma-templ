---
title: "Control"
description: "Bulma form control wrapper."
category: "Form"
atomicLevel: "MOLECULE"
bulmaClass: ".control"
---

# Control

**Atomic Level:** `MOLECULE`

Bulma form control wrapper.

## Description

Control wraps interactive form elements such as

input, select, textarea or button.

It is responsible only for layout concerns like

icon positioning and width expansion.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Expanded` | `bool` | Expanded adds the is-expanded modifier. |
| `HasIconsLeft` | `bool` | HasIconsLeft adds the has-icons-left modifier. |
| `HasIconsRight` | `bool` | HasIconsRight adds the has-icons-right modifier. |
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.control` container. |

## Usage

```go
@form.Control(
    form.ControlProps{},
    form.Items{form.Html("...")},
)
```

## Bulma Reference

[Bulma Control Documentation](https://bulma.io/documentation/form/control/)

## Source

**File:** `form/control.templ`
