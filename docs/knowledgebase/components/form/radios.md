---
title: "Radios"
description: "Bulma radios container."
category: "Form"
atomicLevel: "MOLECULE"
bulmaClass: ".radios"
---

# Radios

**Atomic Level:** `MOLECULE`

Bulma radios container.

## Description

Radios renders the Bulma `.radios` wrapper,

which groups multiple Radio components

into a single logical set.

This component provides structural grouping only

and does not manage radio state.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.radios` container. |

## Usage

```go
@form.Radios(
    form.RadiosProps{},
    form.Items{form.Html("...")},
)
```

## Bulma Reference

[Bulma Radios Documentation](https://bulma.io/documentation/form/radios/)

## Source

**File:** `form/radio.templ`
