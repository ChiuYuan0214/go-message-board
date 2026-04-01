# Project Notes

## Local Verification

- Prefer `docker-compose.yml` for cross-service integration checks.
- For `backend/chat`, use `docker-compose.yml` first when you need Redis, MySQL, and the chat server together for WebSocket verification.
- If you fall back to manual local runs, note env overrides explicitly and clean up temporary processes after testing.

## Reference Services

- `backend/general` is the primary reference for the depin layered pattern.
- `backend/chat` follows the same depin shape, but WebSocket lifecycle concerns still need separate service responsibilities.
