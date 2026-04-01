# GetProfileList

**File:** `backend/general/services/articles.go`

## Signature

```go
func GetProfileList(page, size int64, userId uint64, selfUserId uint64) []types.ArticleListData
```

## Parameters

| Name | Type | Description |
|------|------|-------------|
| `userId` | `uint64` | Caller — for `myScore` and `hasCollec` |
| `selfUserId` | `uint64` | The user whose articles to list |

## Behaviour

Returns articles authored by `selfUserId`, paginated, sorted by publish_time DESC.
