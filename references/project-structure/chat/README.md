# Chat Service Reference

`backend/chat` is the WebSocket chat service on port `9080`.

## Layer Map

- `infra`: MySQL (`infra.RDB`), Redis (`infra.Cache`), DynamoDB (`infra.Dynamo`)
- `repo`: token, follow, history
- `services`: chat, token, event, history, message, follow-list, follow, notify
- `jobs`: scheduler
- `routes`: chat handler, router
- `store`: shared chat state
- `types`: websocket/runtime/history payloads

## Infra Responsibilities

- `infra.MySQL`: [Run](infra/mysql/Run.md), [DB](infra/mysql/DB.md)
- `infra.Redis`: [Run](infra/redis/Run.md), [GetToken](infra/redis/GetToken.md)
- `infra.DynamoDB`: [Run](infra/dynamo_db/Run.md), [Client](infra/dynamo_db/Client.md)

## Repo Responsibilities

- `repo.TokenImpl`: [GetToken](repo/token/GetToken.md)
- `repo.FollowImpl`: [GetFollowerIDs](repo/follow/GetFollowerIDs.md), [GetFollowIDs](repo/follow/GetFollowIDs.md)
- `repo.HistoryImpl`: [GetHistory](repo/history/GetHistory.md), [GetHistoryLimit20](repo/history/GetHistoryLimit20.md), [BatchInsert](repo/history/BatchInsert.md)

## Services

- `services.ChatImpl`: [InitChatClient](services/chat/InitChatClient.md), [ListenChatEvent](services/chat/ListenChatEvent.md)
- `services.TokenImpl`: [ValidateToken](services/token/ValidateToken.md), [UseTokenChecker](services/token/UseTokenChecker.md)
- `services.EventImpl`: [RunEventLoop](services/event/RunEventLoop.md), [HandleEvent](services/event/HandleEvent.md)
- `services.MessageImpl`: [SendMessage](services/message/SendMessage.md)
- `services.HistoryImpl`: [GetHistory](services/history/GetHistory.md)
- `services.FollowListImpl`: [InitFollowerList](services/follow_list/InitFollowerList.md), [InitFollowList](services/follow_list/InitFollowList.md)
- `services.FollowImpl`: [AddFollow](services/follow/AddFollow.md), [RemoveFollow](services/follow/RemoveFollow.md), [RemoveFollower](services/follow/RemoveFollower.md)
- `services.NotifyImpl`: [NotifyLogin](services/notify/NotifyLogin.md), [NotifyLogout](services/notify/NotifyLogout.md)

## Routes and State

- `routes.ChatHandler`: [Run](routes/chat_handler/Run.md), [handleChats](routes/chat_handler/handleChats.md)
- `routes.RouterImpl`: [Handle](routes/router/Handle.md), [Serve](routes/router/Serve.md)
- [ChatStore](store/ChatStore.md)

## Key Types

- [Client](types/Client.md), [Message](types/Message.md), [SendMap](types/SendMap.md), [RequestEvent](types/RequestEvent.md)
- [History](types/History.md), [Notification](types/Notification.md), [UserInfoList](types/UserInfoList.md), [ServerMessage](types/ServerMessage.md)
- [DynamoChat](types/DynamoChat.md), [DynamoClient](types/DynamoClient.md), [RedisCache](types/RedisCache.md)

## Wiring Notes

- `main.go` registers `infra -> repo -> services -> jobs -> routes`, then calls `depin.Run()`.
- `routes.ChatHandler` depends on `services.Chat`, `services.Token`, and `services.Event`.
- Prefer `docker-compose.yml` when validating chat with Redis/MySQL together.
