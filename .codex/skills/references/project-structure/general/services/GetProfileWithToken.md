# GetProfileWithToken

**File:** `backend/general/services/profile.go`

## Signature

```go
func GetProfileWithToken(userId uint64) (*types.SelfProfile, int)
```

## Returns

| Value | Notes |
|-------|-------|
| `*types.SelfProfile` | Full self-profile; nil on error |
| `int` | `0` = success, HTTP status on error |

## Behaviour

Same as [GetProfileWithId](GetProfileWithId.md) but includes private fields: email, phone, address, creationTime.
Only call this for the authenticated user's own profile.
