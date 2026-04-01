# ArticleImpl.UpdateArticle

**File:** `backend/general/service/article.go`
**Struct:** `ArticleImpl`

```go
func (s *ArticleImpl) UpdateArticle(userId uint64, articleId uint64, data *types.UpdateArticleData) (string, int)
```

- Service method on `ArticleImpl` coordinating `article` business logic in general service.
