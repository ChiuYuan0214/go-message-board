# Routes: /verifyCode, /resendVerificationCode

**File:** `backend/security/` (handler)

## Endpoints

| Method | Path | Auth | Body | Response |
|--------|------|------|------|----------|
| POST | `/verifyCode` | No | `{userId, code}` | `{status}` |
| POST | `/resendVerificationCode` | No | `{email, password}` | `{status}` |

## Notes

- `/verifyCode`: validates code, calls [ActivateUser](../services/ActivateUser.md), invalidates code
- `/resendVerificationCode`: re-verifies password, invalidates old codes, issues new code + sends email
