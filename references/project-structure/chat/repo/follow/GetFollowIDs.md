# FollowImpl.GetFollowIDs

**File:** `backend/chat/repo/follow.go`
**Struct:** `FollowImpl`

```go
func (r *FollowImpl) GetFollowIDs(userId uint64) ([]uint64, error)
```

- Queries MySQL `followers` for followed `user_id` by `follower_id`
- Used during chat connect to build the caller's in-memory `FollowList`
