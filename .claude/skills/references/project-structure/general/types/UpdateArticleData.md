# UpdateArticleData

**Purpose:** Request — update article
**File:** `backend/general/types/request.go`

## Fields

| Field | JSON | Notes |
|-------|------|-------|
| `Title` | `title` | Optional — omit to keep existing |
| `Content` | `content` | Optional — omit to keep existing |
| `Tags` | `tags` | Full replacement list |

## Used in

- [UpdateArticle](../services/UpdateArticle.md)
- Route: [PUT /article](../routes/article.md)
