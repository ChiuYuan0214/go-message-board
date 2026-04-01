# ArticlesImpl.GetTagList

**File:** `backend/general/service/articles.go`
**Struct:** `ArticlesImpl`

```go
func (s *ArticlesImpl) GetTagList(page, size int64, tag string) (data []types.ArticleListData)
```

- Service method on `ArticlesImpl` coordinating `articles` business logic in general service.
