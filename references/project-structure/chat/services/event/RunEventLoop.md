# EventImpl.RunEventLoop

**File:** `backend/chat/services/event.go`
**Struct:** `EventImpl`

```go
func (s *EventImpl) RunEventLoop()
```

- Consumes events from `ChatStore.Broadcast`
- Dispatches each event to the relevant injected service
- Keeps event routing separate from connection lifecycle
- Started once from `routes.ChatHandler.Run()`
