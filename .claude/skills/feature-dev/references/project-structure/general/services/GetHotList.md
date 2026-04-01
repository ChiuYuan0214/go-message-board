# GetHotList

**File:** `backend/general/services/articles.go`

## Signature

```go
func GetHotList(page, size int64, userId uint64) []types.ArticleListData
```

## Behaviour

Gets hot articles from a Redis list maintained by a background job. Order reflects hotness score computed by the scheduler.
