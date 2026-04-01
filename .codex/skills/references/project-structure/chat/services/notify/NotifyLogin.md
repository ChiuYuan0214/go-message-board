# NotifyImpl.NotifyLogin

**File:** `backend/chat/services/notify.go`
**Struct:** `NotifyImpl`

```go
func (s *NotifyImpl) NotifyLogin(userId uint64)
```

- Broadcasts login notifications to online followers and follows
- Uses only in-memory chat state
- Emits `follow-login` to followers and `follower-login` to follows
