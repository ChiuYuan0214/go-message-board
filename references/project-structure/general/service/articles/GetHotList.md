# ArticlesImpl.GetHotList

**File:** `backend/general/service/articles.go`
**Struct:** `ArticlesImpl`

```go
func (s *ArticlesImpl) GetHotList(page, size int64, userId uint64) (data []types.ArticleListData)
```

- Service method on `ArticlesImpl` coordinating `articles` business logic in general service.
