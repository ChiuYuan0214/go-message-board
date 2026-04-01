# Routes: /article

**File:** `backend/general/routes/article.go`
**Handler:** `ArticleHandler`

## Endpoints

| Method | Auth | Params / Body | Response |
|--------|------|---------------|----------|
| GET | No | `?articleId=` | `{status, data: Article}` |
| POST | Yes | Body: [AddArticleData](../types/AddArticleData.md) + `publishTime` string | `{status, id}` |
| PUT | Yes | `?articleId=` + Body: [UpdateArticleData](../types/UpdateArticleData.md) | `{status}` |
| DELETE | Yes | `?articleId=` | `{status}` |

## Notes

- GET also calls [GetTagsByArticleId](../services/GetTagsByArticleId.md) and attaches tags to response
- POST parses `publishTime` as `"2006-01-02T15:04"` before calling [InsertArticle](../services/InsertArticle.md)
- PUT calls [DeleteRemovedTags](../services/DeleteRemovedTags.md) + [InsertTags](../services/InsertTags.md) after article update
- DELETE cascades via [DeleteArticle](../services/DeleteArticle.md)
