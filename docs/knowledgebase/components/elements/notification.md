---
title: "Notification"
description: "Bulma notification component."
category: "Elements"
atomicLevel: "ATOM"
bulmaClass: ".notification"
---

# Notification

**Atomic Level:** `ATOM`

Bulma notification component.

## Description

Notification renders a Bulma `.notification` container.

It does not manage close buttons, state, or behavior.

Any additional elements (such as a delete button)

must be composed explicitly by the caller.

## Props

| Property | Type | Description |
|----------|------|-------------|
| `Color` | `string` | Color defines the Bulma color modifier (is-primary, is-info, is-success, is-warning, is-danger). |
| `Attr` | `templ.Attributes` | Attr contains additional HTML attributes for the `.notification` container. |

## Usage

```go
@elements.Notification(
    elements.NotificationProps{},
    elements.Items{elements.Html("...")},
)
```

## Bulma Reference

[Bulma Notification Documentation](https://bulma.io/documentation/elements/notification/)

## Source

**File:** `elements/notification.templ`
