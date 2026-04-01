# InsertVerificationCode

**File:** `backend/security/services/` (verification service)

## Signature

```go
func InsertVerificationCode(userId int64, code int32, expireTime time.Time) int64
```

## Returns

New `code_id`. `0` on error.

## Behaviour

Inserts a verification code record into `verification_codes`. Used after [AddNewUser](AddNewUser.md) during registration flow.
