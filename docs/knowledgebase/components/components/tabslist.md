---
title: "TabsList"
description: "tabs list container."
category: "Components"
atomicLevel: "MOLECULE"
bulmaClass: ".tabs-list"
---

# TabsList

**Atomic Level:** `MOLECULE`

tabs list container.

## Description

TabsList defines the structural <ul> wrapper inside Tabs.

It does NOT manage active state or navigation behavior.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the <ul> element. |

## Usage

```go
@components.TabsList(
    components.TabsListProps{},
    components.Items{components.Html("...")},
)
```

## Bulma Reference

[Bulma TabsList Documentation](https://bulma.io/documentation/components/tabs-list/)

## Source

**File:** `components/tabs.templ`
