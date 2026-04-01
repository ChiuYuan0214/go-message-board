# Stream Service Reference

`backend/stream` is the live-stream / socket service on port `5000`.

## Layer Map

- `store`: in-memory stream state
- `services`: live, owner, watcher, record handlers
- `routes`: socket and HLS handlers
- `types`: live/client/owner-watcher runtime shapes

## Shared State and Types

- [Live](types/Live.md)
- [Client](types/Client.md)
- [Owner-Watcher](types/Owner-Watcher.md)

## Services

- [OwnerService-Handle](services/OwnerService-Handle.md)
- [WatcherService-Handle](services/WatcherService-Handle.md)
- [OwnerRecordService-Handle](services/OwnerRecordService-Handle.md)
- [WatcherRecordService-Handle](services/WatcherRecordService-Handle.md)
- [LiveService-PushStream](services/LiveService-PushStream.md)

## Notes

- Stream docs are still lighter than the other three services.
- Use these files for discovery, then confirm the real socket/store flow in source before reusing behavior.
