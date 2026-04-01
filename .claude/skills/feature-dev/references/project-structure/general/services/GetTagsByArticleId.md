# GetTagsByArticleId

**File:** `backend/general/services/article.go`

## Signature

```go
func GetTagsByArticleId(articleId string) []string
```

## Returns

Array of tag name strings. Returns empty slice on error (never nil).

## Behaviour

Joins `tags` with `article_tag_maps` to get all tag names for the given article.
Returns `[]string{}` on DB error instead of surfacing the error.

## When to reuse

Call after [GetArticle](GetArticle.md) to populate `article.Tags`.
