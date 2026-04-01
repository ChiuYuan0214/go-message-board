# Routes: /register

**File:** `backend/security/` (handler)

## Endpoints

| Method | Auth | Body | Response |
|--------|------|------|----------|
| POST | No | `{username, email, password, phone, job, address}` | `{status}` |

## Notes

- Validates email uniqueness via [CheckEmailExist](../services/CheckEmailExist.md)
- Concurrently: creates user via [AddNewUser](../services/AddNewUser.md) + sends verification email
- Issues a verification code via [InsertVerificationCode](../services/InsertVerificationCode.md)
