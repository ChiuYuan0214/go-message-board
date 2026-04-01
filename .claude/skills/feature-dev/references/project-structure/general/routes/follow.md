# Routes: /follow

**File:** `backend/general/routes/follow.go`
**Handler:** `FollowHandler`

## Endpoints

| Method | Auth | Body | Response |
|--------|------|------|----------|
| POST | Yes | `{followee: uint64}` | `{status}` |
| DELETE | Yes | `{followee: uint64}` | `{status}` |
