# Redis.HGet

**File:** `backend/general/infra/redis.go`
**Struct:** `Redis`

```go
func (i *Redis) HGet(key1 string, key2 string) string
```

- Redis helper used by general repos/services for auth tokens and ranking/cache state.
