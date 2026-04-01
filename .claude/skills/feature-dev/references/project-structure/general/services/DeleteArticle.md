# DeleteArticle

**File:** `backend/general/services/article.go`

## Signature

```go
func DeleteArticle(userId uint64, articleId uint64) (string, int)
```

## Parameters

| Name | Type | Description |
|------|------|-------------|
| `userId` | `uint64` | Caller — ownership is validated |
| `articleId` | `uint64` | Article to delete |

## Returns

| Value | Notes |
|-------|-------|
| `string` | Error message; `""` on success |
| `int` | HTTP status on error; `0` on success |

## Behaviour

Cascade deletes in order: tags → votes → comments → article_tag_maps → collections → article.
Returns `400` if `userId` doesn't own the article.
