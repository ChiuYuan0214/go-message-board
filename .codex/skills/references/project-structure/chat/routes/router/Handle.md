# RouterImpl.Handle

**File:** `backend/chat/routes/router.go`
**Struct:** `RouterImpl`

```go
func (r *RouterImpl) Handle(pattern string, handler http.HandlerFunc)
```

- Registers a handler on the default `net/http` mux
- Chat route setup currently uses the default mux, not a custom router library
