# FollowListImpl.InitFollowerList

**File:** `backend/chat/services/follow-list.go`
**Struct:** `FollowListImpl`

```go
func (s *FollowListImpl) InitFollowerList(wg *sync.WaitGroup, userId uint64)
```

- Loads follower ids via `repo.Follow`
- Stores them on the in-memory client
- Sends `online-follower-list` with the currently online subset
- Used only during connection bootstrap
