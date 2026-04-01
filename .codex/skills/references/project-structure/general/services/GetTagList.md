# GetTagList

**File:** `backend/general/services/articles.go`

## Signature

```go
func GetTagList(page, size int64, tag string) []types.ArticleListData
```

## Parameters

| Name | Type | Description |
|------|------|-------------|
| `tag` | `string` | Tag name to filter by |

## Behaviour

Returns articles that have the given tag, via join on `article_tag_maps` and `tags`.
Note: `userId` is not passed — `myScore` and `hasCollec` are always `0`/`false`.
