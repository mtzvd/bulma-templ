<a name="unreleased"></a>
## [Unreleased]

### Chore
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


[Unreleased]: https://github.com/mtzvd/bulma-templ/compare/v1.0.0...HEAD
