# go-message-board — Depin Pattern

Use `depin` for backend dependency injection. Prefer struct-based components and avoid package-level mutable dependencies.

## Layers

```text
infra   external clients and connections
repo    data access
service business logic
routes  transport handlers
```

## Rules

- `repo/`, `service/`, and `routes/` follow one file, one main struct.
- Keep interfaces in `interface.go`, grouped by responsibility.
- Repos return data plus `error`, never HTTP status.
- Services own validation, orchestration, and route-facing status translation.
- Routes hold `router Router` plus only the service interfaces they need.
- WebSocket handlers may inject multiple services, but responsibilities should still stay split.

## Lifecycle

- Every registered component implements `Run() error` and `Stop()`.
- `depin.Set(...)` only registers dependencies.
- `depin.Run()` is required to inject fields and execute `Run()`.
- `depin.RunAndSet(...)` is useful when you need the initialized instance immediately, such as `routes.Router`.

```go
router := depin.RunAndSet[routes.Router](new(routes.RouterImpl))
depin.Set[infra.RDB](new(infra.MySQL))
depin.Set[repo.Article](new(repo.ArticleImpl))
depin.Set[service.Article](new(service.ArticleImpl))
depin.Set[routes.Handler](new(routes.ArticleHandler))
depin.Run()
router.Serve()
```

## Registration Order

```text
1. routes.Router
2. infra.*
3. repo.*
4. service.*
5. jobs / schedulers
6. routes.Handler
7. depin.Run()
8. router.Serve()
```

## Checklist

- Add interface in `repo/interface.go` or `service/interface.go`
- Add one impl struct in its own file
- Add compile guard at the top of impl files
- Register dependencies bottom-up in `main.go`
- Call `depin.Run()` before serving traffic
