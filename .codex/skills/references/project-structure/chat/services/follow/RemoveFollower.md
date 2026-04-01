# FollowImpl.RemoveFollower

**File:** `backend/chat/services/follow.go`
**Struct:** `FollowImpl`

```go
func (s *FollowImpl) RemoveFollower(event *types.RequestEvent)
```

- Removes `TargetUserId` from the caller's in-memory `FollowerList`
- In-memory only; used to keep chat presence state aligned with follow changes
