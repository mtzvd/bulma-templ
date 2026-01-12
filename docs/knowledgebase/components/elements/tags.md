---
title: "Tags"
description: "Bulma tags container."
category: "Elements"
atomicLevel: "MOLECULE"
bulmaClass: ".tags"
---

# Tags

**Atomic Level:** `MOLECULE`

Bulma tags container.

## Description

Tags groups multiple Tag components together

using the Bulma `.tags` wrapper.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `HasAddons` | `bool` | HasAddons applies the `has-addons` modifier. |
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.tags` container. |

## Usage

```go
@elements.Tags(
    elements.TagsProps{},
    elements.Items{elements.Html("...")},
)
```

## Bulma Reference

[Bulma Tags Documentation](https://bulma.io/documentation/elements/tags/)

## Source

**File:** `elements/tag.templ`
