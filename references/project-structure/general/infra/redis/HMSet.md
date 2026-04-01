# Redis.HMSet

**File:** `backend/general/infra/redis.go`
**Struct:** `Redis`

```go
func (i *Redis) HMSet(key string, vcm *map[string]string) error
```

- Redis helper used by general repos/services for auth tokens and ranking/cache state.
