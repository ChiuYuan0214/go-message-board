# Project Notes

## Local Verification

- Prefer `docker-compose.yml` for integration checks that span multiple services.
- For `backend/chat`, `docker-compose.yml` is the default path when you need Redis, MySQL, and the chat server together for WebSocket verification.
- For `backend/chat`, prefer `make lint` for repeatable Go static analysis, then `make build` / `make test` for narrower checks.
- `backend/chat` lint also includes local readability checks for struct/slice initialization style; keep the "blank line between unrelated code blocks" rule in manual review because it is not enforced reliably by a simple script.
- If you fall back to manual local runs, write down env overrides and clean up temporary processes after testing.

## Reference Services

- `backend/general` is the primary reference for the depin layered pattern.
- `backend/chat` follows the same depin shape, but WebSocket lifecycle concerns should stay split across small services.
