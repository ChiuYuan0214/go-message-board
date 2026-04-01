# AddArticleData

**Purpose:** Request — create article
**File:** `backend/general/types/request.go`

## Fields

| Field | JSON | Required | Notes |
|-------|------|----------|-------|
| `Title` | `title` | Yes | |
| `Content` | `content` | Yes | |
| `PublishTime` | `publishTime` | Yes | Format: `"2006-01-02T15:04"` |
| `Tags` | `tags` | No | Array of tag name strings |

## Used in

- [InsertArticle](../services/InsertArticle.md)
- Route: [POST /article](../routes/article.md)
