---
title: "Button"
description: "Bulma button component."
category: "Elements"
atomicLevel: "ATOM"
bulmaClass: ".button"
---

# Button

**Atomic Level:** `ATOM`

Bulma button component.

## Description

Button renders a Bulma `.button` element.

The component renders either an `<a>` or a `<button>`

depending on whether Href is provided.

This component is purely presentational and does NOT:

- manage state

- infer behavior

- render icons implicitly

Icons MUST be composed explicitly via the Icon component.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Href` | `string` | Href defines the link URL. When empty, a native <button> element is rendered. |
| `Type` | `string` | Type defines the button type attribute. Applicable only when rendering <button>. (button \| submit \| reset) |
| `Color` | `string` | Color defines the Bulma color modifier. |
| `Size` | `string` | Size defines the Bulma size modifier. |
| `State` | `string` | State defines the Bulma state modifier (is-hovered, is-focused, is-active, is-loading). |
| `FullWidth` | `bool` | FullWidth applies the `is-fullwidth` modifier. |
| `Disabled` | `bool` | Disabled applies the disabled state.  For <a> elements, this is purely visual. For <button> elements, the native `disabled` attribute is applied. |
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes. |

## Usage

```go
@elements.Button(
    elements.ButtonProps{},
    elements.Items{elements.Html("...")},
)
```

## Bulma Reference

[Bulma Button Documentation](https://bulma.io/documentation/elements/button/)

## Source

**File:** `elements/button.templ`
