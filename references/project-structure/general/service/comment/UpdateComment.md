# CommentImpl.UpdateComment

**File:** `backend/general/service/comment.go`
**Struct:** `CommentImpl`

```go
func (s *CommentImpl) UpdateComment(userId uint64, data *types.UpdateCommentData) (string, int)
```

- Service method on `CommentImpl` coordinating `comment` business logic in general service.
