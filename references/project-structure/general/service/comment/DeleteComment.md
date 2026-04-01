# CommentImpl.DeleteComment

**File:** `backend/general/service/comment.go`
**Struct:** `CommentImpl`

```go
func (s *CommentImpl) DeleteComment(userId uint64, commentIdStr string) (string, int)
```

- Service method on `CommentImpl` coordinating `comment` business logic in general service.
