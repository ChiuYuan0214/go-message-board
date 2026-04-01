# GetNewestList

**File:** `backend/general/services/articles.go`

## Signature

```go
func GetNewestList(page, size int64, userId uint64) []types.ArticleListData
```

## Parameters

| Name | Type | Description |
|------|------|-------------|
| `page` | `int64` | Page number (1-based) |
| `size` | `int64` | Items per page |
| `userId` | `uint64` | Caller — for `myScore` and `hasCollec` |

## Returns

Array of [ArticleListData](../types/ArticleListData.md). Empty slice on error.

## Behaviour

Raw SQL with pagination. Sorted by `publish_time DESC`. Includes author, top comment, and vote data.
Tags are **not** populated — caller handles that.
