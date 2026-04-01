# Redis.SetToken

**File:** `backend/general/infra/redis.go`
**Struct:** `Redis`

```go
func (i *Redis) SetToken(userId int64, token types.Token) error
```

- Redis helper used by general repos/services for auth tokens and ranking/cache state.
