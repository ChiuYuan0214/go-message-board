# Redis.GetToken

**File:** `backend/general/infra/redis.go`
**Struct:** `Redis`

```go
func (i *Redis) GetToken(userId int64) (string, error)
```

- Redis helper used by general repos/services for auth tokens and ranking/cache state.
