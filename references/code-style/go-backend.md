# Go Backend Code Style

Applies to all backend services. Prefer the current depin-based structure over older package-level patterns.

## Structure

- Follow dependency order: `entities -> types -> repo -> service -> route`.
- In `repo/`, `service/`, and `routes/`, prefer one file and one main struct.
- Use narrow interfaces per responsibility.
- When adding or changing methods in `repo/`, `service/`, or `routes/`, update the matching `interface.go` in the same folder in the same task.
- Register dependencies in `main.go`, then call `depin.Run()` before serving traffic.

## Repo Style

- Repos depend on `infra`, not on routes.
- Return data plus `error`, never HTTP status.
- Add compile guards such as `var _ Article = (*ArticleImpl)(nil)`.
- Named returns are fine when they keep the method short and clear.

```go
type ArticleImpl struct {
    db infra.RDB
}

func (r *ArticleImpl) GetArticleById(articleId uint64) (article types.Article, err error) {
    err = r.db.Orm().Where("article_id = ?", articleId).First(&article).Error
    return
}
```

## Service Style

- Services depend on repos or other services, not raw globals.
- Services own validation, ownership checks, workflow coordination, and status translation.
- Common return shapes are `(data, int)`, `(string, int)`, `bool`, or `uint64`, depending on caller needs.
- Log unexpected infra/repo errors before returning.

```go
type ArticleImpl struct {
    articleRepo repo.Article
}

func (s *ArticleImpl) GetArticle(userId uint64, articleId string) (*types.Article, int) {
    article, err := s.articleRepo.GetArticleDetail(userId, articleId)
    if err != nil {
        log.Println(err)
        return nil, http.StatusInternalServerError
    }
    return &article, 0
}
```

## Route Style

- Handler structs hold `router Router` plus only the service interfaces they need.
- `Run()` registers routes.
- Keep handlers thin: parse input, call services, shape response.
- For WebSocket handlers, keep upgrade logic in routes and delegate token checks, event loops, history, and notifications to services.

```go
type ArticleHandler struct {
    router         Router
    articleService service.Article
}

func (h *ArticleHandler) Run() (err error) {
    h.router.Get("/article", h.get)
    return
}
```

## Entities and Types

- Use explicit struct tags.
- Keep request/response types separate from persistence entities when helpful.
- Avoid extra validation tags unless the repo already relies on them.

## Naming

- Methods: verb-first PascalCase like `GetArticle`, `InsertArticle`
- Impl structs: `<Domain>Impl`
- Handler structs: `<Domain>Handler`
- Private helper methods: lowercase
- Error messages: lowercase, usually end with `.`

## Imports

Use standard grouping:
- stdlib
- third-party
- local module

## Local Readability Rules

- For keyed struct literals, put each field on its own line. Avoid single-line forms like `X{A: 1, B: 2}`.
- Leave one blank line between adjacent code blocks when the variables in those blocks are not directly part of the same local flow.
- For zero-value struct creation, prefer `var x X` over `x := X{}`.
- For empty pointer-to-struct creation, prefer `new(X)` over `&X{}`.
- For empty slice declarations, prefer `var xs []T` over `xs := []T{}` unless you truly need a non-nil empty slice at that point.

## Google-Style and Linting Guidance

- Keep interfaces small and consumer-oriented.
- Prefer concrete types until an interface is needed for injection or testing.
- Avoid unnecessary abbreviations in exported names.
- Handle errors explicitly; do not silently discard them.
- Prefer early returns to reduce nesting.
- Keep methods focused on one responsibility.
- Remove dead code, stale shims, and unused imports promptly.
- Optimize for readability over cleverness.

## Avoid

- Package-level dependency holders like `var db`, `var cache`, `var dynamo`
- Standalone service logic when the codebase expects struct methods
- Forgetting `depin.Run()` after dependency registration
- Collapsing multiple repo/service/route responsibilities into one large struct without a strong reason
