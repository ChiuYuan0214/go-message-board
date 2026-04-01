# RemoveFollower (Chat)

**File:** `backend/chat/services/`

## Signature

```go
func RemoveFollower(event *types.RequestEvent)
```

## Behaviour

Removes `event.TargetUserId` from the client's in-memory `FollowerList`. In-memory only.
