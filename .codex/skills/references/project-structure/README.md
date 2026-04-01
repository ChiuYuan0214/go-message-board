# Codebase Reference Index

Folder-based reference for backend services. Use this to locate ownership, reusable methods, and the current layering shape before opening source files.

Use this file as the entry point for feature planning:
- identify the owning service first
- check which layers already exist (`infra`, `repo`, `service`, `routes`, `types`, `entities`)
- open the smallest relevant reference file before reading source code
- treat method names and short descriptions as discovery hints only
- before reusing an existing method, confirm the real source code still matches the intended business logic
- if schema, args, API, or layering changes, update this index and the affected leaf reference files in the same task

---

## General Service — port 8080 (Gin + GORM · MySQL + Redis)

### Layer Map

- `infra`: MySQL (`infra.RDB`), Redis (`infra.Cache`)
- `repo`: article, vote, comment, tag, tag-map, collection, follower, profile
- `service`: article, articles, collection, comment, comments, follow, follower, follows, profile, view, vote
- `routes`: article, articles, collections, comment, comments, follow, follower, follows, profile, view, vote

### Infra and Repo Index

- `infra`: `infra/mysql.go`, `infra/redis.go`, `infra/interface.go`
- `repo`: `repo/article.go`, `repo/collection.go`, `repo/comment.go`, `repo/follower.go`, `repo/profile.go`, `repo/tag.go`, `repo/tag_map.go`, `repo/vote.go`, `repo/interface.go`

### Entities
| File | Table |
|------|-------|
| [article-tag-map](general/entities/article-tag-map.md) | `article_tag_maps` |
| [collection](general/entities/collection.md) | `collections` |
| [comment](general/entities/comment.md) | `comments` |
| [follower](general/entities/follower.md) | `followers` |
| [tag](general/entities/tag.md) | `tags` |
| [vote](general/entities/vote.md) | `votes` |

### Types
| File | Purpose |
|------|---------|
| [Article](general/types/Article.md) | Response — full article detail |
| [ArticleListData](general/types/ArticleListData.md) | Response — article list item |
| [AddArticleData](general/types/AddArticleData.md) | Request — create article |
| [UpdateArticleData](general/types/UpdateArticleData.md) | Request — update article |
| [CollectionData](general/types/CollectionData.md) | Response — collection list item |
| [WriteCollectionData](general/types/WriteCollectionData.md) | Request — add/remove collection |
| [CommentListData](general/types/CommentListData.md) | Response — comment list item |
| [AddCommentData](general/types/AddCommentData.md) | Request — create comment |
| [UpdateCommentData](general/types/UpdateCommentData.md) | Request — update comment |
| [Profile](general/types/Profile.md) | Response — public profile |
| [SelfProfile](general/types/SelfProfile.md) | Response — self profile (auth'd) |
| [Follower](general/types/Follower.md) | Response — follower/follow user |
| [FollowData](general/types/FollowData.md) | Request — follow/unfollow |
| [FollowerData](general/types/FollowerData.md) | Request — remove follower |
| [User](general/types/User.md) | DB user record |
| [Token](general/types/Token.md) | Auth token |

### Services
| File | Function |
|------|----------|
| [GetArticle](general/services/GetArticle.md) | Fetch full article with vote counts and collection status |
| [InsertArticle](general/services/InsertArticle.md) | Create article record |
| [UpdateArticle](general/services/UpdateArticle.md) | Update article; validates ownership |
| [DeleteArticle](general/services/DeleteArticle.md) | Delete article + cascade |
| [GetTagsByArticleId](general/services/GetTagsByArticleId.md) | Get tag names for an article |
| [InsertTags](general/services/InsertTags.md) | Create tags and article_tag_maps |
| [DeleteRemovedTags](general/services/DeleteRemovedTags.md) | Remove stale tag mappings |
| [GetNewestList](general/services/GetNewestList.md) | Articles sorted by publish time |
| [GetViewList](general/services/GetViewList.md) | Articles by view count (Redis) |
| [GetHotList](general/services/GetHotList.md) | Hot articles (Redis) |
| [GetProfileList](general/services/GetProfileList.md) | Articles by a specific user |
| [GetTagList](general/services/GetTagList.md) | Articles filtered by tag |
| [GetCollections](general/services/GetCollections.md) | User's saved collections |
| [AddCollection](general/services/AddCollection.md) | Save article to collection |
| [RemoveCollection](general/services/RemoveCollection.md) | Remove article from collection |
| [AddComment](general/services/AddComment.md) | Create comment on article |
| [UpdateComment](general/services/UpdateComment.md) | Update comment; validates ownership |
| [DeleteComment](general/services/DeleteComment.md) | Delete comment + cascade votes |
| [GetComments](general/services/GetComments.md) | All comments for an article |
| [AddFollow](general/services/AddFollow.md) | Follow a user |
| [RemoveFollow](general/services/RemoveFollow.md) | Unfollow a user |
| [RemoveFollower](general/services/RemoveFollower.md) | Remove a follower |
| [GetFollowers](general/services/GetFollowers.md) | Users following this user |
| [GetFollows](general/services/GetFollows.md) | Users this user follows |
| [GetProfileWithId](general/services/GetProfileWithId.md) | Public profile by user ID |
| [GetProfileWithToken](general/services/GetProfileWithToken.md) | Self profile (authenticated) |
| [Vote](general/services/Vote.md) | Create or update a vote |
| [UpdateVote](general/services/UpdateVote.md) | Toggle/update existing vote |
| [RecordView](general/services/RecordView.md) | Increment article view in Redis |

### Routes
| File | Path |
|------|------|
| [article](general/routes/article.md) | `GET/POST/PUT/DELETE /article` |
| [articles](general/routes/articles.md) | `GET /articles` |
| [collections](general/routes/collections.md) | `GET/POST/DELETE /collections` |
| [comment](general/routes/comment.md) | `POST/PUT/DELETE /comment` |
| [comments](general/routes/comments.md) | `GET /comments` |
| [follow](general/routes/follow.md) | `POST/DELETE /follow` |
| [follower](general/routes/follower.md) | `GET/DELETE /follower` |
| [profile](general/routes/profile.md) | `GET /profile` |
| [vote](general/routes/vote.md) | `POST/PUT /vote` |
| [view](general/routes/view.md) | `PUT /view` |

---

## Security Service — port 7080 (net/http · MySQL + Redis)

### Layer Map

- `infra`: MySQL (`infra.RDB`), Redis (`infra.Cache`)
- `repo`: auth, register, profile, user
- `store`: in-memory active user directory
- `jobs`: user directory sync from MySQL
- `services`: auth, register, profile, user directory
- `routes`: register, verify-code, login, update-password, update-profile, upload-image, users

### Wiring Index

- `infra`: `infra/db.go`, `infra/cache.go`, `infra/interface.go`
- `repo`: `repo/auth.go`, `repo/register.go`, `repo/profile.go`, `repo/user.go`, `repo/interface.go`
- `store`: `store/users.go`
- `jobs`: `jobs/init-users.go`
- `services`: `services/login.go`, `services/register.go`, `services/profile.go`, `services/common.go`, `services/interface.go`
- `routes`: `routes/base.go`, `routes/interface.go`, `routes/login.go`, `routes/register.go`, `routes/resend-veri-code.go`, `routes/verfiy-code.go`, `routes/update-password.go`, `routes/update-profile.go`, `routes/upload-image.go`, `routes/users.go`, `routes/interceptor.go`

### Wiring Notes

- `main.go` now wires concrete infra -> repo -> services -> jobs -> routes instances directly instead of using package-level `UsePool` / `UseCache` globals.
- `repo` owns MySQL / Redis access for auth, register, profile, and user-directory reads.
- `store.UsersStore` owns the in-memory user list used by `/users`.
- `jobs.UsersSyncJob` refreshes the in-memory user directory through `repo.User`.
- Security is not on depin yet, but the runtime dependencies are now explicit and injected.

### Services
| File | Function |
|------|----------|
| [CheckEmailExist](security/services/CheckEmailExist.md) | Check if email already registered |
| [AddNewUser](security/services/AddNewUser.md) | Insert new user; hashes password |
| [InsertVerificationCode](security/services/InsertVerificationCode.md) | Insert verification code record |
| [InvalidateVerificationCodes](security/services/InvalidateVerificationCodes.md) | Mark all user codes invalid |
| [ActivateUser](security/services/ActivateUser.md) | Set user is_active = true |
| [VerifyPasswordByEmail](security/services/VerifyPasswordByEmail.md) | Verify email + password |
| [Login](security/services/Login.md) | Authenticate and return token |
| [GenerateToken](security/services/GenerateToken.md) | Create JWT; store in Redis |
| [VerifyToken](security/services/VerifyToken.md) | Validate token against Redis |
| [GetUserIdFromToken](security/services/GetUserIdFromToken.md) | Parse JWT → userId |
| [UpdateColumnsById](security/services/UpdateColumnsById.md) | Dynamic user column update |
| [VerifyPasswordByUserId](security/services/VerifyPasswordByUserId.md) | Verify password by user ID |
| [UpdatePassword](security/services/UpdatePassword.md) | Hash and update password |
| [InsertProfileImageInfo](security/services/InsertProfileImageInfo.md) | Upsert profile image record |

### Routes
| File | Path |
|------|------|
| [register](security/routes/register.md) | `POST /register` |
| [verify-code](security/routes/verify-code.md) | `POST /verifyCode`, `POST /resendVerificationCode` |
| [login](security/routes/login.md) | `POST/PUT /login` |
| [update-password](security/routes/update-password.md) | `PUT /updatePassword` |
| [update-profile](security/routes/update-profile.md) | `POST /updateProfile`, `POST /uploadImage` |
| [users](security/routes/users.md) | `GET /users` |

---

## Chat Service — port 9080 (Gorilla WebSocket · MySQL + DynamoDB + Redis)

### Layer Map

- `infra`: MySQL (`infra.RDB`), Redis (`infra.Cache`), DynamoDB (`infra.Dynamo`)
- `repo`: token, follow, history
- `service`: chat, token, event, history, message, follow-list, follow, notify
- `jobs`: scheduler
- `routes`: chat handler

### Infra and Repo Index

- `infra`: `infra/mysql.go`, `infra/redis.go`, `infra/dynamo.go`, `infra/interface.go`
- `repo`: `repo/token.go`, `repo/follow.go`, `repo/history.go`, `repo/interface.go`

### Infra Responsibilities
| Impl | Purpose |
|------|---------|
| `infra.MySQL` · [Run](chat/infra/mysql/Run.md) / [DB](chat/infra/mysql/DB.md) | MySQL connection lifecycle and raw DB handle access |
| `infra.Redis` · [Run](chat/infra/redis/Run.md) / [GetToken](chat/infra/redis/GetToken.md) | Redis client lifecycle and token lookup |
| `infra.DynamoDB` · [Run](chat/infra/dynamo_db/Run.md) / [Client](chat/infra/dynamo_db/Client.md) | DynamoDB client lifecycle and typed client access |

### Key Types
| File | Purpose |
|------|---------|
| [Client](chat/types/Client.md) | WebSocket connection and state |
| [Message](chat/types/Message.md) | Chat message format |
| [SendMap](chat/types/SendMap.md) | Per-client message cache |
| [RequestEvent](chat/types/RequestEvent.md) | Client request format |
| [History](chat/types/History.md) | Message history response |
| [Notification](chat/types/Notification.md) | Follow/follower online status event |
| [UserInfoList](chat/types/UserInfoList.md) | Online follow/follower list payload |
| [ServerMessage](chat/types/ServerMessage.md) | Server-side error / control payload |
| [DynamoChat](chat/types/DynamoChat.md) | DynamoDB persistence record for chat history |
| [DynamoClient](chat/types/DynamoClient.md) | Typed DynamoDB wrapper used by history repo |
| [RedisCache](chat/types/RedisCache.md) | Legacy Redis token helper kept for old type coverage |

### Repo Responsibilities
| Impl | Purpose |
|------|---------|
| `repo.TokenImpl` · [GetToken](chat/repo/token/GetToken.md) | Token lookup via Redis |
| `repo.FollowImpl` · [GetFollowerIDs](chat/repo/follow/GetFollowerIDs.md) / [GetFollowIDs](chat/repo/follow/GetFollowIDs.md) | Follow/follower lookup via MySQL |
| `repo.HistoryImpl` · [GetHistory](chat/repo/history/GetHistory.md) / [GetHistoryLimit20](chat/repo/history/GetHistoryLimit20.md) / [BatchInsert](chat/repo/history/BatchInsert.md) | History read/write via DynamoDB |

### Services
| Impl / Reference | Responsibility |
|------|----------|
| `services.ChatImpl` · [InitChatClient](chat/services/chat/InitChatClient.md) / [ListenChatEvent](chat/services/chat/ListenChatEvent.md) | WebSocket client lifecycle and read loop |
| `services.TokenImpl` · [ValidateToken](chat/services/token/ValidateToken.md) / [UseTokenChecker](chat/services/token/UseTokenChecker.md) | Token validation and periodic token check |
| `services.EventImpl` · [RunEventLoop](chat/services/event/RunEventLoop.md) / [HandleEvent](chat/services/event/HandleEvent.md) | Broadcast event loop and event dispatch |
| `services.MessageImpl` · [SendMessage](chat/services/message/SendMessage.md) | Send message to target user |
| `services.HistoryImpl` · [GetHistory](chat/services/history/GetHistory.md) | History fetch and cache merge |
| `services.FollowListImpl` · [InitFollowerList](chat/services/follow_list/InitFollowerList.md) / [InitFollowList](chat/services/follow_list/InitFollowList.md) | Load follow graphs on connect |
| `services.FollowImpl` · [AddFollow](chat/services/follow/AddFollow.md) / [RemoveFollow](chat/services/follow/RemoveFollow.md) / [RemoveFollower](chat/services/follow/RemoveFollower.md) | Update in-memory follow state |
| `services.NotifyImpl` · [NotifyLogin](chat/services/notify/NotifyLogin.md) / [NotifyLogout](chat/services/notify/NotifyLogout.md) | Broadcast login/logout notifications |

### Routes and State

| Impl / Reference | Purpose |
|------|---------|
| `routes.ChatHandler` · [Run](chat/routes/chat_handler/Run.md) / [handleChats](chat/routes/chat_handler/handleChats.md) | WebSocket route registration and connection bootstrap |
| `routes.RouterImpl` · [Handle](chat/routes/router/Handle.md) / [Serve](chat/routes/router/Serve.md) | `net/http` route registration and server startup |
| [ChatStore](chat/store/ChatStore.md) | Shared in-memory chat state |

### Wiring Notes

- `main.go` registers `infra -> repo -> service -> jobs -> routes`, then calls `depin.Run()`.
- `routes.ChatHandler` currently depends on `services.Chat`, `services.Token`, and `services.Event`.
- `docker-compose.yml` is the preferred integration-test path when validating chat with Redis/MySQL together.

---

## Stream Service — port 5000 (Gin + WebSocket · in-memory)

### Layer Map

- `store`: in-memory stream state
- `service`: live, owner, watcher, record handlers
- `routes`: socket and HLS handlers

### State and Service Index

- `store`: `store/stream.go`
- `services`: `services/live.go`, `services/live-record.go`, `services/owner.go`, `services/owner-record.go`, `services/watcher.go`, `services/watcher-record.go`, `services/rtmp.go`
- `routes`: `routes/socket.go`, `routes/hls.go`, `routes/base.go`, `routes/utils.go`

### Key Types
| File | Purpose |
|------|---------|
| [Live](stream/types/Live.md) | Live stream session |
| [Client](stream/types/Client.md) | Base WebSocket client |
| [Owner-Watcher](stream/types/Owner-Watcher.md) | Broadcaster / viewer |

### Services
| File | Function |
|------|----------|
| [OwnerService-Handle](stream/services/OwnerService-Handle.md) | Read from owner; push to watchers |
| [WatcherService-Handle](stream/services/WatcherService-Handle.md) | Passive receive loop |
| [OwnerRecordService-Handle](stream/services/OwnerRecordService-Handle.md) | Owner text message handling |
| [WatcherRecordService-Handle](stream/services/WatcherRecordService-Handle.md) | Watcher text message handling |
| [LiveService-PushStream](stream/services/LiveService-PushStream.md) | Broadcast binary stream |

---

## Reference Backlog

These source areas exist in the repo but do not yet have enough leaf references for fast planning. Prefer adding them when touching the related area:

- `general`: `infra/*`, `repo/*`, `service/*`, shared route plumbing such as `routes/router.go`, `routes/interface.go`, `routes/utils.go`, `routes/follows.go`
- `security`: `infra/*`, `store/users.go`, shared route helpers/interceptors, and service layer files that are not yet indexed by structure
- `stream`: `store/stream.go`, route files, and most service/type files still need structure-level references
