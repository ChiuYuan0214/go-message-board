# Routes: /updatePassword

**File:** `backend/security/` (handler)

## Endpoints

| Method | Auth | Body | Response |
|--------|------|------|----------|
| PUT | Yes | `{oldPassword, newPassword}` | `{status}` |

## Notes

Calls [VerifyPasswordByUserId](../services/VerifyPasswordByUserId.md) first. Only updates if old password is correct.
