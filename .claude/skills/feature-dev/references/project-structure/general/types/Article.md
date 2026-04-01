# Article

**Purpose:** Response — full single-article detail view
**File:** `backend/general/types/article.go`

## Fields

| Field | JSON | GORM | Notes |
|-------|------|------|-------|
| `ArticleId` | `articleId` | `primaryKey` | |
| `UserId` | `userId` | `column:user_id` | |
| `Author` | `author` | `-` | Joined from `users.username` |
| `AuthorImage` | `authorImage` | `-` | Joined from `images.file_name` |
| `Title` | `title` | `column:title` | |
| `Content` | `content` | `column:content` | |
| `TopCommentId` | `topCommentId` | `column:top_comment_id` | |
| `Edited` | `edited` | `column:edited` | |
| `ViewCount` | `viewCount` | `column:view_count` | |
| `VoteUp` | `voteUp` | `-` | Aggregated via subquery |
| `VoteDown` | `voteDown` | `-` | Aggregated via subquery |
| `MyScore` | `myScore` | `-` | Caller's current vote (`1`/`-1`/`0`) |
| `HasCollec` | `hasCollec` | `-` | Whether caller saved this article |
| `PublishTime` | `publishTime` | `column:publish_time` | |
| `CreationTime` | `creationTime` | `column:creation_time` | |
| `UpdateTime` | `updateTime` | `column:update_time` | |
| `Tags` | `tags` | `-` | Populated separately after query |

## Used in

- [GetArticle](../services/GetArticle.md) — returns this type
- Route: [GET /article](../routes/article.md)
