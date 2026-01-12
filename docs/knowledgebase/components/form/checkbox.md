---
title: "Checkbox"
description: "Bulma checkbox form control."
category: "Form"
atomicLevel: "ATOM"
bulmaClass: ".checkbox"
---

# Checkbox

**Atomic Level:** `ATOM`

Bulma checkbox form control.

## Description

Checkbox renders a Bulma `.checkbox` label together with

the native `<input type="checkbox">` element inside.

IMPORTANT:

This component intentionally exposes two attribute sets:

- WrapperAttr applies to the `<label class="checkbox">` element

- Attr applies to the native `<input>` element

Use WrapperAttr for layout, spacing and JS hooks.

Use Attr for form-related attributes (name, value, checked, etc.).

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Checked` | `bool` | Checked controls the checked state of the checkbox. |
| `Disabled` | `bool` | Disabled disables the native checkbox input. |
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the native <input type="checkbox"> element.  Use this for: - name - value - data-* - x-model - aria-* |
| `WrapperAttr` | `templ.Attributes` | WrapperAttr contains additional HTML attributes for the <label class="checkbox"> wrapper.  Use this for: - layout-related classes - spacing helpers - Alpine.js directives |

## Usage

```go
@form.Checkbox(
    form.CheckboxProps{},
    form.Items{form.Html("...")},
)
```

## Bulma Reference

[Bulma Checkbox Documentation](https://bulma.io/documentation/form/checkbox/)

## Source

**File:** `form/checkbox.templ`
