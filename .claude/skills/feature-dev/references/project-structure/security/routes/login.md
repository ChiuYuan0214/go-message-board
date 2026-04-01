# Routes: /login

**File:** `backend/security/` (handler)

## Endpoints

| Method | Auth | Body | Response |
|--------|------|------|----------|
| POST | No | `{email, password}` | `{status, token, expireTime, userId}` |
| PUT | No | `{userId, token}` | `{status, token, expireTime}` |

## Notes

- POST → [Login](../services/Login.md) — authenticates and issues new token
- PUT → refreshes an existing token via [GenerateToken](../services/GenerateToken.md) after [VerifyToken](../services/VerifyToken.md)
