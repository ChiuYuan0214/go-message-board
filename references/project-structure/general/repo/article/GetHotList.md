# ArticleImpl.GetHotList

**File:** `backend/general/repo/article.go`
**Struct:** `ArticleImpl`

```go
func (r *ArticleImpl) GetHotList(userId uint64, articleIds []string) (list []types.ArticleListData, err error)
```

- Repo method on `ArticleImpl` for the `article` persistence path in general service.
