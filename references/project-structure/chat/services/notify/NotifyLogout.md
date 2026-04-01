# NotifyImpl.NotifyLogout

**File:** `backend/chat/services/notify.go`
**Struct:** `NotifyImpl`

```go
func (s *NotifyImpl) NotifyLogout(userId uint64)
```

- Broadcasts logout notifications to online followers and follows
- Uses only in-memory chat state
- Emits `follow-logout` to followers and `follower-logout` to follows
