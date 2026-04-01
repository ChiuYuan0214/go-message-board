# ArticleImpl.GetArticleDetail

**File:** `backend/general/repo/article.go`
**Struct:** `ArticleImpl`

```go
func (r *ArticleImpl) GetArticleDetail(userId uint64, articleId string) (article types.Article, err error)
```

- Repo method on `ArticleImpl` for the `article` persistence path in general service.
