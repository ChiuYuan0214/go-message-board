# ChatHandler.Run

**File:** `backend/chat/routes/chat.go`
**Struct:** `ChatHandler`

```go
func (h *ChatHandler) Run() (err error)
```

- Configures the WebSocket upgrader
- Registers `/chat`
- Starts the injected event loop
- Requires `depin.Run()` to have executed, otherwise route registration never happens
