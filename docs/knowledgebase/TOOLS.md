# Documentation Tools for Go Projects

Recommended tools for serving and publishing the Bulma-Templ knowledge base.

## Go-Based Documentation Tools

### 1. **Hugo** (Recommended)

**Best for:** Static site generation with Go

- **Language:** Go
- **Setup Time:** 5-10 minutes
- **Features:**
  - Blazing fast build times
  - Built-in markdown support
  - Excellent theming (Docsy, Book, Learn)
  - Automatic navigation from file structure
  - Full-text search with Algolia/Lunr.js
  - Hot reload during development

**Installation:**
```bash
# Windows (Chocolatey)
choco install hugo-extended

# macOS
brew install hugo

# Linux
snap install hugo
```

**Quick Setup:**
```bash
# Create Hugo site
hugo new site docs-site
cd docs-site

# Install documentation theme (Docsy recommended)
git clone https://github.com/google/docsy.git themes/docsy

# Copy knowledgebase to content/
cp -r ../docs/knowledgebase/* content/docs/

# Configure hugo.toml
cat > hugo.toml << EOF
baseURL = 'https://bulma-templ.dev/'
title = 'Bulma-Templ Documentation'
theme = 'docsy'

[params]
  description = "Type-safe Bulma components for Go"
  copyright = "Bulma-Templ Contributors"
EOF

# Run dev server
hugo server

# Build for production
hugo
```

**Pros:**
- ✅ Written in Go (same ecosystem)
- ✅ Extremely fast builds
- ✅ Rich theme ecosystem
- ✅ Great for technical documentation

**Cons:**
- ⚠️ Requires learning Hugo templating
- ⚠️ Themes can be complex to customize

**Best Themes for API Docs:**
- [Docsy](https://www.docsy.dev/) - Google's documentation theme
- [Book](https://themes.gohugo.io/themes/hugo-book/) - Clean, simple
- [Learn](https://learn.netlify.app/) - Technical docs focused

---

### 2. **MkDocs**

**Best for:** Python-based simple documentation

- **Language:** Python (but works great with Go projects)
- **Setup Time:** 3-5 minutes
- **Features:**
  - Simple configuration
  - Material theme is beautiful
  - Built-in search
  - Code highlighting
  - Easy deployment

**Installation:**
```bash
pip install mkdocs mkdocs-material
```

**Quick Setup:**
```bash
# Create mkdocs project
mkdocs new docs-site
cd docs-site

# Configure mkdocs.yml
cat > mkdocs.yml << EOF
site_name: Bulma-Templ Documentation
theme:
  name: material
  palette:
    primary: indigo
    accent: light-blue
  features:
    - navigation.tabs
    - navigation.sections
    - toc.integrate
    - search.suggest

nav:
  - Home: index.md
  - Components:
    - Elements: components/elements/
    - Components: components/components/
    - Form: components/form/
    - Layout: components/layout/
EOF

# Copy docs
cp -r ../docs/knowledgebase/components docs/

# Run dev server
mkdocs serve

# Build for production
mkdocs build
```

**Pros:**
- ✅ Simplest setup
- ✅ Beautiful Material theme
- ✅ Excellent search
- ✅ Great mobile support

**Cons:**
- ⚠️ Not Go-based
- ⚠️ Requires Python

---

### 3. **VitePress**

**Best for:** Modern, Vue-based documentation

- **Language:** JavaScript/TypeScript
- **Setup Time:** 5-10 minutes
- **Features:**
  - Fast Vite-powered builds
  - Modern UI/UX
  - Vue components in markdown
  - Excellent code highlighting
  - Built-in search

**Installation:**
```bash
npm init vitepress@latest
```

**Quick Setup:**
```bash
# Create project
mkdir docs-site && cd docs-site
npm init -y
npm install -D vitepress

# Initialize VitePress
npx vitepress init

# Copy knowledgebase
cp -r ../docs/knowledgebase/* docs/

# Run dev server
npm run docs:dev

# Build for production
npm run docs:build
```

**Pros:**
- ✅ Modern, fast, beautiful
- ✅ Great developer experience
- ✅ Vue ecosystem

**Cons:**
- ⚠️ Requires Node.js
- ⚠️ Not Go-based

---

### 4. **Docusaurus**

**Best for:** React-based documentation with versioning

- **Language:** JavaScript/TypeScript (React)
- **Setup Time:** 10-15 minutes
- **Features:**
  - Powerful versioning
  - i18n support
  - Plugin ecosystem
  - MDX support (JSX in Markdown)

**Installation:**
```bash
npx create-docusaurus@latest docs-site classic
```

**Pros:**
- ✅ Enterprise-grade features
- ✅ Versioning support
- ✅ React ecosystem

**Cons:**
- ⚠️ More complex
- ⚠️ Slower builds than Hugo/VitePress
- ⚠️ Not Go-based

---

### 5. **Docsify** (Zero Build)

**Best for:** Ultra-simple setup without build step

- **Language:** JavaScript (runtime)
- **Setup Time:** 2 minutes
- **Features:**
  - No build step required
  - Renders markdown at runtime
  - Easy deployment (just static files)
  - Plugins for search, theming, etc.

**Quick Setup:**
```bash
# Create index.html
cat > index.html << 'EOF'
<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <title>Bulma-Templ Docs</title>
  <link rel="stylesheet" href="//cdn.jsdelivr.net/npm/docsify/themes/vue.css">
</head>
<body>
  <div id="app"></div>
  <script>
    window.$docsify = {
      name: 'Bulma-Templ',
      loadSidebar: true,
      subMaxLevel: 2,
      search: 'auto'
    }
  </script>
  <script src="//cdn.jsdelivr.net/npm/docsify/lib/docsify.min.js"></script>
  <script src="//cdn.jsdelivr.net/npm/docsify/lib/plugins/search.min.js"></script>
</body>
</html>
EOF

# Copy docs
cp -r docs/knowledgebase/* .

# Serve with any web server
python3 -m http.server 3000
```

**Pros:**
- ✅ Simplest possible setup
- ✅ No build step
- ✅ Works with any HTTP server

**Cons:**
- ⚠️ Runtime markdown parsing (slower)
- ⚠️ Less features than build-based tools

---

## Recommendation Matrix

| Tool | Setup Time | Build Speed | Features | Go Native | Best For |
|------|-----------|-------------|----------|-----------|----------|
| **Hugo** | ⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ✅ | **Go projects** |
| **MkDocs** | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐⭐⭐ | ❌ | Quick setup |
| **VitePress** | ⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ❌ | Modern UI |
| **Docusaurus** | ⭐⭐⭐ | ⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ❌ | Versioning |
| **Docsify** | ⭐⭐⭐⭐⭐ | ⭐⭐ | ⭐⭐⭐ | ❌ | Zero-config |

## Our Recommendation: **Hugo with Docsy Theme**

**Why:**
1. ✅ **Go-native** - Same language as Bulma-Templ
2. ✅ **Lightning fast** - Builds 106 pages in milliseconds
3. ✅ **Docsy theme** - Designed for component documentation
4. ✅ **Automatic sidebar** - Generated from file structure
5. ✅ **Search built-in** - Algolia or Lunr.js integration
6. ✅ **GitHub Pages ready** - Easy deployment

## Quick Start with Hugo + Docsy

```bash
# 1. Install Hugo
choco install hugo-extended  # Windows
# or
brew install hugo           # macOS

# 2. Create site
hugo new site bulma-templ-docs
cd bulma-templ-docs

# 3. Add Docsy theme as submodule
git init
git submodule add https://github.com/google/docsy.git themes/docsy
git submodule update --init --recursive

# 4. Configure hugo.toml
cat > hugo.toml << EOF
baseURL = "https://bulma-templ.dev/"
title = "Bulma-Templ Documentation"
theme = "docsy"

[params]
  description = "Type-safe Bulma components for Go using Templ"
  github_repo = "https://github.com/mtzvd/bulma-templ"

[params.ui]
  sidebar_menu_compact = true
  navbar_logo = true

[menu]
  [[menu.main]]
    name = "Documentation"
    weight = 10
    url = "/docs/"
  [[menu.main]]
    name = "GitHub"
    weight = 50
    url = "https://github.com/mtzvd/bulma-templ"
EOF

# 5. Copy knowledgebase content
mkdir -p content/docs
cp -r ../docs/knowledgebase/components content/docs/

# 6. Create index page
cat > content/_index.md << EOF
---
title: "Bulma-Templ"
---

# Type-Safe Bulma Components for Go

Complete documentation for all 106 components.
EOF

# 7. Run dev server
hugo server -D

# 8. Build for production
hugo

# 9. Deploy to GitHub Pages
# Output is in public/ directory
```

## Deployment Options

### GitHub Pages (Free)

```bash
# Build
hugo

# Deploy (using gh-pages branch)
git subtree push --prefix public origin gh-pages
```

### Netlify (Free)

1. Connect GitHub repo
2. Set build command: `hugo`
3. Set publish directory: `public`

### Vercel (Free)

1. Import project
2. Framework preset: Hugo
3. Build command: `hugo`
4. Output directory: `public`

## Integration with docgen

Update `cmd/docgen/main.go` to add Hugo frontmatter format:

```go
// Hugo-specific frontmatter
sb.WriteString("---\n")
sb.WriteString(fmt.Sprintf("title: \"%s\"\n", component.Name))
sb.WriteString(fmt.Sprintf("description: \"%s\"\n", component.Description))
sb.WriteString(fmt.Sprintf("weight: %d\n", weight)) // For ordering
sb.WriteString("---\n\n")
```

---

## Conclusion

For a **Go project like Bulma-Templ**, we recommend:

1. **First choice:** Hugo + Docsy theme (Go-native, fast, feature-rich)
2. **Alternative:** MkDocs Material (simplest setup, beautiful)
3. **For modern UI:** VitePress (if Node.js is acceptable)

All options work great with the auto-generated markdown from `docgen`!
