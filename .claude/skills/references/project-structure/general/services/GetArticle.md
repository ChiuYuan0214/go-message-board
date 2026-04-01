# GetArticle

**File:** `backend/general/services/article.go`

## Signature

```go
func GetArticle(userId uint64, articleId string) (*types.Article, int)
```

## Parameters

| Name | Type | Description |
|------|------|-------------|
| `userId` | `uint64` | Caller's user ID — used to compute `myScore` and `hasCollec` |
| `articleId` | `string` | Target article ID |

## Returns

| Value | Notes |
|-------|-------|
| `*types.Article` | Full article detail; nil on error |
| `int` | `0` = success, `http.StatusInternalServerError` on DB error |

## Behaviour

Single raw SQL query with subqueries that computes in one pass:
- Author info (joined from `users`, `images`)
- Vote counts (`voteUp`, `voteDown`) via COUNT subqueries
- Caller's vote score (`myScore`)
- Whether caller has collected the article (`hasCollec`)

Does **not** populate `Tags` — call [GetTagsByArticleId](GetTagsByArticleId.md) separately.

## When to reuse

Any endpoint that returns a single article's full detail view.
