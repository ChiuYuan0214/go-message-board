# FollowImpl.RemoveFollow

**File:** `backend/chat/services/follow.go`
**Struct:** `FollowImpl`

```go
func (s *FollowImpl) RemoveFollow(event *types.RequestEvent)
```

- Removes `TargetUserId` from the caller's in-memory `FollowList`
- In-memory only; assumes the real follow relationship has already been changed elsewhere
