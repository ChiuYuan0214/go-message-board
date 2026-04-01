# ChatHandler.handleChats

**File:** `backend/chat/routes/chat.go`
**Struct:** `ChatHandler`

```go
func (h *ChatHandler) handleChats(w http.ResponseWriter, r *http.Request)
```

- Upgrades the incoming request to WebSocket
- Extracts the token from query params
- Delegates token validation to `services.Token`
- Starts the chat read loop and token checker on success
- This is the handshake boundary between HTTP and the long-lived chat session
