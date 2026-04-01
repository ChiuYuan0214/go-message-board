# Codebase Reference Index

Use this file only to choose the owning backend, then jump into that service's `README.md`.

## Backend Services

- [general](general/README.md) — article/forum HTTP API on port `8080`
- [security](security/README.md) — auth/profile HTTP API on port `7080`
- [chat](chat/README.md) — WebSocket chat server on port `9080`
- [stream](stream/README.md) — live stream / socket server on port `5000`

## Usage

- Treat method names and short descriptions as discovery hints only.
- After finding a candidate method through the docs, confirm the real source code before reusing it.
- If schema, args, API, or layering changes, update the owning service README and the affected leaf docs in the same task.
