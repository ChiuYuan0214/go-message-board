# UpdateCommentData

**Purpose:** Request — update comment
**File:** `backend/general/types/request.go`

## Fields

| Field | JSON | Notes |
|-------|------|-------|
| `CommentId` | `commentId` | Target comment |
| `Title` | `title` | Optional |
| `Content` | `content` | Optional |

## Used in

- [UpdateComment](../services/UpdateComment.md)
- Route: [PUT /comment](../routes/comment.md)
