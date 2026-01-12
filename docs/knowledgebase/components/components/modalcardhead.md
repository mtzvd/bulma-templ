---
title: "ModalCardHead"
description: ""
category: "Components"
atomicLevel: "MOLECULE"
bulmaClass: ".modal-card-head"
---

# ModalCardHead

**Atomic Level:** `MOLECULE`

## Description

Typically contains the title and close button.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.modal-card-head` container. |

## Usage

```go
@components.ModalCardHead(
    components.ModalCardHeadProps{},
    components.Items{components.Html("...")},
)
```

## Bulma Reference

[Bulma ModalCardHead Documentation](https://bulma.io/documentation/components/modal-card-head/)

## Source

**File:** `components/modal.templ`
