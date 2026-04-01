# UpdateColumnsById

**File:** `backend/security/services/` (user service)

## Signature

```go
func UpdateColumnsById(data interface{}, id *uint64) (string, int)
```

## Returns

`("", 0)` on success. `(errorMessage, httpStatus)` on failure.

## Behaviour

Dynamic column update on the `users` table using the provided struct fields.
Used for profile updates (username, phone, job, address).
