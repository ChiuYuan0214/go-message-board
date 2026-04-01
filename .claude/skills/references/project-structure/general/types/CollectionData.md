# CollectionData

**Purpose:** Response — collection list item
**File:** `backend/general/types/collection.go`

## Fields

| Field | JSON | Notes |
|-------|------|-------|
| `ArticleId` | `articleId` | |
| `UserId` | `userId` | Article author |
| `Title` | `title` | |
| `Content` | `content` | |
| `Author` | `author` | |
| `AuthorImage` | `authorImage` | |
| `VoteUp` | `voteUp` | |
| `VoteDown` | `voteDown` | |
| `MyScore` | `myScore` | |
| `HasCollec` | `hasCollec` | Always true in collection context |
| `PublishTime` | `publishTime` | |

## Used in

- [GetCollections](../services/GetCollections.md)
- Route: [GET /collections](../routes/collections.md)
