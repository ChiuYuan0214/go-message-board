# Redis.LRange

**File:** `backend/general/infra/redis.go`
**Struct:** `Redis`

```go
func (i *Redis) LRange(key string, page, size int64) []string
```

- Redis helper used by general repos/services for auth tokens and ranking/cache state.
