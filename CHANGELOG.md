<a name="unreleased"></a>
## [Unreleased]


<a name="v1.0.2"></a>
## [v1.0.2] - 2026-01-12
### Chore
- Merge main into staging
- Merge main into develop

### Docs
- Add Git workflow and branch protection guides

### Feat
- Add generated templ files and documentation for pkg.go.dev visibility
- Add automatic code formatting (gofmt + templ fmt)

### Fix
- Build golangci-lint with Go 1.25 in CI ([#2](https://github.com/mtzvd/bulma-templ/issues/2))
- Build golangci-lint from source with Go 1.25 in CI
- Remove duplicate issues section in golangci config
- Merge duplicate issues section in golangci config
- Update golangci-lint config (use exclude-dirs instead of skip-dirs)

### Style
- Format code with gofmt and templ fmt

### Test
- Remove test file
- Check pre-commit hook


<a name="v1.0.1"></a>
## [v1.0.1] - 2026-01-12
### Chore
- Update CHANGELOG for v1.0.1
- Prepare v1.0.1 release ([#1](https://github.com/mtzvd/bulma-templ/issues/1))

### Feat
- Add CHANGELOG automation with git-chglog


<a name="v1.0.0"></a>
## v1.0.0 - 2026-01-11
### Docs
- merge README for v1.0 (concise but complete)
- Update README and add DESIGN_SYSTEM documentation for clarity and consistency

### Feat
- complete CI/CD setup and project restructuring
- Add unit tests for pagination and button components; enhance documentation for clarity
- Add starter page example for Bulma-Templ design system
- Refactor layout components to use new elements package for rendering
- **examples:** add code display to kitchen sink examples

### Refactor
- update comments to use "defines" for consistency across components

### BREAKING CHANGE

Generated *_templ.go files are no longer tracked in git.
Users must run 'templ generate' after clone/pull.

This brings the project to production-ready state with comprehensive
CI/CD automation, multi-platform support, and clean git history.


[Unreleased]: https://github.com/mtzvd/bulma-templ/compare/v1.0.2...HEAD
[v1.0.2]: https://github.com/mtzvd/bulma-templ/compare/v1.0.1...v1.0.2
[v1.0.1]: https://github.com/mtzvd/bulma-templ/compare/v1.0.0...v1.0.1
