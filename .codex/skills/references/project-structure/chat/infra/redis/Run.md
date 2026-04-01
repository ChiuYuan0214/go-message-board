# Redis.Run

**File:** `backend/chat/infra/redis.go`
**Struct:** `Redis`

```go
func (r *Redis) Run() (err error)
```

- Initializes the Redis client for chat
- Verifies connectivity with `PING`
- Fails fast through the depin lifecycle if Redis is unavailable
- This is the Redis readiness gate for WebSocket auth checks
