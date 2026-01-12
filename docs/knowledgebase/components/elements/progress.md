---
title: "Progress"
description: "Bulma progress bar."
category: "Elements"
atomicLevel: "ATOM"
bulmaClass: ".progress"
---

# Progress

**Atomic Level:** `ATOM`

Bulma progress bar.

## Description

Progress defines a thin wrapper around the native <progress>

element styled with Bulma classes.

This component does NOT compute or format progress values.

Any content rendered inside <progress> is fully controlled

by the caller, exactly as allowed by Bulma.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Color` | `string` | Color defines the Bulma color modifier (is-primary, is-link, is-info, is-success, etc.). |
| `Size` | `string` | Size defines the Bulma size modifier (is-small, is-medium, is-large). |
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the <progress> element. |
| `Max` | `int` | Max is the maximum value of the progress bar (maps to the "max" attribute). Use ProgressUnset (-1) to omit this attribute. |
| `Current` | `int` | Current is the current progress value (maps to the "value" attribute). Use ProgressUnset (-1) to omit this attribute and show indeterminate progress. |

## Usage

```go
@elements.Progress(
    elements.ProgressProps{},
    elements.Items{elements.Html("...")},
)
```

## Bulma Reference

[Bulma Progress Documentation](https://bulma.io/documentation/elements/progress/)

## Source

**File:** `elements/progress.templ`
