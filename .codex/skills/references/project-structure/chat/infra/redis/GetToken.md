# Redis.GetToken

**File:** `backend/chat/infra/redis.go`
**Struct:** `Redis`

```go
func (r *Redis) GetToken(userId uint64) (string, error)
```

- Reads the current auth token for a user from Redis
- Used by token-related repo and service code during WebSocket authentication
- Current chat repo usage only needs token lookup, not generic Redis operations
