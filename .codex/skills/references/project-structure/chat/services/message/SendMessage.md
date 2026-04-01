# MessageImpl.SendMessage

**File:** `backend/chat/services/message.go`
**Struct:** `MessageImpl`

```go
func (s *MessageImpl) SendMessage(reqMsg *types.RequestEvent)
```

- Builds a `message` event from the request
- Stores it in the sender's `SendMap`
- Pushes to the target connection immediately if the target is online
- Does not persist immediately; Dynamo sync is deferred to the scheduler path
