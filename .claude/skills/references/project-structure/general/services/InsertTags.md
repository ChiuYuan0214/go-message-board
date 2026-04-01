# InsertTags

**File:** `backend/general/services/article.go`

## Signature

```go
func InsertTags(articleId uint64, tags []string) bool
```

## Parameters

| Name | Type | Description |
|------|------|-------------|
| `articleId` | `uint64` | Article to attach tags to |
| `tags` | `[]string` | Tag names to add |

## Returns

`true` on success, `false` on any DB error.

## Behaviour

1. Finds which tags already exist in `tags` table
2. Creates any missing tags
3. Inserts `article_tag_maps` rows for tags not already mapped to this article (idempotent)

## When to reuse

Call after [InsertArticle](InsertArticle.md) or alongside [UpdateArticle](UpdateArticle.md).
