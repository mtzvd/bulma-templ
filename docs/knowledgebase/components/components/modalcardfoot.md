---
title: "ModalCardFoot"
description: ""
category: "Components"
atomicLevel: "MOLECULE"
bulmaClass: ".modal-card-foot"
---

# ModalCardFoot

**Atomic Level:** `MOLECULE`

## Description

Typically contains action buttons.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.modal-card-foot` container. |

## Usage

```go
@components.ModalCardFoot(
    components.ModalCardFootProps{},
    components.Items{components.Html("...")},
)
```

## Bulma Reference

[Bulma ModalCardFoot Documentation](https://bulma.io/documentation/components/modal-card-foot/)

## Source

**File:** `components/modal.templ`
