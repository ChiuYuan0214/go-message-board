# CommentImpl.Create

**File:** `backend/general/repo/comment.go`
**Struct:** `CommentImpl`

```go
func (r *CommentImpl) Create(userId, articleId uint64, title, content string) (newComment *entities.Comment, err error)
```

- Repo method on `CommentImpl` for the `comment` persistence path in general service.
