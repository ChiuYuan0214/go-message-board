# ArticlesImpl.GetNewestList

**File:** `backend/general/service/articles.go`
**Struct:** `ArticlesImpl`

```go
func (s *ArticlesImpl) GetNewestList(page, size int64, userId uint64) (list []types.ArticleListData)
```

- Service method on `ArticlesImpl` coordinating `articles` business logic in general service.
