# Comment

**Table:** `comments`
**File:** `backend/general/entities/comment.go`

## Fields

| Field | Type | Column | Notes |
|-------|------|--------|-------|
| `CommentId` | `uint64` | `comment_id` | PK, auto-increment |
| `UserId` | `uint64` | `user_id` | FK → users |
| `ArticleId` | `uint64` | `article_id` | FK → articles |
| `Title` | `string` | `title` | |
| `Content` | `string` | `content` | |
| `Edited` | `bool` | `edited` | default false |
| `UpdateTime` | `time.Time` | `update_time` | auto on update |

## Related

- Services: [AddComment](../services/AddComment.md), [UpdateComment](../services/UpdateComment.md), [DeleteComment](../services/DeleteComment.md), [GetComments](../services/GetComments.md)
