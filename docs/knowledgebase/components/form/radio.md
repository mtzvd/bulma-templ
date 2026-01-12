---
title: "Radio"
description: "Bulma radio form control."
category: "Form"
atomicLevel: "ATOM"
bulmaClass: ".radio"
---

# Radio

**Atomic Level:** `ATOM`

Bulma radio form control.

## Description

Radio renders a Bulma `.radio` label together with

the native `<input type="radio">` element inside.

IMPORTANT:

This component intentionally exposes two attribute sets:

- WrapperAttr applies to the `<label class="radio">` element

- Attr applies to the native `<input>` element

Use WrapperAttr for layout, spacing and JS hooks.

Use Attr for form-related attributes (name, value, checked, etc.).

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Checked` | `bool` | Checked controls the checked state of the radio button. |
| `Disabled` | `bool` | Disabled disables the native radio input. |
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the native <input type="radio"> element.  Use this for: - name (REQUIRED for grouping) - value - data-* - x-model - aria-* |
| `WrapperAttr` | `templ.Attributes` | WrapperAttr contains additional HTML attributes for the <label class="radio"> wrapper.  Use this for: - layout-related classes - spacing helpers - Alpine.js directives |

## Usage

```go
@form.Radio(
    form.RadioProps{},
    form.Items{form.Html("...")},
)
```

## Bulma Reference

[Bulma Radio Documentation](https://bulma.io/documentation/form/radio/)

## Source

**File:** `form/radio.templ`
