# ArticleImpl.GetNewsList

**File:** `backend/general/repo/article.go`
**Struct:** `ArticleImpl`

```go
func (r *ArticleImpl) GetNewsList(userId uint64, start, size int64) (list []types.ArticleListData, err error)
```

- Repo method on `ArticleImpl` for the `article` persistence path in general service.
