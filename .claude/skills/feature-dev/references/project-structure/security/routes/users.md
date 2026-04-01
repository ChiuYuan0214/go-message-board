# Routes: /users

**File:** `backend/security/` (handler)

## Endpoints

| Method | Auth | Params | Response |
|--------|------|--------|----------|
| GET | Yes | `?name=` | `{status, list: [{userId, username}]}` |

## Notes

Searches users by username (LIKE). Excludes the authenticated caller from results.
Also serves uploaded images at `GET /uploads/images/{filename}` (static file server, no auth).
