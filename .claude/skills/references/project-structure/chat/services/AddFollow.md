# AddFollow (Chat)

**File:** `backend/chat/services/`

## Signature

```go
func AddFollow(event *types.RequestEvent)
```

## Behaviour

Adds `event.TargetUserId` to the client's in-memory `FollowList`.
If the target is online, notifies them that this user has followed them.

> Note: This only updates the in-memory state. The actual DB write happens via the general service's [AddFollow](../../general/services/AddFollow.md).
