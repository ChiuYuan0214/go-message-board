# FollowImpl.GetFollowerIDs

**File:** `backend/chat/repo/follow.go`
**Struct:** `FollowImpl`

```go
func (r *FollowImpl) GetFollowerIDs(userId uint64) ([]uint64, error)
```

- Queries MySQL `followers` for `follower_id` by `user_id`
- Used during chat connect to build the caller's in-memory `FollowerList`
