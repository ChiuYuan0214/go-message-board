# RedisCache

**Category:** legacy helper type
**File:** `backend/chat/types/redis-cache.go`

## Fields

| Field | Type | Notes |
|-------|------|-------|
| `Client` | `*redis.Client` | Redis client |
| `Ctx` | `context.Context` | Request context for Redis commands |

## Methods

| Method | Notes |
|--------|-------|
| `GetToken(userId)` | Reads the current token string from Redis |

## Notes

- Kept mainly for source discovery
- Current depin flow prefers `infra.Redis` plus `repo.TokenImpl`
