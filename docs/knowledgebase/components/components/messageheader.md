---
title: "MessageHeader"
description: ""
category: "Components"
atomicLevel: "MOLECULE"
bulmaClass: ".message-header"
---

# MessageHeader

**Atomic Level:** `MOLECULE`

## Description

MessageHeader is used to display the message title

and optional controls (e.g. delete button).

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.message-header` container. |

## Usage

```go
@components.MessageHeader(
    components.MessageHeaderProps{},
    components.Items{components.Html("...")},
)
```

## Bulma Reference

[Bulma MessageHeader Documentation](https://bulma.io/documentation/components/message-header/)

## Source

**File:** `components/message.templ`
