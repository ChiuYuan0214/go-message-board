# Routes: /updateProfile, /uploadImage

**File:** `backend/security/` (handler)

## Endpoints

| Method | Path | Auth | Notes |
|--------|------|------|-------|
| POST | `/updateProfile` | Yes | Fields: `username`, `phone`, `job`, `address`; calls [UpdateColumnsById](../services/UpdateColumnsById.md) |
| POST | `/uploadImage` | Yes | Multipart form; max 5MB; saved as `img_{userId}{ext}`; calls [InsertProfileImageInfo](../services/InsertProfileImageInfo.md) |
