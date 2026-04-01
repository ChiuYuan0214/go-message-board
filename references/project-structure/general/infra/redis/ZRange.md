# Redis.ZRange

**File:** `backend/general/infra/redis.go`
**Struct:** `Redis`

```go
func (i *Redis) ZRange(key string, page, pageSize int64) []string
```

- Redis helper used by general repos/services for auth tokens and ranking/cache state.
