# SelfProfile

**Purpose:** Response — authenticated user's own profile (superset of Profile)
**File:** `backend/general/types/profile.go`

## Fields

All fields from [Profile](Profile.md), plus:

| Field | JSON | Notes |
|-------|------|-------|
| `Email` | `email` | Only visible to self |
| `Phone` | `phone` | Only visible to self |
| `Address` | `address` | Only visible to self |
| `CreationTime` | `creationTime` | Account creation date |

## Used in

- [GetProfileWithToken](../services/GetProfileWithToken.md)
- Route: [GET /profile](../routes/profile.md) (authenticated)
