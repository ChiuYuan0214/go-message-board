# Routes: /comment

**File:** `backend/general/routes/comment.go`
**Handler:** `CommentHandler`

## Endpoints

| Method | Auth | Body | Response |
|--------|------|------|----------|
| POST | Yes | [AddCommentData](../types/AddCommentData.md) | `{status, id}` |
| PUT | Yes | [UpdateCommentData](../types/UpdateCommentData.md) | `{status}` |
| DELETE | Yes | `?commentId=` | `{status}` |
