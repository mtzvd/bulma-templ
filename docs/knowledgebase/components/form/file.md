---
title: "File"
description: "Bulma file upload component."
category: "Form"
atomicLevel: "MOLECULE"
bulmaClass: ".file"
---

# File

**Atomic Level:** `MOLECULE`

Bulma file upload component.

## Description

File renders the Bulma `.file` wrapper.

It provides visual styling and layout for

file input controls and related UI elements.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Color` | `string` | Color defines the file color modifier (e.g. is-primary, is-danger). |
| `Size` | `string` | Size defines the file size modifier (is-small, is-medium, is-large). |
| `FullWidth` | `bool` | FullWidth adds the is-fullwidth modifier. |
| `Boxed` | `bool` | Boxed adds the is-boxed modifier. |
| `HasName` | `bool` | HasName adds the has-name modifier. |
| `Align` | `string` | Align defines alignment modifier (is-centered, is-right). |
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.file` wrapper. |

## Usage

```go
@form.File(
    form.FileProps{},
    form.Items{form.Html("...")},
)
```

## Bulma Reference

[Bulma File Documentation](https://bulma.io/documentation/form/file/)

## Source

**File:** `form/file.templ`
