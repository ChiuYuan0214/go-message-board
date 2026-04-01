# User

**Purpose:** Internal — DB user record
**File:** `backend/general/types/user.go`

## Fields

| Field | JSON | GORM | Notes |
|-------|------|------|-------|
| `UserId` | `userId` | `primaryKey` | |
| `Username` | `username` | `column:username` | |
| `Email` | `email` | `column:email` | |
| `Password` | `password` | `column:password` | Hashed |
| `CreationTime` | `creationTime` | `column:creation_time` | |

## Notes

Mostly used as a join target in general service queries. Auth and profile management live in the security service.
