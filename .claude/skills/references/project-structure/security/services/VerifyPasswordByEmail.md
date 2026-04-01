# VerifyPasswordByEmail

**File:** `backend/security/services/` (auth service)

## Signature

```go
func VerifyPasswordByEmail(email, password *string) int64
```

## Returns

| Value | Meaning |
|-------|---------|
| `userId > 0` | Success — returns user ID |
| `0` | Wrong email or password |
| `-1` | Account not activated |

## When to reuse

Core credential check used by [Login](Login.md).
