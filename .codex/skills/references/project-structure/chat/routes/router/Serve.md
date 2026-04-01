# RouterImpl.Serve

**File:** `backend/chat/routes/router.go`
**Struct:** `RouterImpl`

```go
func (r *RouterImpl) Serve()
```

- Starts the HTTP server on `constants.PORT`
- Wraps the default mux with `UseCORS()`
- Final serving step after all depin-managed handlers have run their `Run()` hooks
