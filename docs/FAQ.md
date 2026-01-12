# FAQ — Bulma-Templ

Frequently Asked Questions about the Bulma-Templ library.

## Table of Contents

- [Getting Started](#getting-started)
- [Architecture & Design](#architecture--design)
- [Component Usage](#component-usage)
- [Common Patterns](#common-patterns)
- [Integration](#integration)
- [Troubleshooting](#troubleshooting)
- [Contributing](#contributing)

---

## Getting Started

### What is Bulma-Templ?

Bulma-Templ is a **type-safe, SSR-first** component library that provides 1:1 mappings of [Bulma CSS](https://bulma.io) components for Go using the [Templ](https://templ.guide) template engine.

**Key features:**
- Complete Bulma coverage (elements, components, forms, layouts)
- Type-safe Props structs with IDE autocomplete
- Zero internal state — pure rendering functions
- Explicit composition via `Items` pattern
- No JavaScript required for rendering

### Why use Bulma-Templ instead of writing HTML directly?

**Benefits:**

1. **Type safety** — Compile-time validation of component props
2. **IDE support** — Autocomplete for all component properties
3. **Consistency** — Enforces Bulma patterns and best practices
4. **Refactoring** — Change component structure without breaking HTML
5. **Composition** — Explicit, predictable component nesting
6. **Documentation** — Props are self-documenting

**Example comparison:**

```go
// ❌ Raw HTML (error-prone, no autocomplete)
<button class="button is-primary is-large">Click Me</button>

// ✅ Bulma-Templ (type-safe, autocomplete, refactorable)
@elements.Button(
    elements.ButtonProps{Color: "is-primary", Size: "is-large"},
    elements.Items{elements.Html("Click Me")},
)
```

### How do I install Bulma-Templ?

```bash
# Install the library
go get github.com/mtzvd/bulma-templ
```

That's it! Generated `*_templ.go` files are included in the repository for pkg.go.dev compatibility.

### What versions of Go and Bulma are supported?

- **Go:** 1.25+ (uses latest Go features)
- **Bulma CSS:** 1.0+ (library maps to Bulma 1.0 API)
- **Templ:** Latest version (install via `go install`)

### Do I need to include Bulma CSS separately?

**Yes.** Bulma-Templ generates HTML with Bulma CSS classes, but does NOT include the CSS itself.

**Add Bulma CSS via CDN:**

```html
<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bulma@1.0.0/css/bulma.min.css">
```

**Or install via npm/yarn:**

```bash
npm install bulma
```

Then include it in your build process.

---

## Architecture & Design

### Why the `Items` pattern instead of variadic arguments?

The `Items` pattern provides:

1. **Explicit composition** — Clear visual structure
2. **Type safety** — All children are `templ.Component`
3. **Predictable order** — No hidden sorting or re-ordering
4. **Zero magic** — What you write is what renders

**Comparison:**

```go
// ❌ Variadic (implicit, order unclear)
Card(CardHeader(...), CardContent(...), CardFooter(...))

// ✅ Items (explicit, structured)
@Card(CardProps{}, Items{
    CardHeader(...),
    CardContent(...),
    CardFooter(...),
})
```

### Why separate Props struct instead of individual parameters?

**Reasons:**

1. **Named parameters** — Go doesn't have named args, Props provides this
2. **Optional parameters** — Zero values work as sensible defaults
3. **Extensibility** — Easy to add new props without breaking API
4. **IDE autocomplete** — Discover all available options
5. **Documentation** — Props are self-documenting

**Example:**

```go
// ✅ Props struct (readable, extensible)
Button(ButtonProps{
    Color: "is-primary",
    Size: "is-large",
    Disabled: true,
}, Items{...})

// ❌ Individual params (unclear, breaks on changes)
Button("is-primary", "is-large", true, false, false, Items{...})
```

### Why 1:1 Bulma mapping instead of custom abstractions?

**Design philosophy:**

> Bulma-Templ is a **thin, transparent layer** over Bulma CSS. It provides type safety without re-inventing Bulma.

**Benefits:**

1. **Bulma docs are your docs** — No need to learn new abstractions
2. **Predictable HTML** — Output matches Bulma examples exactly
3. **Easy migration** — Convert HTML to components mechanically
4. **No surprises** — Components do exactly what Bulma docs say
5. **Stability** — No opinions means fewer breaking changes

**Example:**

```go
// ✅ Direct Bulma mapping
Button(ButtonProps{Color: "is-primary", Size: "is-large"}, Items{...})

// ❌ Custom abstraction (requires new docs, new learning)
PrimaryButton(size: Large, variant: Solid, ...)
```

### What is the "zero internal state" principle?

**Principle:** Components NEVER manage state internally. They are pure rendering functions.

**This means:**

- ❌ No component-level variables
- ❌ No event handlers in component code
- ❌ No lifecycle hooks
- ✅ All state is passed via Props
- ✅ Components render HTML only

**Example:**

```go
// ✅ Correct: State is external
templ MyPage(isOpen bool) {
    @Modal(ModalProps{IsActive: isOpen}, Items{...})
}

// ❌ Wrong: Component manages state (not possible in this library)
// Components don't have internal state variables
```

### Why SSR-first? Can I use this with client-side frameworks?

**SSR-first means:**

- Components render complete HTML on the server
- No hydration required
- Works with JavaScript disabled
- Perfect for HTMX, Alpine.js, or no JS at all

**Can you use it with React/Vue/Svelte?**

Not recommended. This library is designed for **server-side rendering**. If you need client-side reactivity, use Bulma's official framework integrations.

**Use cases:**
- ✅ Go web servers (net/http, Fiber, Echo, Chi)
- ✅ HTMX applications
- ✅ Alpine.js for interactivity
- ✅ Server-side rendered pages
- ❌ Single-page applications (SPA)
- ❌ React/Vue/Svelte (use framework-specific Bulma libs)

---

## Component Usage

### What's the difference between `Html()` and `templ.Raw()`?

**`elements.Html(string)`** — Renders a string as a `templ.Component`

```go
elements.Html("Hello World")  // Returns templ.Component
```

**`templ.Raw(string)`** — Renders unescaped HTML (DANGEROUS)

```go
templ.Raw("<strong>HTML</strong>")  // Renders raw HTML
```

**When to use:**

- **Use `Html()`** for most cases (safe, escaped text)
- **Use `templ.Raw()`** only when you MUST render unescaped HTML (trusted sources only!)

**Example:**

```go
// ✅ Safe text rendering
@Button(ButtonProps{}, Items{Html("Click Me")})

// ⚠️ Unescaped HTML (only for trusted content!)
@Content(ContentProps{}, Items{templ.Raw("<p>Trusted <strong>HTML</strong></p>")})
```

### How do I add custom CSS classes to components?

Use the `Attr` field in Props:

```go
@Button(
    ButtonProps{
        Color: "is-primary",
        Attr: templ.Attributes{
            "class": "my-custom-class",
        },
    },
    Items{Html("Custom Button")},
)
```

**Output:**

```html
<button class="button is-primary my-custom-class">Custom Button</button>
```

**Note:** Custom classes are **merged** with Bulma classes, not replaced.

### How do I add Alpine.js or HTMX attributes?

Use the `Attr` escape hatch:

```go
// Alpine.js example
@Button(
    ButtonProps{
        Color: "is-primary",
        Attr: templ.Attributes{
            "x-on:click": "count++",
            "x-text":     "count",
        },
    },
    Items{},
)

// HTMX example
@Button(
    ButtonProps{
        Attr: templ.Attributes{
            "hx-post":   "/api/submit",
            "hx-target": "#result",
        },
    },
    Items{Html("Submit")},
)
```

### How do I pass data attributes or ARIA attributes?

Same as above — use `Attr`:

```go
@Button(
    ButtonProps{
        Attr: templ.Attributes{
            "data-id":        "123",
            "data-analytics": "click-event",
            "aria-label":     "Close dialog",
            "role":           "button",
        },
    },
    Items{Html("Close")},
)
```

### How do I conditionally render components?

Use Templ's `if` statements:

```go
templ MyPage(showNotification bool) {
    if showNotification {
        @Notification(
            NotificationProps{Color: "is-success"},
            Items{Html("Success!")},
        )
    }
}
```

### How do I loop over components?

Use Templ's `for` loops:

```go
templ UserList(users []User) {
    @Box(BoxProps{}, Items{})
        for _, user := range users {
            @Tag(TagProps{}, Items{Html(user.Name)})
        }
    }
}
```

### Can I nest components?

**Yes!** Components are designed for nesting via `Items`:

```go
@Card(CardProps{}, Items{
    CardHeader(CardHeaderProps{}, Items{
        CardHeaderTitle(CardHeaderTitleProps{}, Items{
            Html("Card Title"),
        }),
    }),
    CardContent(CardContentProps{}, Items{
        Content(ContentProps{}, Items{
            Html("<p>Card body text</p>"),
        }),
    }),
    CardFooter(CardFooterProps{}, Items{
        CardFooterItem(CardFooterItemProps{}, Items{
            Html("Save"),
        }),
    }),
})
```

---

## Common Patterns

### How do I create a form with validation?

```go
templ LoginForm(emailError string) {
    @form.Field(form.FieldProps{}, form.Items{
        form.Label(form.LabelProps{}, form.Items{
            elements.Html("Email"),
        }),
        form.Control(form.ControlProps{}, form.Items{
            form.Input(form.InputProps{
                Type:        "email",
                Placeholder: "user@example.com",
                Name:        "email",
            }),
        }),
        if emailError != "" {
            form.Help(form.HelpProps{Color: "is-danger"}, form.Items{
                elements.Html(emailError),
            })
        },
    })
}
```

### How do I create a responsive navbar?

```go
templ Navbar(isActive bool) {
    @components.Navbar(components.NavbarProps{}, components.Items{
        components.NavbarBrand(components.NavbarBrandProps{}, components.Items{
            components.NavbarItem(components.NavbarItemProps{
                Href: "/",
            }, components.Items{
                elements.Html("My Site"),
            }),
            components.NavbarBurger(components.NavbarBurgerProps{
                Attr: templ.Attributes{
                    "x-on:click": "navbarOpen = !navbarOpen",
                },
            }, components.Items{}),
        }),
        components.NavbarMenu(components.NavbarMenuProps{
            IsActive: isActive,
        }, components.Items{
            components.NavbarStart(components.NavbarStartProps{}, components.Items{
                components.NavbarItem(components.NavbarItemProps{
                    Href: "/about",
                }, components.Items{
                    elements.Html("About"),
                }),
            }),
        }),
    })
}
```

### How do I create a modal with Alpine.js?

```go
templ MyPage() {
    <div x-data="{ modalOpen: false }">
        @elements.Button(
            elements.ButtonProps{
                Color: "is-primary",
                Attr: templ.Attributes{
                    "x-on:click": "modalOpen = true",
                },
            },
            elements.Items{elements.Html("Open Modal")},
        )

        @components.Modal(
            components.ModalProps{
                Attr: templ.Attributes{
                    "x-show": "modalOpen",
                },
            },
            components.Items{
                components.ModalBackground(components.ModalBackgroundProps{
                    Attr: templ.Attributes{
                        "x-on:click": "modalOpen = false",
                    },
                }, components.Items{}),
                components.ModalContent(components.ModalContentProps{}, components.Items{
                    elements.Box(elements.BoxProps{}, elements.Items{
                        elements.Html("<p>Modal content here</p>"),
                    }),
                }),
                components.ModalClose(components.ModalCloseProps{
                    Size: "is-large",
                    Attr: templ.Attributes{
                        "x-on:click": "modalOpen = false",
                    },
                }, components.Items{}),
            },
        )
    </div>
}
```

### How do I create a button with an icon?

```go
@elements.Button(
    elements.ButtonProps{Color: "is-primary"},
    elements.Items{
        elements.Icon(elements.IconProps{}, elements.Items{
            elements.Html(`<i class="fas fa-save"></i>`),
        }),
        elements.Html("Save"),
    },
)
```

**Or with icon-only button:**

```go
@elements.Button(
    elements.ButtonProps{},
    elements.Items{
        elements.Icon(elements.IconProps{}, elements.Items{
            elements.Html(`<i class="fas fa-trash"></i>`),
        }),
    },
)
```

### How do I create a multi-column layout?

```go
@columns.Columns(columns.ColumnsProps{}, columns.Items{
    columns.Column(columns.ColumnProps{Size: "is-half"}, columns.Items{
        elements.Box(elements.BoxProps{}, elements.Items{
            elements.Html("Left column"),
        }),
    }),
    columns.Column(columns.ColumnProps{Size: "is-half"}, columns.Items{
        elements.Box(elements.BoxProps{}, elements.Items{
            elements.Html("Right column"),
        }),
    }),
})
```

### How do I create tabs with content switching?

```go
templ TabsExample() {
    <div x-data="{ activeTab: 'profile' }">
        @components.Tabs(components.TabsProps{}, components.Items{
            components.TabsList(components.TabsListProps{}, components.Items{
                components.TabsItem(components.TabsItemProps{
                    Attr: templ.Attributes{
                        ":class": "{ 'is-active': activeTab === 'profile' }",
                    },
                }, components.Items{
                    components.TabsLink(components.TabsLinkProps{
                        Attr: templ.Attributes{
                            "x-on:click.prevent": "activeTab = 'profile'",
                            "href":               "#",
                        },
                    }, components.Items{
                        elements.Html("Profile"),
                    }),
                }),
                components.TabsItem(components.TabsItemProps{
                    Attr: templ.Attributes{
                        ":class": "{ 'is-active': activeTab === 'settings' }",
                    },
                }, components.Items{
                    components.TabsLink(components.TabsLinkProps{
                        Attr: templ.Attributes{
                            "x-on:click.prevent": "activeTab = 'settings'",
                            "href":               "#",
                        },
                    }, components.Items{
                        elements.Html("Settings"),
                    }),
                }),
            }),
        })

        <div x-show="activeTab === 'profile'">
            Profile content
        </div>
        <div x-show="activeTab === 'settings'">
            Settings content
        </div>
    </div>
}
```

---

## Integration

### How do I use Bulma-Templ with HTMX?

Bulma-Templ works perfectly with HTMX. Use `Attr` to add HTMX attributes:

```go
templ TodoItem(todo Todo) {
    @elements.Box(
        elements.BoxProps{
            Attr: templ.Attributes{
                "hx-delete":  fmt.Sprintf("/todos/%d", todo.ID),
                "hx-target":  "this",
                "hx-swap":    "outerHTML",
            },
        },
        elements.Items{
            elements.Html(todo.Title),
            elements.Button(
                elements.ButtonProps{
                    Color: "is-danger",
                    Size:  "is-small",
                    Attr: templ.Attributes{
                        "hx-trigger": "click",
                    },
                },
                elements.Items{elements.Html("Delete")},
            ),
        },
    )
}
```

### How do I use Bulma-Templ with Echo/Fiber/Chi?

**Echo example:**

```go
package main

import (
    "github.com/labstack/echo/v4"
    "github.com/mtzvd/bulma-templ/elements"
)

func main() {
    e := echo.New()

    e.GET("/", func(c echo.Context) error {
        return Page().Render(c.Request().Context(), c.Response().Writer)
    })

    e.Start(":8080")
}

templ Page() {
    <!DOCTYPE html>
    <html>
        <head>
            <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bulma@1.0.0/css/bulma.min.css">
        </head>
        <body>
            @elements.Button(
                elements.ButtonProps{Color: "is-primary"},
                elements.Items{elements.Html("Click Me")},
            )
        </body>
    </html>
}
```

**Fiber/Chi:** Similar pattern — render the templ component to `http.ResponseWriter`.

### Can I use this library with Tailwind CSS instead of Bulma?

**No.** This library is specifically designed for Bulma CSS. The components generate Bulma-specific class names.

If you want Tailwind, consider:
- Writing raw Templ templates with Tailwind classes
- Creating a similar library for Tailwind (different project)

### How do I add custom JavaScript to components?

**Don't.** Components should not include JavaScript.

**Instead:**

1. **Use Alpine.js** via `Attr` for simple interactivity
2. **Add `<script>` tags** in your page template
3. **Use HTMX** for server-driven interactions

**Example:**

```go
templ Page() {
    <!DOCTYPE html>
    <html>
        <head>
            <script src="https://unpkg.com/alpinejs@3.x.x/dist/cdn.min.js" defer></script>
        </head>
        <body>
            <div x-data="{ count: 0 }">
                @elements.Button(
                    elements.ButtonProps{
                        Attr: templ.Attributes{
                            "x-on:click": "count++",
                        },
                    },
                    elements.Items{elements.Html("Increment")},
                )
                <span x-text="count"></span>
            </div>
        </body>
    </html>
}
```

---

## Troubleshooting

### Error: `not enough arguments in call to Component`

**Cause:** Missing `Props` or `Items` parameter.

**Solution:** All components require both `Props` and `Items`:

```go
// ❌ Wrong
Button(Items{Html("Click")})

// ✅ Correct
Button(ButtonProps{}, Items{Html("Click")})
```

### Error: `cannot use string as Items value`

**Cause:** Trying to pass a string directly instead of wrapping in `Items`.

**Solution:** Use `elements.Html()` to convert strings:

```go
// ❌ Wrong
Button(ButtonProps{}, "Click Me")

// ✅ Correct
Button(ButtonProps{}, Items{Html("Click Me")})
```

### Error: `undefined: elements` or `undefined: components`

**Cause:** Missing import statement.

**Solution:** Import the package:

```go
import (
    "github.com/mtzvd/bulma-templ/elements"
    "github.com/mtzvd/bulma-templ/components"
    "github.com/mtzvd/bulma-templ/form"
)
```

### Error: `templ command not found`

**Cause:** Templ CLI is not installed.

**Solution:**

```bash
go install github.com/a-h/templ/cmd/templ@latest
```

Make sure `$GOPATH/bin` is in your `$PATH`.

### Generated `*_templ.go` files are missing (for contributors)

**Cause:** You're modifying `.templ` files and need to regenerate.

**Solution:**

```bash
templ generate
# or
task templ
```

**Note:** Generated files ARE tracked in git for pkg.go.dev compatibility. Always regenerate after modifying `.templ` files.

### Components are rendering but styles are missing

**Cause:** Bulma CSS is not included in your HTML.

**Solution:** Add Bulma CSS to your `<head>`:

```html
<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bulma@1.0.0/css/bulma.min.css">
```

### How do I debug the generated HTML?

**Option 1:** Use browser DevTools to inspect the rendered HTML.

**Option 2:** Print the component output:

```go
package main

import (
    "context"
    "os"
    "github.com/mtzvd/bulma-templ/elements"
)

func main() {
    component := elements.Button(
        elements.ButtonProps{Color: "is-primary"},
        elements.Items{elements.Html("Test")},
    )

    component.Render(context.Background(), os.Stdout)
}
```

### Why are my custom classes not appearing?

**Check:**

1. Are you using `Attr` correctly?
2. Are you inspecting the right element in DevTools?
3. Are your classes being overridden by Bulma CSS?

**Example:**

```go
@Button(
    ButtonProps{
        Attr: templ.Attributes{
            "class": "my-custom-class",  // ✅ Merged with Bulma classes
        },
    },
    Items{Html("Button")},
)
```

**Output:**

```html
<button class="button my-custom-class">Button</button>
```

---

## Contributing

### How can I contribute to Bulma-Templ?

See [CONTRIBUTING.md](../CONTRIBUTING.md) for detailed guidelines.

**Welcome contributions:**
- Bug fixes
- Documentation improvements
- Test coverage
- Examples
- Accessibility enhancements

**Out of scope:**
- New components beyond Bulma
- Framework-specific integrations
- State management features
- Custom styling/theming

### How do I report a bug?

1. Check existing [GitHub issues](https://github.com/mtzvd/bulma-templ/issues)
2. Create a new issue with:
   - Go version
   - Templ version
   - Minimal reproduction code
   - Expected vs actual behavior

### How do I request a new component?

**Check first:**
- Is it part of [Bulma CSS](https://bulma.io/documentation/)?
- If yes → create an issue or submit a PR
- If no → out of scope (Bulma-first principle)

### Can I add convenience methods or helper functions?

**Generally no.** The library follows a strict 1:1 Bulma mapping principle.

**Exception:** Infrastructure primitives like `BaseElement`, `Items`, `Html()` that support the core architecture.

### How do I run tests locally?

```bash
# Run all tests
go test ./...

# Run with coverage
go test ./... -cover

# Run specific package
go test ./elements -v

# Run with race detection
go test ./... -race
```

### How do I add a new component?

See [CONTRIBUTING.md](../CONTRIBUTING.md) for the full process.

**Quick summary:**

1. Check Bulma documentation for HTML structure
2. Create `componentname.templ` in appropriate package
3. Define `ComponentNameProps` struct
4. Implement component following existing patterns
5. Add tests in `package_test.go`
6. Run `templ generate`
7. Submit PR

---

## Additional Resources

- **Bulma Documentation:** https://bulma.io/documentation/
- **Templ Documentation:** https://templ.guide
- **Project Documentation:**
  - [CANONICAL_PROJECT_CONTEXT.md](CANONICAL_PROJECT_CONTEXT.md) — Architecture
  - [DESIGN_SYSTEM.md](DESIGN_SYSTEM.md) — Design patterns
  - [CONTRIBUTING.md](../CONTRIBUTING.md) — Contribution guide
  - [COMMENT_STYLE.md](COMMENT_STYLE.md) — Code style

---

**Questions not answered here?** [Open an issue](https://github.com/mtzvd/bulma-templ/issues) or check the [documentation](.).
