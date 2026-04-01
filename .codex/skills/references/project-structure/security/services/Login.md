# Login

**File:** `backend/security/services/` (auth service)

## Signature

```go
func Login(email, password string) (uint64, *types.Token)
```

## Returns

| Value | Notes |
|-------|-------|
| `uint64` | User ID; `0` on failure |
| `*types.Token` | JWT token + expiry; nil on failure |

## Behaviour

Calls [VerifyPasswordByEmail](VerifyPasswordByEmail.md), then [GenerateToken](GenerateToken.md).
Returns `(0, nil)` if credentials are wrong or account is inactive.
