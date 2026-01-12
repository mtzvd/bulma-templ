---
title: "TabsLink"
description: ""
category: "Components"
atomicLevel: "ATOM"
bulmaClass: ".tabs-link"
---

# TabsLink

**Atomic Level:** `ATOM`

## Description

TabsLink is used inside TabsItem to provide

correct UX (cursor, hover, focus) as per Bulma docs.

It does not manage tab state.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Href` | `string` | Href defines the tab link URL. If empty, TabsLink is rendered as a non-navigational interactive element. |
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the <a> element. |

## Usage

```go
@components.TabsLink(
    components.TabsLinkProps{},
    components.Items{components.Html("...")},
)
```

## Bulma Reference

[Bulma TabsLink Documentation](https://bulma.io/documentation/components/tabs-link/)

## Source

**File:** `components/tabs.templ`
