# AddNewUser

**File:** `backend/security/services/` (user service)

## Signature

```go
func AddNewUser(username, email, password, phone, job, address string) int64
```

## Returns

New user ID. `0` on error.

## Behaviour

Hashes the password before inserting. Returns the new `user_id`.
