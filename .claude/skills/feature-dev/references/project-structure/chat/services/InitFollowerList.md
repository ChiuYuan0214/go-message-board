# InitFollowerList

**File:** `backend/chat/services/`

## Signature

```go
func InitFollowerList(wg *sync.WaitGroup, conn *websocket.Conn, userId uint64)
```

## Behaviour

Queries MySQL for the user's followers. Sets `client.FollowerList`.
Also sends a `userInfoList` event to the client listing which followers are currently online.
Called concurrently in [InitChatClient](InitChatClient.md).
