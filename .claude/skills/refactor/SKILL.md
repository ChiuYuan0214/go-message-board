---
name: refactor
description: >
  Repository-pattern migration skill for the go-message-board general service.
  Use this skill whenever the user wants to refactor, restructure, optimize, or
  migrate service code — especially when keywords like refactor, 重構, cleanup,
  整理, optimize, 優化 appear, or when they ask to migrate functions from the
  global-db pattern to the repo/service struct pattern. Also trigger when the
  user asks to continue, finish, or complete an ongoing refactor, or references
  the infra/repo/services layering.
---

# Refactor Skill — Repository Pattern Migration

## Before Anything Else — Load Context

Read these before writing any code:

1. **`.claude/skills/rules/engineering.md`** — rules that apply to every task
2. **`.claude/skills/references/code-style/go-backend.md`** — naming conventions, GORM patterns, import style
3. **`.claude/skills/references/project-structure/README.md`** — existing service functions and types (avoid re-implementing what already exists)

---

This project's general service is being migrated from standalone service
functions (using a global `db *gorm.DB`) to a layered repository pattern.
Always read the current state of the files before writing anything.

## The Three-Layer Architecture

```
infra/          RDB interface wrapping *gorm.DB
repo/           per-domain interface + ArticleImpl (raw SQL lives here)
services/       per-domain struct with injected repo; business logic
routes/         wires up structs, calls service methods
```

### Layer contracts

**`infra/interface.go`** — already done, do not modify:
```go
type RDB interface { Orm() *gorm.DB }
```

**`repo/interface.go`** — one interface per domain, add methods as needed:
```go
type Article interface {
    GetArticleDetail(userId uint64, articleId string) (types.Article, error)
    // add new methods here
}
```

**`repo/article.go`** — `ArticleImpl` implements `repo.Article` via `infra.RDB`:
```go
type ArticleImpl struct { db infra.RDB }

func (r *ArticleImpl) SomeMethod(...) (..., error) {
    err = r.db.Orm()./* gorm chain */.Error
    return
}
```

**`services/article.go`** — `ArticleImpl` holds injected repos, exposes business methods:
```go
type ArticleImpl struct { articleRepo repo.Article }

func (s *ArticleImpl) SomeMethod(...) (..., int) {
    result, err := s.articleRepo.SomeRepoMethod(...)
    if err != nil { ... return nil, http.StatusInternalServerError }
    return result, 0
}
```

## Migration Workflow (per function)

For each standalone service function that still uses the global `db`:

1. **Read** the function in `services/article.go`. Identify the SQL/GORM query.

2. **`repo/interface.go`** — add the method signature to the `Article` interface.
   Return `(T, error)` — never HTTP codes from repo.

3. **`repo/article.go`** — implement the method on `repo.ArticleImpl`.
   - Use `r.db.Orm()` to access GORM.
   - Use named return values + bare `return` for brevity (see existing `GetArticleDetail`).

4. **`services/article.go`** — convert the standalone function to a method on
   `services.ArticleImpl`. It should call `s.articleRepo.TheNewRepoMethod(...)`.
   Keep HTTP status codes and business-logic checks here, not in repo.
   Remove the original standalone function.

5. **`routes/article.go`** — update any call sites that used the old standalone
   `services.FuncName(...)` to use the service struct method instead.
   Ensure `ArticleHandler` holds (or can access) an `*services.ArticleImpl`.

6. **Verify** the compile still works conceptually: no leftover references to
   the removed standalone function, no direct `db.` usage in services.

## Remaining Work in article.go (as of 2026-04-01)

These standalone functions still use the global `db` and need migrating:

| Function | Notes |
|---|---|
| `InsertArticle` | creates article, returns new ID |
| `UpdateArticle` | ownership check + partial update |
| `DeleteArticle` | cascades to votes, comments, tags, collections |
| `GetTagsByArticleId` | joins tags via article_tag_maps |
| `InsertTags` | upsert tags + article_tag_maps |
| `DeleteRemovedTags` | removes stale tag associations |

## Conventions to Follow

- **No HTTP codes in repo** — repo returns `error`; service translates to codes.
- **Named returns** in repo methods — follow the style of `GetArticleDetail`.
- **Compile guard** at top of repo file: `var _ Article = (*ArticleImpl)(nil)`
- **One interface per domain** — `repo.Article`, `repo.Comment`, etc.
- **Don't touch unrelated files** — scope each change to one domain or one layer.
- **Service struct, not package-level vars** — no new `var db` additions anywhere.

## Checklist Before Finishing

- [ ] `repo/interface.go` has all new method signatures
- [ ] `repo/article.go` implements them all (compile guard passes)
- [ ] `services/article.go` has no remaining standalone functions using global `db`
- [ ] `routes/article.go` calls service methods, not old standalone functions
- [ ] Project compiles (`go build ./...` in `backend/general/`)
