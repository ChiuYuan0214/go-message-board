# GetViewList

**File:** `backend/general/services/articles.go`

## Signature

```go
func GetViewList(page, size int64, userId uint64) []types.ArticleListData
```

## Behaviour

Gets the top-viewed articles using article IDs stored in Redis (sorted set by view count).
Falls back gracefully if Redis is empty.
