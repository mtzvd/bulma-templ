---
title: "Modal"
description: ""
category: "Components"
atomicLevel: "ORGANISM"
bulmaClass: ".modal"
---

# Modal

**Atomic Level:** `ORGANISM`

## Description

Modal defines the root container for modal dialogs.

It is responsible only for structure and CSS modifiers.

Open / close state must be managed externally.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Active` | `bool` | Active applies the `is-active` modifier. Useful for server-side controlled modals. |
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.modal` container. Commonly used for x-data, x-show, aria-*, data-*. |

## Usage

```go
@components.Modal(
    components.ModalProps{},
    components.Items{components.Html("...")},
)
```

## Bulma Reference

[Bulma Modal Documentation](https://bulma.io/documentation/components/modal/)

## Source

**File:** `components/modal.templ`
