---
title: "ModalContent"
description: ""
category: "Components"
atomicLevel: "MOLECULE"
bulmaClass: ".modal-content"
---

# ModalContent

**Atomic Level:** `MOLECULE`

## Description

ModalContent is used in the simple (non-card) modal variant.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.modal-content` container. |

## Usage

```go
@components.ModalContent(
    components.ModalContentProps{},
    components.Items{components.Html("...")},
)
```

## Bulma Reference

[Bulma ModalContent Documentation](https://bulma.io/documentation/components/modal-content/)

## Source

**File:** `components/modal.templ`
