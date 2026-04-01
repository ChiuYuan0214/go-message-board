# go-message-board — Depin DI Architecture

The `general` service uses [depin](https://github.com/ChiuYuan0214/depin) for
dependency injection. Every component is a struct; there are no package-level
variables or standalone functions.

---

## Four Layers

```
infra/      External connections — *gorm.DB, Redis client
repo/       Raw DB queries — one impl per domain, returns (T, error)
service/    Business logic — holds injected repos, returns (T, int) or bool
routes/     HTTP handlers — holds injected Router + service, registers routes
```

---

## Layer Rules

### infra
- `infra.RDB` wraps `*gorm.DB` via `.Orm()`.
- `infra.Cache` wraps Redis.
- Never imported by `routes/`.

### repo
- One interface per domain in `repo/interface.go`.
- Impl struct holds `db infra.RDB`.
- Methods use named returns + bare `return`.
- Never return HTTP codes — only `(T, error)`.
- Compile guard at top: `var _ Foo = (*FooImpl)(nil)`.

### service
- One interface per domain in `service/interface.go`.
- Impl struct holds injected repo interfaces (and `infra.Cache` if needed).
- Business logic lives here: ownership checks, error-to-HTTP-code translation.
- Returns `(string, int)` for errors (message + status), or `bool`, or a data type.
- Compile guard at top: `var _ Foo = (*FooImpl)(nil)`.

### routes
- Each handler is a struct with `router Router` + one injected service interface.
- `Run()` registers all routes for that domain via `router.Get/Post/Put/Delete`.
- `Stop()` is a no-op `{}`.
- No `init*` functions, no package-level `var`.

---

## depin Lifecycle

Every struct must implement `Run() error` and `Stop()`.  
depin calls `Run()` after injecting all fields and `Stop()` on shutdown.

```go
// infra — use RunAndSet when you need the value (Router needs .Serve())
router := depin.RunAndSet[routes.Router](new(routes.RouterImpl))

// everything else — use Set
depin.Set[infra.RDB](new(infra.MySQL))
depin.Set[repo.Article](new(repo.ArticleImpl))
depin.Set[service.Article](new(service.ArticleImpl))
depin.Set[routes.Handler](new(routes.ArticleHandler))
```

---

## Registration Order in main.go

Always register bottom-up so depin can resolve dependencies:

```
1. routes.Router      (RunAndSet — needed for router.Serve() later)
2. infra.*            (RDB, Cache)
3. repo.*             (depend on infra)
4. service.*          (depend on repo + infra)
5. routes.Handler     (depend on routes.Router + service)
6. router.Serve()
```

---

## Adding a New Domain — Checklist

- [ ] `repo/interface.go` — add `type Foo interface { ... }`
- [ ] `repo/foo.go` — `FooImpl` with `db infra.RDB`, compile guard, named returns
- [ ] `service/interface.go` — add `type Foo interface { ... }`
- [ ] `service/foo.go` — `FooImpl` with injected repos, compile guard
- [ ] `routes/foo.go` — `FooHandler` with `router Router` + `fooService service.Foo`, `Run()` registers routes
- [ ] `main.go` — `depin.Set` for repo → service → handler (in that order)
