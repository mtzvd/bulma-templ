# Documentation Generation — Bulma-Templ Design (Canonical)

This document defines the documentation generation pipeline for the **Bulma-Templ Design System**.

This document is **normative**: all future documentation tooling MUST comply with it.

**Version:** v1.0  
**Reference Code:** DOCGEN-PLAN-V1

---

## 1. Purpose

The documentation generation system extracts component documentation directly from `.templ` source files and generates human-readable Markdown documentation.

The system serves three purposes:

1. **Single Source of Truth**: Documentation lives in code, not separate files
2. **Discoverability**: Generate browsable component reference documentation
3. **Example Validation**: Ensure Example blocks remain accurate and copy-paste-able

The generated documentation is a reference guide, not a tutorial or learning resource.

---

## 2. Non-Goals

The documentation generation system explicitly does **NOT**:

- Generate interactive documentation or live previews
- Execute or validate Example blocks at generation time
- Parse or understand Go/templ semantics
- Infer missing documentation from code structure
- Generate API documentation for non-component code
- Support multiple output formats (only Markdown)
- Generate tutorials, guides, or learning materials
- Extract documentation from Kitchen Sink examples
- Perform static analysis or type checking
- Support annotations, tags, or DSL syntax
- Auto-generate Example blocks from usage patterns
- Create cross-references or dependency graphs
- Generate changelog or migration guides
- Support versioned documentation
- Provide search functionality
- Generate table of contents beyond simple lists

---

## 3. Source of Truth

### Primary Source

`.templ` files in the project are the **single source of truth** for component documentation.

### Excluded Sources

The following are **explicitly excluded** as documentation sources:

- Kitchen Sink examples (`examples/kitchensink/`)
- Test files (`*_test.go`)
- Generated Go files (`*_templ.go`)
- Markdown files in `docs/` (these document architecture, not components)
- Code comments outside component-level declarations
- Git history or commit messages
- External documentation sites

### Discovery

The generator discovers components by:

1. Scanning directories: `elements/`, `form/`, `components/`, `layout/`, `columns/`, `grid/`
2. Reading all `.templ` files (NOT `*_templ.go`)
3. Extracting component-level comments that match the canonical format

---

## 4. Extracted Data Model

For each component, the generator extracts:

### Required Fields

- **Component Name**: Extracted from type declaration (e.g., `ButtonProps` → `Button`)
- **Title Line**: First line of component comment (e.g., `Button — Bulma button component.`)
- **Atomic Level**: Exact string after `Atomic level:` (ATOM, MOLECULE, or ORGANISM)
- **Description**: All lines between atomic level and Example block or type declaration
- **Props Fields**: Field names, types, and comments from Props struct

### Optional Fields

- **Example Block**: Content of `Example:` section if present
- **Non-Responsibility**: Lines containing `does NOT`, `is NOT`, `MUST NOT`, etc.

### Non-Extracted Data

The generator does **NOT** extract:

- Function/method implementations
- Template body content
- Variable declarations outside Props
- Import statements
- Internal helper functions
- Test code

---

## 5. Parsing Rules

### Comment Block Detection

Component-level comments are detected by:

1. Finding a line matching: `// ComponentName — description.`
2. Continuing until reaching `type ComponentNameProps struct {`
3. All lines in between are part of the component comment block

### Atomic Level Extraction

The atomic level is extracted by exact regex match:

```
^// Atomic level: (ATOM|MOLECULE|ORGANISM)$
```

**Rules**:
- Case-sensitive
- Must be exact match (no variants)
- Must appear after title line
- Must appear before description

### Example Block Detection

Example blocks are detected by:

1. Finding a line exactly matching: `// Example:`
2. Starting a code block on the next line
3. Continuing until a non-comment line or type declaration

**Format Rules**:
- Example block starts with `// Example:`
- Following line is blank comment `//`
- Code lines are indented and prefixed with `//   ` (two spaces)
- Block ends at first non-comment line

### Field Comment Extraction

For each field in Props struct:

1. Locate comment line(s) immediately preceding field declaration
2. Extract comment text after `//`
3. Preserve multi-line field comments
4. Stop at field declaration line

### Whitespace Handling

Parser is **tolerant** of:
- Extra spaces after `//`
- Inconsistent indentation in Example blocks
- Blank lines within comment blocks
- CRLF vs LF line endings

Parser is **strict** about:
- `// Example:` must be exact (no variations)
- Atomic level values (ATOM, MOLECULE, ORGANISM only)
- Title line format (em dash `—` required)

---

## 6. Example Block Contract

### Purpose

Example blocks demonstrate **minimal, canonical usage** of a component using only its public API.

### Format Requirements

Example blocks MUST:

- Use valid Go/templ syntax (conceptually)
- Show component initialization with Props struct
- Use `elements.Items` for composition
- Be copy-paste-able (with appropriate imports)
- Use only public API (exported types and functions)

Example blocks MUST NOT:

- Show edge cases or advanced usage
- Include tutorial-style comments
- Demonstrate multiple patterns
- Use unexported helpers
- Reference Kitchen Sink code
- Show validation or error handling
- Include setup or teardown code

### Canonical Patterns

**Atom Example**:
```go
Button(
    ButtonProps{
        Color: "is-primary",
    },
    elements.Items{
        elements.Html("Click me"),
    },
)
```

**Molecule Example**:
```go
Field(
    FieldProps{},
    elements.Items{
        Label(LabelProps{}, elements.Items{elements.Html("Name")}),
        Control(ControlProps{}, elements.Items{
            Input(InputProps{Type: "text"}),
        }),
    },
)
```

**Organism Example**:
```go
Navbar(
    NavbarProps{
        Color: "is-primary",
    },
    elements.Items{
        NavbarBrand(NavbarBrandProps{}, elements.Items{...}),
        NavbarMenu(NavbarMenuProps{}, elements.Items{...}),
    },
)
```

### Interpretation

The generator:

- Extracts Example blocks as-is (no transformation)
- Preserves indentation and formatting
- Does NOT execute or validate examples
- Does NOT generate syntax highlighting metadata

---

## 7. Markdown Output Structure

### File Organization

Generated documentation structure:

```
docs/generated/
    index.md              # Component index by atomic level
    elements/
        Button.md
        Icon.md
        ...
    form/
        Input.md
        Field.md
        ...
    components/
        Card.md
        Navbar.md
        ...
    layout/
        Container.md
        ...
    columns/
        Columns.md
        ...
    grid/
        Grid.md
        ...
```

### Component Page Template

Each component page follows this structure:

```markdown
# ComponentName

**Atomic Level:** ATOM/MOLECULE/ORGANISM

Brief description from title line.

## Description

Full description text from component comment.

## Props

| Field | Type | Description |
|-------|------|-------------|
| Color | string | Bulma color modifier. |
| Size  | string | Bulma size modifier. |
| ...   | ...    | ... |

## Example

\`\`\`go
Button(
    ButtonProps{
        Color: "is-primary",
    },
    elements.Items{
        elements.Html("Click me"),
    },
)
\`\`\`

## Non-Responsibilities

- Does NOT manage state
- Does NOT handle routing
```

### Index Page Structure

The index page groups components by atomic level:

```markdown
# Bulma-Templ Component Reference

## Atoms

- [Button](elements/Button.md)
- [Icon](elements/Icon.md)
- ...

## Molecules

- [Field](form/Field.md)
- [CardHeader](components/CardHeader.md)
- ...

## Organisms

- [Navbar](components/Navbar.md)
- [Card](components/Card.md)
- ...
```

### Metadata

Each page includes metadata header:

```markdown
---
component: Button
atomic_level: ATOM
package: elements
generated: 2026-01-12T10:30:00Z
---
```

---

## 8. Planned CLI Interface (Conceptual)

### Command Structure

```bash
bulma-templ-docgen [flags]
```

### Input

The CLI reads:
- Source `.templ` files from workspace root
- No configuration files
- No external metadata

### Flags

**Required**:
- `--output <dir>`: Output directory for generated Markdown (default: `docs/generated`)

**Optional**:
- `--clean`: Remove output directory before generation
- `--verbose`: Print detailed parsing information
- `--check`: Validate documentation completeness without generating files

**Explicitly NOT Supported**:
- Custom templates or output formats
- Filtering by package or component
- Incremental generation
- Watch mode
- Preview server

### Output

The CLI:
- Creates output directory if it doesn't exist
- Generates one Markdown file per component
- Generates index.md with component list
- Prints summary: `Generated documentation for N components`

### Exit Codes

- `0`: Success
- `1`: Parsing error (malformed comment blocks)
- `2`: File I/O error
- `3`: Invalid flags

### Example Usage

```bash
# Generate documentation
bulma-templ-docgen --output docs/generated

# Clean and regenerate
bulma-templ-docgen --output docs/generated --clean

# Validate without generating
bulma-templ-docgen --check
```

---

## 9. Failure Modes and Edge Cases

### Parsing Failures

**Missing Atomic Level**:
- Behavior: Skip component, log warning
- Rationale: Atomic level is required per COMMENT_STYLE.md

**Malformed Example Block**:
- Behavior: Include Example section with "Parse error" message
- Rationale: Preserve partial documentation, signal issue

**Missing Title Line**:
- Behavior: Skip component, log error
- Rationale: Cannot generate documentation without component name

**Inconsistent Indentation**:
- Behavior: Preserve as-is, do not reformat
- Rationale: Respect author's formatting choices

### Edge Cases

**Multiple Example Blocks**:
- Behavior: Use first Example block only
- Rationale: Canonical format allows exactly one Example

**Props Struct Not Found**:
- Behavior: Generate documentation without Props table
- Rationale: Some helpers may not have Props

**Nested Comments**:
- Behavior: Treat all as single comment block
- Rationale: Simplicity over precision

**Unicode in Comments**:
- Behavior: Preserve UTF-8 as-is
- Rationale: Markdown supports Unicode

### Non-Failures

The following are **explicitly NOT errors**:

- Missing Example block (optional)
- Missing non-responsibility section (optional)
- Extra fields in Props beyond documented standard
- Comments in template body (ignored)
- Unused imports in `.templ` file

---

## 10. What This System Will NEVER Do

The documentation generation system is **intentionally limited** and will NEVER:

1. **Execute Code**: No evaluation of Example blocks or template logic
2. **Semantic Analysis**: No type inference, dependency resolution, or semantic parsing
3. **Cross-Referencing**: No links between related components (manual only)
4. **Validation**: No verification that examples compile or run
5. **Transformation**: No reformatting, prettifying, or rewriting of examples
6. **Annotation Support**: No tags, decorators, or structured metadata beyond standard comments
7. **Tutorial Generation**: No step-by-step guides or learning paths
8. **API Inference**: No automatic documentation from function signatures
9. **Version Management**: No changelog generation or version comparison
10. **Interactive Features**: No search, filtering, or dynamic navigation
11. **Multiple Formats**: No HTML, PDF, or other output formats
12. **Localization**: No translation or multi-language support
13. **Kitchen Sink Integration**: No extraction from example applications
14. **Auto-Completion**: No IDE integration or language server features
15. **Semantic Linking**: No automatic detection of component relationships

### Rationale

These limitations are **architectural decisions**, not implementation gaps:

- **Simplicity**: Keep the generator simple and maintainable
- **Reliability**: Avoid fragile parsing or semantic analysis
- **Stability**: Minimize dependencies and complexity
- **Conservatism**: Match project philosophy of explicit over implicit

---

## Final Notes

This design is **frozen** for v1.0. Changes require architectural review.

The documentation generator is a **reflection tool**, not an inference engine. It surfaces what is explicitly written, nothing more.

When in doubt, the generator does **less**, not more.
