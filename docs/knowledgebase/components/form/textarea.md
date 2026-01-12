---
title: "Textarea"
description: "Bulma textarea form element."
category: "Form"
atomicLevel: "ATOM"
bulmaClass: ".textarea"
---

# Textarea

**Atomic Level:** `ATOM`

Bulma textarea form element.

## Description

Textarea renders a Bulma `.textarea` form control.

It is responsible only for visual appearance and

HTML attributes, and does not manage layout,

validation logic or state.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Value` | `string` | Value defines the textarea content. |
| `Placeholder` | `string` | Placeholder defines the textarea placeholder text. |
| `Rows` | `int` | Rows defines the visible number of text lines. |
| `Color` | `string` | Color defines the textarea color modifier (e.g. is-primary, is-danger). |
| `Size` | `string` | Size defines the textarea size modifier (is-small, is-medium, is-large). |
| `Disabled` | `bool` | Disabled disables the textarea. |
| `ReadOnly` | `bool` | ReadOnly makes the textarea read-only. |
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the <textarea> element. |

## Usage

```go
@form.Textarea(
    form.TextareaProps{},
    form.Items{form.Html("...")},
)
```

## Bulma Reference

[Bulma Textarea Documentation](https://bulma.io/documentation/form/textarea/)

## Source

**File:** `form/textarea.templ`
