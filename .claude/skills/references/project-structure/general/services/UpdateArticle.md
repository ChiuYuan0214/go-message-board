# UpdateArticle

**File:** `backend/general/services/article.go`

## Signature

```go
func UpdateArticle(userId uint64, articleId uint64, data *types.UpdateArticleData) (string, int)
```

## Parameters

| Name | Type | Description |
|------|------|-------------|
| `userId` | `uint64` | Caller — ownership is validated |
| `articleId` | `uint64` | Target article |
| `data` | `*types.UpdateArticleData` | Fields to update (title, content) |

## Returns

| Value | Notes |
|-------|-------|
| `string` | Error message; `""` on success |
| `int` | HTTP status on error; `0` on success |

## Behaviour

- Returns `400` if article not found
- Returns `400` if `userId != article.UserId`
- Only updates non-empty fields (partial update)
- Tags are handled separately by [DeleteRemovedTags](DeleteRemovedTags.md) + [InsertTags](InsertTags.md)
