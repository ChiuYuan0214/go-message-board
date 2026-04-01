# Redis.SAdd

**File:** `backend/general/infra/redis.go`
**Struct:** `Redis`

```go
func (i *Redis) SAdd(key string, val string) error
```

- Redis helper used by general repos/services for auth tokens and ranking/cache state.
