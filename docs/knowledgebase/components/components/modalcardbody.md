---
title: "ModalCardBody"
description: ""
category: "Components"
atomicLevel: "MOLECULE"
bulmaClass: ".modal-card-body"
---

# ModalCardBody

**Atomic Level:** `MOLECULE`

## Description

Used for the main content of a card-style modal.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.modal-card-body` container. |

## Usage

```go
@components.ModalCardBody(
    components.ModalCardBodyProps{},
    components.Items{components.Html("...")},
)
```

## Bulma Reference

[Bulma ModalCardBody Documentation](https://bulma.io/documentation/components/modal-card-body/)

## Source

**File:** `components/modal.templ`
