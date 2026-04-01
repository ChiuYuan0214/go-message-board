# Collection

**Table:** `collections`
**File:** `backend/general/entities/collection.go`

## Fields

| Field | Type | Column | Notes |
|-------|------|--------|-------|
| `UserId` | `uint64` | `user_id` | Composite PK |
| `ArticleId` | `uint64` | `article_id` | Composite PK |

## Purpose

Records that a user has saved ("collected") an article. Many-to-many between users and articles.

## Related

- Services: [AddCollection](../services/AddCollection.md), [RemoveCollection](../services/RemoveCollection.md), [GetCollections](../services/GetCollections.md)
