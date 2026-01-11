# Git Workflow

## Branch Structure

- **`develop`** — Основная ветка разработки (default branch)
- **`staging`** — Тестирование перед релизом
- **`main`** — Только стабильные релизы (защищена)

## Workflow

### Ежедневная разработка

```bash
# Работа в develop
git checkout develop
git pull origin develop

# Внесение изменений
git add .
git commit -m "feat: Add new feature"
git push origin develop
```

### Подготовка к релизу

```bash
# 1. Merge develop в staging для тестирования
git checkout staging
git pull origin staging
git merge develop
git push origin staging

# 2. Тестирование на staging
# ...

# 3. Если всё OK, создаём PR: staging -> main
# Через GitHub UI создаём Pull Request
```

### Создание релиза

```bash
# После merge в main через PR:
git checkout main
git pull origin main

# Создаём релиз (автоматизировано)
task release -- v1.1.0

# Или вручную:
git tag v1.1.0
git push origin v1.1.0
gh release create v1.1.0 --notes-file release_notes.md
```

### Hotfix (срочное исправление в production)

```bash
# Создаём hotfix ветку от main
git checkout main
git pull origin main
git checkout -b hotfix/critical-bug

# Исправляем
git add .
git commit -m "fix: Critical bug in production"

# Создаём PR: hotfix/critical-bug -> main
# После merge, также merge в develop и staging
```

## Branch Protection Rules

Настройте защиту веток на GitHub:

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
git checkout develop   # для разработки
git checkout staging   # для тестирования
git checkout main      # только для чтения

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
