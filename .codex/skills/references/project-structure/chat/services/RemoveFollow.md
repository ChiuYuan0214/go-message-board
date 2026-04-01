# RemoveFollow (Chat)

**File:** `backend/chat/services/`

## Signature

```go
func RemoveFollow(event *types.RequestEvent)
```

## Behaviour

Removes `event.TargetUserId` from the client's in-memory `FollowList`. In-memory only.
