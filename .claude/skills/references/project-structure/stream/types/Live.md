# Live

**Purpose:** Represents an active live stream session
**File:** `backend/stream/types/`

## Fields / Methods

| Member | Type | Notes |
|--------|------|-------|
| `liveId` | `string` | Session identifier |
| `owner` | `*Owner` | The broadcaster |
| `watchers` | `WatcherMap` | `sync.Map` of userId → Watcher |
| `watcherCount` | `int` | |
| `isStart` | `bool` | Whether stream is active |
| `GetLiveId()` | method | |
| `GetOwner()` | method | |
| `WatcherJoin(w)` | method | Add watcher to map |
| `WatcherExit(id)` | method | Remove watcher from map |
| `GetWatchers()` | method | Returns WatcherMap |
| `GetWatcher(id)` | method | Lookup single watcher |
| `HasWatcher(id)` | method | |

## Notes

All sessions are in-memory — no DB persistence. There is no multi-room support currently (single hardcoded session).
