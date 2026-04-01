# InitFollowList

**File:** `backend/chat/services/`

## Signature

```go
func InitFollowList(wg *sync.WaitGroup, conn *websocket.Conn, userId uint64)
```

## Behaviour

Same as [InitFollowerList](InitFollowerList.md) but for the users this client follows.
Sets `client.FollowList` and sends online status of follows to the client.
