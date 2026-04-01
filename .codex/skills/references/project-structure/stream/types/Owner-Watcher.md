# Owner / Watcher

**Purpose:** Broadcaster and viewer types for live stream
**File:** `backend/stream/types/`

## Owner

Inherits all fields/methods from [Client](Client.md). The broadcaster.
- Sends stream data via `liveConn`
- Sends/receives text events (chat, reactions, polls, feedback) via `recordConn`

## Watcher

Inherits all fields/methods from [Client](Client.md). A viewer.
- Receives stream data via `liveConn` (passive)
- Sends/receives text events via `recordConn`

## OwnerRecord / WatcherRecord

Separate connection objects for the recording/interaction channel (`recordConn`).
Used when the client connects to `/socket` separately from the main stream `/chat`.

## Current Limitation

Owner is identified by `userId == 123 && liveId == 1` (hardcoded). Single session only.
