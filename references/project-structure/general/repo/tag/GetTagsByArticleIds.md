# TagImpl.GetTagsByArticleIds

**File:** `backend/general/repo/tag.go`
**Struct:** `TagImpl`

```go
func (r *TagImpl) GetTagsByArticleIds(articleIds []string) (data []dto.ArticleTag, err error)
```

- Repo method on `TagImpl` for the `tag` persistence path in general service.
