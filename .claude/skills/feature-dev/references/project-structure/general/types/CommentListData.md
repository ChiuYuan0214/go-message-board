# CommentListData

**Purpose:** Response — comment item in list
**File:** `backend/general/types/comment.go`

## Fields

| Field | JSON | Notes |
|-------|------|-------|
| `CommentId` | `commentId` | |
| `UserId` | `userId` | Commenter's user ID |
| `Commenter` | `commenter` | `users.username` |
| `CommenterImage` | `commenterImage` | `images.file_name` |
| `Title` | `title` | |
| `Content` | `content` | |
| `CreationTime` | `creationTime` | |
| `VoteUp` | `voteUp` | Aggregated |
| `VoteDown` | `voteDown` | Aggregated |

## Used in

- [GetComments](../services/GetComments.md)
- Route: [GET /comments](../routes/comments.md)
