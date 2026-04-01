# AddCommentData

**Purpose:** Request — create comment
**File:** `backend/general/types/request.go`

## Fields

| Field | JSON | Required | Notes |
|-------|------|----------|-------|
| `ArticleId` | `articleId` | Yes | Target article |
| `Title` | `title` | Yes | |
| `Content` | `content` | Yes | |

## Used in

- [AddComment](../services/AddComment.md)
- Route: [POST /comment](../routes/comment.md)
