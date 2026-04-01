# UpdatePassword

**File:** `backend/security/services/` (user service)

## Signature

```go
func UpdatePassword(userId *uint64, password *string) bool
```

## Returns

`true` on success.

## Behaviour

Hashes the new password then updates `users.password` for the given user.
