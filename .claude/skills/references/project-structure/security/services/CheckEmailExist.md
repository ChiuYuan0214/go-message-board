# CheckEmailExist

**File:** `backend/security/services/` (user service)

## Signature

```go
func CheckEmailExist(email string) bool
```

## Returns

`true` if email is already registered.

## When to reuse

Call before [AddNewUser](AddNewUser.md) to prevent duplicate registration.
