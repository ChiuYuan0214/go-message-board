# Redis.Del

**File:** `backend/general/infra/redis.go`
**Struct:** `Redis`

```go
func (i *Redis) Del(key string) error
```

- Redis helper used by general repos/services for auth tokens and ranking/cache state.
