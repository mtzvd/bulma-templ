# Branch Protection Setup on GitHub

## Quick Setup via Web UI

### 1. Open Repository Settings
https://github.com/mtzvd/bulma-templ/settings/branches

### 2. Change Default Branch to `develop`
- In "Default branch" section, click ↔️
- Select `develop`
- Confirm

### 3. Protect `main` Branch

Click **Add rule** and configure:

**Branch name pattern:** `main`

✅ **Require a pull request before merging**
   - Require approvals: 1
   - ✅ Dismiss stale pull request approvals when new commits are pushed

✅ **Require status checks to pass before merging**
   - ✅ Require branches to be up to date before merging
   - Search for status checks: `test`, `build`, `lint` (if CI/CD is configured)

✅ **Require conversation resolution before merging**

✅ **Do not allow bypassing the above settings** (applies to admins too)

### 4. Protect `staging` Branch

Click **Add rule** and configure:

**Branch name pattern:** `staging`

✅ **Require a pull request before merging**

✅ **Require status checks to pass before merging** (optional)
   - Search for status checks: `test`, `build`

### 5. `develop` Branch (Optional)

Can be left unprotected for faster development, or add minimal protection:

**Branch name pattern:** `develop`

✅ **Require status checks to pass before merging** (recommended)

---

## Result

After configuration:
- ❌ **Cannot** push directly to `main`
- ✅ **Only via Pull Request** with checks
- ✅ **Safe** — accidental commits to main are impossible
- ✅ **Code Review** — approval required before merge

## Workflow After Setup

```bash
# Development (can push directly)
git checkout develop
git commit -m "feat: Something"
git push origin develop

# Trying to push to main (will be rejected)
git checkout main
git push origin main  # ❌ ERROR: protected branch

# Correct way: via PR
# 1. develop -> staging (can be via PR or merge)
# 2. staging -> main (ONLY via PR)
```
