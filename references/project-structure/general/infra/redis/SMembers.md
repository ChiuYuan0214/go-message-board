# Redis.SMembers

**File:** `backend/general/infra/redis.go`
**Struct:** `Redis`

```go
func (i *Redis) SMembers(key string) []string
```

- Redis helper used by general repos/services for auth tokens and ranking/cache state.
