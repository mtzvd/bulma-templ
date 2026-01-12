---
title: "ModalBackground"
description: ""
category: "Components"
atomicLevel: "MOLECULE"
bulmaClass: ".modal-background"
---

# ModalBackground

**Atomic Level:** `MOLECULE`

## Description

ModalBackground is typically used to darken the background

and capture click events to close the modal.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.modal-background` element. |

## Usage

```go
@components.ModalBackground(
    components.ModalBackgroundProps{},
    components.Items{components.Html("...")},
)
```

## Bulma Reference

[Bulma ModalBackground Documentation](https://bulma.io/documentation/components/modal-background/)

## Source

**File:** `components/modal.templ`
