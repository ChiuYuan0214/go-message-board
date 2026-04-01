# InsertProfileImageInfo

**File:** `backend/security/services/` (user service)

## Signature

```go
func InsertProfileImageInfo(userId *uint64, fileName *string, desc *string) (string, int)
```

## Returns

`("", 0)` on success. `(errorMessage, httpStatus)` on failure.

## Behaviour

Upserts a row in `images` table (`user_id`, `file_name`, `descript`).
Called after saving the uploaded image file to disk.
