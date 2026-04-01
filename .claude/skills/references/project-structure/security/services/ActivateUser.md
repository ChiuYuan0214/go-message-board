# ActivateUser

**File:** `backend/security/services/` (user service)

## Signature

```go
func ActivateUser(userId uint64) bool
```

## Returns

`true` on success.

## Behaviour

Sets `users.is_active = true` for the given user. Called after successful email verification.
