---
title: "Select"
description: "Bulma select form component."
category: "Form"
atomicLevel: "MOLECULE"
bulmaClass: ".select"
---

# Select

**Atomic Level:** `MOLECULE`

Bulma select form component.

## Description

Select renders a Bulma `.select` wrapper together with

the native `<select>` element inside.

IMPORTANT:

This component intentionally exposes two separate

attribute sets:

- WrapperAttr applies to the `.select` wrapper

- Attr applies to the native `<select>` element

This distinction is required because Bulma styles

and JavaScript hooks (e.g. Alpine.js) are often applied

to the wrapper, while form-related attributes

(name, value, data-*, x-model, etc.) must be applied

to the native control.

Options must be provided via the component content.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Color` | `string` | Color defines the select color modifier (e.g. is-primary, is-danger). |
| `Size` | `string` | Size defines the select size modifier (is-small, is-medium, is-large). |
| `Multiple` | `bool` | Multiple enables multiple selection.  Note: This affects both the Bulma wrapper and the native <select> element. |
| `FullWidth` | `bool` | FullWidth adds the is-fullwidth modifier to the Bulma wrapper. |
| `Disabled` | `bool` | Disabled disables the native <select> element. |
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the native <select> element.  Use this for: - name - value - data-* - x-model - aria-* |
| `WrapperAttr` | `templ.Attributes` | WrapperAttr contains additional HTML attributes for the Bulma `.select` wrapper.  Use this for: - layout-related classes - visibility helpers - Alpine.js state (x-data, x-show, etc.) - non-form-related data attributes |

## Usage

```go
@form.Select(
    form.SelectProps{},
    form.Items{form.Html("...")},
)
```

## Bulma Reference

[Bulma Select Documentation](https://bulma.io/documentation/form/select/)

## Source

**File:** `form/select.templ`
