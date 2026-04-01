# WriteCollectionData

**Purpose:** Request — add or remove a collection
**File:** `backend/general/types/collection.go`

## Fields

| Field | JSON | Notes |
|-------|------|-------|
| `ArticleId` | `articleId` | Target article |

## Used in

- [AddCollection](../services/AddCollection.md), [RemoveCollection](../services/RemoveCollection.md)
- Route: [POST/DELETE /collections](../routes/collections.md)
