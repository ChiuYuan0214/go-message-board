# DeleteRemovedTags

**File:** `backend/general/services/article.go`

## Signature

```go
func DeleteRemovedTags(articleId uint64, tags []string) bool
```

## Parameters

| Name | Type | Description |
|------|------|-------------|
| `articleId` | `uint64` | Article being updated |
| `tags` | `[]string` | The new complete tag list |

## Returns

`true` on success, `false` on DB error.

## Behaviour

Deletes `article_tag_maps` rows for any tags that were previously mapped to the article
but are not in the new `tags` list. Use before [InsertTags](InsertTags.md) on update.
