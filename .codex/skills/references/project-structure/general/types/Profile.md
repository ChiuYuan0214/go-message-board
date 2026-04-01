# Profile

**Purpose:** Response — public profile
**File:** `backend/general/types/profile.go`

## Fields

| Field | JSON | Notes |
|-------|------|-------|
| `UserId` | `userId` | |
| `Username` | `username` | |
| `Job` | `job` | |
| `IsActive` | `isActive` | |
| `ImagePath` | `imagePath` | Profile image filename |
| `ArticleCount` | `articleCount` | Aggregated |
| `CommentCount` | `commentCount` | Aggregated |
| `UpVoteCount` | `upVoteCount` | Total upvotes received |

## Used in

- [GetProfileWithId](../services/GetProfileWithId.md)
- Route: [GET /profile](../routes/profile.md) (unauthenticated)
