# ArticleImpl.GetArticle

**File:** `backend/general/service/article.go`
**Struct:** `ArticleImpl`

```go
func (s *ArticleImpl) GetArticle(userId uint64, articleId string) (article *types.Article, code int)
```

- Service method on `ArticleImpl` coordinating `article` business logic in general service.
