---
title: "ModalClose"
description: ""
category: "Components"
atomicLevel: "MOLECULE"
bulmaClass: ".modal-close"
---

# ModalClose

**Atomic Level:** `MOLECULE`

## Description

ModalClose is used in the simple modal variant.

Closing behavior must be implemented externally.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Size` | `string` | Size defines the close button size (is-small, is-medium, is-large). |
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the close button. |

## Usage

```go
@components.ModalClose(
    components.ModalCloseProps{},
)
```

## Bulma Reference

[Bulma ModalClose Documentation](https://bulma.io/documentation/components/modal-close/)

## Source

**File:** `components/modal.templ`
