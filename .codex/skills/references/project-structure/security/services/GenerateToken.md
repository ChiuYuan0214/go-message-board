# GenerateToken

**File:** `backend/security/services/` (auth service)

## Signature

```go
func GenerateToken(userId uint64) *types.Token
```

## Returns

`*types.Token` with JWT string and expiry. nil on error.

## Behaviour

Creates a HS256 JWT with `userId` claim and 30-minute expiry. Stores the token in Redis for validation.
The general service middleware reads from Redis to authenticate requests.
