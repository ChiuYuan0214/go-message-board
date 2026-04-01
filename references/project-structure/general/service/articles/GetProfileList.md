# ArticlesImpl.GetProfileList

**File:** `backend/general/service/articles.go`
**Struct:** `ArticlesImpl`

```go
func (s *ArticlesImpl) GetProfileList(page, size int64, userId uint64, selfUserId uint64) (data []types.ArticleListData)
```

- Service method on `ArticlesImpl` coordinating `articles` business logic in general service.
