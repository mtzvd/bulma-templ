# Git Workflow

## Branch Structure

- **`develop`** — Main development branch (default branch)
- **`staging`** — Testing before release
- **`main`** — Stable releases only (protected)

## Workflow

### Daily Development

```bash
# Work in develop
git checkout develop
git pull origin develop

# Make changes
git add .
git commit -m "feat: Add new feature"
git push origin develop
```

### Release Preparation

```bash
# 1. Merge develop into staging for testing
git checkout staging
git pull origin staging
git merge develop
git push origin staging

# 2. Test on staging
# ...

# 3. If all OK, create PR: staging -> main
# Create Pull Request via GitHub UI
```

### Creating a Release

```bash
# After merge to main via PR:
git checkout main
git pull origin main

# Create release (automated)
task release -- v1.1.0

# Or manually:
git tag v1.1.0
git push origin v1.1.0
gh release create v1.1.0 --notes-file release_notes.md
```

### Hotfix (urgent production fix)

```bash
# Create hotfix branch from main
git checkout main
git pull origin main
git checkout -b hotfix/critical-bug

# Fix the issue
git add .
git commit -m "fix: Critical bug in production"

# Create PR: hotfix/critical-bug -> main
# After merge, also merge into develop and staging
```

## Branch Protection Rules

Configure branch protection on GitHub:

### Settings → Branches → Branch protection rules

**For `main` branch:**
- ✅ Require a pull request before merging
- ✅ Require approvals (1)
- ✅ Dismiss stale pull request approvals when new commits are pushed
- ✅ Require status checks to pass before merging
  - ✅ Require branches to be up to date before merging
  - Status checks: `test`, `lint`, `build`
- ✅ Require conversation resolution before merging
- ✅ Do not allow bypassing the above settings

**For `staging` branch:**
- ✅ Require a pull request before merging
- ✅ Require status checks to pass before merging
  - Status checks: `test`, `lint`, `build`

**For `develop` branch:**
- ✅ Require status checks to pass before merging (optional)
- Allow direct pushes for faster development

## Default Branch

Set `develop` as default branch:
1. Go to repository Settings
2. Branches → Default branch
3. Switch to `develop`

## Quick Commands

```bash
# Switch branches
git checkout develop   # for development
git checkout staging   # for testing
git checkout main      # read-only

# Sync staging with develop
git checkout staging && git merge develop && git push

# Check current branch
git branch

# See all branches
git branch -a
```

## CI/CD per Branch

- **develop**: Runs tests on every push
- **staging**: Runs tests + optional deployment to staging environment
- **main**: Runs tests + automatic release creation + deployment to production
