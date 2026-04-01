# EventImpl.HandleEvent

**File:** `backend/chat/services/event.go`
**Struct:** `EventImpl`

```go
func (s *EventImpl) HandleEvent(event *types.RequestEvent)
```

- Routes `message`, `history`, `add-follow`, `remove-follow`, and `remove-follower`
- Delegates to injected `Message`, `History`, and `Follow` services
- `refresh-token` is currently a no-op
- This is the central place to extend chat event types
