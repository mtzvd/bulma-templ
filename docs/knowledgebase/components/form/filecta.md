---
title: "FileCTA"
description: "Bulma file call-to-action."
category: "Form"
atomicLevel: "ATOM"
bulmaClass: ".file-c-t-a"
---

# FileCTA

**Atomic Level:** `ATOM`

Bulma file call-to-action.

## Description

FileCTA renders the `.file-cta` element

containing icon and label.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.file-cta` element. |

## Usage

```go
@form.FileCTA(
    form.FileCTAProps{},
    form.Items{form.Html("...")},
)
```

## Bulma Reference

[Bulma FileCTA Documentation](https://bulma.io/documentation/form/file-c-t-a/)

## Source

**File:** `form/file.templ`
