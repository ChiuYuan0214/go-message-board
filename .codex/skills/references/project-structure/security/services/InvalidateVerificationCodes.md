# InvalidateVerificationCodes

**File:** `backend/security/services/` (verification service)

## Signature

```go
func InvalidateVerificationCodes(userId int64) bool
```

## Behaviour

Marks **all** verification codes for the user as invalid (`valid = false`).
Call before issuing a new code to prevent code reuse.
