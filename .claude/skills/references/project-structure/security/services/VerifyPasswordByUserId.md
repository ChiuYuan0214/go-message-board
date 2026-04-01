# VerifyPasswordByUserId

**File:** `backend/security/services/` (auth service)

## Signature

```go
func VerifyPasswordByUserId(userId *uint64, password *string) bool
```

## Returns

`true` if password matches. Used when the caller is already authenticated (e.g. before changing password).
