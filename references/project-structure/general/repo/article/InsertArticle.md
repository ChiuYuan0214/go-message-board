# ArticleImpl.InsertArticle

**File:** `backend/general/repo/article.go`
**Struct:** `ArticleImpl`

```go
func (r *ArticleImpl) InsertArticle(userId uint64, article *types.AddArticleData, publishTime *time.Time) (id uint64, err error)
```

- Repo method on `ArticleImpl` for the `article` persistence path in general service.
