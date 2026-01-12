---
title: "Delete"
description: "Bulma delete button."
category: "Elements"
atomicLevel: "ATOM"
bulmaClass: ".delete"
---

# Delete

**Atomic Level:** `ATOM`

Bulma delete button.

## Description

Delete renders a Bulma `.delete` element as a native <button>.

It does not contain text, state, or behavior logic.

Any interaction (Alpine.js, aria-*, data-*) must be provided

via Attr.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Size` | `string` | Size defines the delete button size (is-small, is-medium, is-large). |
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the <button> element. |

## Usage

```go
@elements.Delete(
    elements.DeleteProps{},
    elements.Items{elements.Html("...")},
)
```

## Bulma Reference

[Bulma Delete Documentation](https://bulma.io/documentation/elements/delete/)

## Source

**File:** `elements/delete.templ`
