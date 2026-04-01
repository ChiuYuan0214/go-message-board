# GetUserIdFromToken

**File:** `backend/security/services/` (auth service)

## Signature

```go
func GetUserIdFromToken(srcToken string) uint64
```

## Returns

Parsed `userId` from JWT claims. `0` if invalid.

## When to reuse

Used by auth middleware to extract the caller's identity from the `Authorization` header.
