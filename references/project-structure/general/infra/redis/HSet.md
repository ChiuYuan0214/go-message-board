# Redis.HSet

**File:** `backend/general/infra/redis.go`
**Struct:** `Redis`

```go
func (i *Redis) HSet(key string, mapKey string, value string) error
```

- Redis helper used by general repos/services for auth tokens and ranking/cache state.
