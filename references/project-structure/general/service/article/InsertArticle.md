# ArticleImpl.InsertArticle

**File:** `backend/general/service/article.go`
**Struct:** `ArticleImpl`

```go
func (s *ArticleImpl) InsertArticle(userId uint64, article *types.AddArticleData, publishTime *time.Time) uint64
```

- Service method on `ArticleImpl` coordinating `article` business logic in general service.
