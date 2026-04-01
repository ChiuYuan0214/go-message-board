# ArticleListData

**Purpose:** Response — article item in list views
**File:** `backend/general/types/article.go`

## Fields

| Field | JSON | Notes |
|-------|------|-------|
| `ArticleId` | `articleId` | |
| `UserId` | `userId` | |
| `Title` | `title` | |
| `Content` | `content` | |
| `Author` | `author` | `users.username` |
| `AuthorImage` | `authorImage` | `images.file_name` |
| `VoteUp` | `voteUp` | Aggregated |
| `VoteDown` | `voteDown` | Aggregated |
| `MyScore` | `myScore` | Caller's vote |
| `HasCollec` | `hasCollec` | Whether caller saved it |
| `CommentTitle` | `commentTitle` | Top comment title |
| `CommentContent` | `commentContent` | Top comment content |
| `CommentUser` | `commentUser` | Top comment author |
| `CommentUserImage` | `commentUserImage` | Top comment author image |
| `PublishTime` | `publishTime` | |
| `Tags` | `tags` | Populated separately |

## Used in

- [GetNewestList](../services/GetNewestList.md), [GetViewList](../services/GetViewList.md), [GetHotList](../services/GetHotList.md), [GetProfileList](../services/GetProfileList.md), [GetTagList](../services/GetTagList.md)
- Route: [GET /articles](../routes/articles.md)
