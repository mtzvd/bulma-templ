---
title: "TabsItem"
description: ""
category: "Components"
atomicLevel: "MOLECULE"
bulmaClass: ".tabs-item"
---

# TabsItem

**Atomic Level:** `MOLECULE`

## Description

Active state is applied on the <li> element

according to Bulma conventions.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Active` | `bool` | Active applies the `is-active` modifier. |
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the <li> element. |

## Usage

```go
@components.TabsItem(
    components.TabsItemProps{},
    components.Items{components.Html("...")},
)
```

## Bulma Reference

[Bulma TabsItem Documentation](https://bulma.io/documentation/components/tabs-item/)

## Source

**File:** `components/tabs.templ`
