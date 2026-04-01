# Client (Stream)

**Purpose:** Base WebSocket client for stream service
**File:** `backend/stream/types/`

## Fields / Methods

| Member | Notes |
|--------|-------|
| `userId` | |
| `liveConn` | Binary WebSocket (stream data) |
| `recordConn` | Text WebSocket (chat/reactions) |
| `live` | Reference to the Live session |
| `PushStream(msg)` | Write binary frame to liveConn |
| `ReadStream()` | Read binary frame from liveConn |
| `Write(msg)` | Write text message to recordConn |
| `Read()` | Read text message from recordConn |
| `Close()` | Close both connections |

## Notes

`Owner` and `Watcher` both embed Client. `OwnerRecord` and `WatcherRecord` also embed Client.
