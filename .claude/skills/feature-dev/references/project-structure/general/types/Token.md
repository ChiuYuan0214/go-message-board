# Token

**Purpose:** Auth token response
**File:** `backend/general/types/token.go`

## Fields

| Field | JSON | Notes |
|-------|------|-------|
| `Token` | `token` | JWT string |
| `ExpireTime` | `expireTime` | Token expiry timestamp |

## Notes

JWT is HS256, 30-minute expiry. Generated and cached in Redis by the security service.
See [GenerateToken](../../security/services/GenerateToken.md) and [VerifyToken](../../security/services/VerifyToken.md).
