---
title: "Message"
description: ""
category: "Components"
atomicLevel: "ORGANISM"
bulmaClass: ".message"
---

# Message

**Atomic Level:** `ORGANISM`

## Description

Message is used to display informational,

success, warning, or error messages.

It does not manage state or closing behavior.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Color` | `string` | Color defines the Bulma color modifier (is-info, is-success, is-warning, is-danger). |
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.message` container. |

## Usage

```go
@components.Message(
    components.MessageProps{},
    components.Items{components.Html("...")},
)
```

## Bulma Reference

[Bulma Message Documentation](https://bulma.io/documentation/components/message/)

## Source

**File:** `components/message.templ`
