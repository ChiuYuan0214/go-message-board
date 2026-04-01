# Tag

**Table:** `tags`
**File:** `backend/general/entities/tags.go`

## Fields

| Field | Type | Column | Notes |
|-------|------|--------|-------|
| `TagId` | `uint64` | `tag_id` | PK, auto-increment |
| `Name` | `string` | `name` | Tag text |

## Purpose

Global tag registry. Tags are shared across articles; the many-to-many mapping lives in `article_tag_maps`.

## Related

- Entity: [ArticleTagMap](article-tag-map.md)
- Services: [InsertTags](../services/InsertTags.md), [DeleteRemovedTags](../services/DeleteRemovedTags.md), [GetTagsByArticleId](../services/GetTagsByArticleId.md)
