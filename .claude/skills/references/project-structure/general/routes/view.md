# Routes: /view

**File:** `backend/general/routes/view.go`
**Handler:** `ViewHandler`

## Endpoints

| Method | Auth | Params | Response |
|--------|------|--------|----------|
| PUT | No | `?articleId=` | `{status}` |

## Notes

Calls [RecordView](../services/RecordView.md). Fire-and-forget; always returns success.
