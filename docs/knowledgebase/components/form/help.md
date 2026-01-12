---
title: "Help"
description: "Bulma help text element."
category: "Form"
atomicLevel: "ATOM"
bulmaClass: ".help"
---

# Help

**Atomic Level:** `ATOM`

Bulma help text element.

## Description

Help defines a visual helper or feedback message

displayed below a form control.

It does NOT perform validation or state management.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Text` | `string` | Text defines the help message content. |
| `Color` | `string` | Color defines the message color (e.g. is-danger, is-success). |
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.help` element. |

## Usage

```go
@form.Help(
    form.HelpProps{},
)
```

## Bulma Reference

[Bulma Help Documentation](https://bulma.io/documentation/form/help/)

## Source

**File:** `form/help.templ`
