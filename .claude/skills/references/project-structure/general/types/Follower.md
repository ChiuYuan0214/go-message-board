# Follower

**Purpose:** Response — user in a follower/follow list
**File:** `backend/general/types/follower.go`

## Fields

| Field | JSON | Notes |
|-------|------|-------|
| `UserId` | `userId` | |
| `Username` | `username` | |
| `UserImage` | `userImage` | Profile image filename |

## Used in

- [GetFollowers](../services/GetFollowers.md), [GetFollows](../services/GetFollows.md)
- Route: [GET /follower](../routes/follower.md)
