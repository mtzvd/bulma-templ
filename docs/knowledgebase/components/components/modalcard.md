---
title: "ModalCard"
description: ""
category: "Components"
atomicLevel: "MOLECULE"
bulmaClass: ".modal-card"
---

# ModalCard

**Atomic Level:** `MOLECULE`

## Description

ModalCard is used for the card-style modal layout.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.modal-card` container. |

## Usage

```go
@components.ModalCard(
    components.ModalCardProps{},
    components.Items{components.Html("...")},
)
```

## Bulma Reference

[Bulma ModalCard Documentation](https://bulma.io/documentation/components/modal-card/)

## Source

**File:** `components/modal.templ`
