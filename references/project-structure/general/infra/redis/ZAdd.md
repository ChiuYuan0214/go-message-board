# Redis.ZAdd

**File:** `backend/general/infra/redis.go`
**Struct:** `Redis`

```go
func (i *Redis) ZAdd(key string, list *([]string)) error
```

- Redis helper used by general repos/services for auth tokens and ranking/cache state.
