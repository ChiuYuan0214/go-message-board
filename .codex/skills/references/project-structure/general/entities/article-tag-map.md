# ArticleTagMap

**Table:** `article_tag_maps`
**File:** `backend/general/entities/article-tag-map.go`

## Fields

| Field | Type | Column | Notes |
|-------|------|--------|-------|
| `ArticleId` | `uint64` | `article_id` | Composite PK |
| `TagId` | `uint64` | `tag_id` | Composite PK |

## Purpose

Join table linking articles to tags. A single article can have multiple tags;
a tag can belong to multiple articles.

## Related

- Entity: [Tag](tag.md)
- Services: [InsertTags](../services/InsertTags.md), [DeleteRemovedTags](../services/DeleteRemovedTags.md), [GetTagsByArticleId](../services/GetTagsByArticleId.md)
