# CommentImpl.GetByArticleId

**File:** `backend/general/repo/comment.go`
**Struct:** `CommentImpl`

```go
func (r *CommentImpl) GetByArticleId(articleId uint64) (data []types.CommentListData, err error)
```

- Repo method on `CommentImpl` for the `comment` persistence path in general service.
