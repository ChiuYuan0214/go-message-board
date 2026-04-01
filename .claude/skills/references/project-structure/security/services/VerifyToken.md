# VerifyToken

**File:** `backend/security/services/` (auth service)

## Signature

```go
func VerifyToken(userId uint64, token string) bool
```

## Behaviour

Compares the provided token against the value stored in Redis for the user.
Returns `false` if expired, wrong, or not in Redis.
