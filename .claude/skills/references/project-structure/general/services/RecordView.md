# RecordView

**File:** `backend/general/services/view.go`

## Signature

```go
func RecordView(articleId string)
```

## Behaviour

Increments the article's view count in Redis (sorted set). A background job later flushes these counts to MySQL and updates the view-list cache.
Fire-and-forget — no return value.
