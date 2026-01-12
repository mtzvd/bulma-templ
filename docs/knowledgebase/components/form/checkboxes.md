---
title: "Checkboxes"
description: "Bulma checkboxes container."
category: "Form"
atomicLevel: "ORGANISM"
bulmaClass: ".checkboxes"
---

# Checkboxes

**Atomic Level:** `ORGANISM`

Bulma checkboxes container.

## Description

Checkboxes renders the Bulma `.checkboxes` wrapper,

which groups multiple Checkbox components

into a single list.

This component provides structural grouping only

and does not manage checkbox state.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.checkboxes` container. |

## Usage

```go
@form.Checkboxes(
    form.CheckboxesProps{},
    form.Items{form.Html("...")},
)
```

## Bulma Reference

[Bulma Checkboxes Documentation](https://bulma.io/documentation/form/checkboxes/)

## Source

**File:** `form/checkbox.templ`
