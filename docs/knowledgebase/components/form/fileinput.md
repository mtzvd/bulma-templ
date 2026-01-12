---
title: "FileInput"
description: "native file input element."
category: "Form"
atomicLevel: "MOLECULE"
bulmaClass: ".file-input"
---

# FileInput

**Atomic Level:** `MOLECULE`

native file input element.

## Description

FileInput renders the `<input type=\"file\">` element.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Disabled` | `bool` | Disabled disables the file input. |
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the <input type="file"> element.  Use this for: - name - accept - multiple - data-* - x-on:change |

## Usage

```go
@form.FileInput(
    form.FileInputProps{},
    form.Items{form.Html("...")},
)
```

## Bulma Reference

[Bulma FileInput Documentation](https://bulma.io/documentation/form/file-input/)

## Source

**File:** `form/file.templ`
