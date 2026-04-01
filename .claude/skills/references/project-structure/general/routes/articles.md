# Routes: /articles

**File:** `backend/general/routes/articles.go`
**Handler:** `ArticlesHandler`

## Endpoints

| Method | Auth | Params | Response |
|--------|------|--------|----------|
| GET | No | `?type=newest\|view\|hot\|profile\|tag&page=&size=&userId=&tag=` | `{status, list: ArticleListData[]}` |

## Notes

- `type=newest` → [GetNewestList](../services/GetNewestList.md)
- `type=view` → [GetViewList](../services/GetViewList.md)
- `type=hot` → [GetHotList](../services/GetHotList.md)
- `type=profile` → [GetProfileList](../services/GetProfileList.md) (requires `?userId=`)
- `type=tag` → [GetTagList](../services/GetTagList.md) (requires `?tag=`)
- Tags are fetched and attached per article after list query
