# Настройка защиты веток на GitHub

## Быстрая настройка через Web UI

### 1. Откройте настройки репозитория
https://github.com/mtzvd/bulma-templ/settings/branches

### 2. Измените default branch на `develop`
- В разделе "Default branch" нажмите ↔️
- Выберите `develop`
- Подтвердите

### 3. Защита ветки `main`

Нажмите **Add rule** и настройте:

**Branch name pattern:** `main`

✅ **Require a pull request before merging**
   - Require approvals: 1
   - ✅ Dismiss stale pull request approvals when new commits are pushed

✅ **Require status checks to pass before merging**
   - ✅ Require branches to be up to date before merging
   - Search for status checks: `test`, `build`, `lint` (если есть CI/CD)

✅ **Require conversation resolution before merging**

✅ **Do not allow bypassing the above settings** (для админов тоже)

### 4. Защита ветки `staging`

Нажмите **Add rule** и настройте:

**Branch name pattern:** `staging`

✅ **Require a pull request before merging**

✅ **Require status checks to pass before merging** (опционально)
   - Search for status checks: `test`, `build`

### 5. Ветка `develop` (опционально)

Можно не защищать для быстрой разработки, или добавить минимальную защиту:

**Branch name pattern:** `develop`

✅ **Require status checks to pass before merging** (рекомендуется)

---

## Результат

После настройки:
- ❌ **Нельзя** делать прямой push в `main`
- ✅ **Только через Pull Request** с проверками
- ✅ **Безопасно** — случайный коммит в main невозможен
- ✅ **Code Review** — требуется одобрение перед merge

## Workflow после настройки

```bash
# Разработка (можно пушить напрямую)
git checkout develop
git commit -m "feat: Something"
git push origin develop

# Попытка запушить в main (будет отклонено)
git checkout main
git push origin main  # ❌ ERROR: protected branch

# Правильный способ: через PR
# 1. develop -> staging (можно через PR или merge)
# 2. staging -> main (ТОЛЬКО через PR)
```
