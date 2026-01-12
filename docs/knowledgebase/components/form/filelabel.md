---
title: "FileLabel"
description: "Bulma file label wrapper."
category: "Form"
atomicLevel: "MOLECULE"
bulmaClass: ".file-label"
---

# FileLabel

**Atomic Level:** `MOLECULE`

Bulma file label wrapper.

## Description

FileLabel wraps the native file input

and all visual file UI elements.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.file-label` element. |

## Usage

```go
@form.FileLabel(
    form.FileLabelProps{},
    form.Items{form.Html("...")},
)
```

## Bulma Reference

[Bulma FileLabel Documentation](https://bulma.io/documentation/form/file-label/)

## Source

**File:** `form/file.templ`
