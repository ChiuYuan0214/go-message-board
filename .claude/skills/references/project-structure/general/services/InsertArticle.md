# InsertArticle

**File:** `backend/general/services/article.go`

## Signature

```go
func InsertArticle(userId uint64, article *types.AddArticleData, publishTime *time.Time) uint64
```

## Parameters

| Name | Type | Description |
|------|------|-------------|
| `userId` | `uint64` | Article author |
| `article` | `*types.AddArticleData` | Title, content, publishTime, tags |
| `publishTime` | `*time.Time` | Pre-parsed publish time |

## Returns

| Value | Notes |
|-------|-------|
| `uint64` | New article ID; `0` on error |

## Behaviour

Creates an `articles` record via `db.Create`. The caller is responsible for calling [InsertTags](InsertTags.md) afterward.
