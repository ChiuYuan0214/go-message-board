# GetProfileWithId

**File:** `backend/general/services/profile.go`

## Signature

```go
func GetProfileWithId(userId uint64) (*types.Profile, int)
```

## Returns

| Value | Notes |
|-------|-------|
| `*types.Profile` | Public profile; nil on error |
| `int` | `0` = success, HTTP status on error |

## Behaviour

Returns public profile fields: username, job, isActive, imagePath, and aggregated counts (articles, comments, upvotes received).
Does not return email, phone, or address.

## When to reuse

Any endpoint that needs to display another user's public profile info.
