# Routes: /comments

**File:** `backend/general/routes/comments.go`
**Handler:** `CommentsHandler`

## Endpoints

| Method | Auth | Params | Response |
|--------|------|--------|----------|
| GET | No | `?articleId=` | `{status, list: CommentListData[]}` |
