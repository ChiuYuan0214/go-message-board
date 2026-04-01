# Routes: /profile

**File:** `backend/general/routes/profile.go`
**Handler:** `ProfileHandler`

## Endpoints

| Method | Auth | Params | Response |
|--------|------|--------|----------|
| GET | Conditional | `?userId=` (unauthenticated) | `{status, data: Profile \| SelfProfile}` |

## Notes

- If a valid token is present → calls [GetProfileWithToken](../services/GetProfileWithToken.md) → returns [SelfProfile](../types/SelfProfile.md)
- If no token → requires `?userId=` → calls [GetProfileWithId](../services/GetProfileWithId.md) → returns [Profile](../types/Profile.md)
