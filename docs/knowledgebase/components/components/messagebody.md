---
title: "MessageBody"
description: ""
category: "Components"
atomicLevel: "MOLECULE"
bulmaClass: ".message-body"
---

# MessageBody

**Atomic Level:** `MOLECULE`

## Description

MessageBody is used to display the main message content.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.message-body` container. |

## Usage

```go
@components.MessageBody(
    components.MessageBodyProps{},
    components.Items{components.Html("...")},
)
```

## Bulma Reference

[Bulma MessageBody Documentation](https://bulma.io/documentation/components/message-body/)

## Source

**File:** `components/message.templ`
