# Codebase Reference Index

Folder-based reference for all backend services. Each service function has its own file.
Use this to find reusable methods before writing new ones.

---

## General Service — port 8080 (Gin + GORM · MySQL + Redis)

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

## Chat Service — port 9080 (Gorilla WebSocket · MySQL + MongoDB + Redis)

### Key Types
| File | Purpose |
|------|---------|
| [Client](chat/types/Client.md) | WebSocket connection and state |
| [Message](chat/types/Message.md) | Chat message format |
| [SendMap](chat/types/SendMap.md) | Per-client message cache |
| [RequestEvent](chat/types/RequestEvent.md) | Client request format |
| [History](chat/types/History.md) | Message history response |

### Services
| File | Function |
|------|----------|
| [InitChatClient](chat/services/InitChatClient.md) | Create/update client on connect |
| [ListenChatEvent](chat/services/ListenChatEvent.md) | Main WebSocket read loop |
| [UseTokenChecker](chat/services/UseTokenChecker.md) | Periodic token validation |
| [SendMessage](chat/services/SendMessage.md) | Send message to target user |
| [GetHistory](chat/services/GetHistory.md) | Fetch message history |
| [InitFollowerList](chat/services/InitFollowerList.md) | Load followers on connect |
| [InitFollowList](chat/services/InitFollowList.md) | Load follows on connect |
| [AddFollow](chat/services/AddFollow.md) | Add to follow list |
| [RemoveFollow](chat/services/RemoveFollow.md) | Remove from follow list |
| [RemoveFollower](chat/services/RemoveFollower.md) | Remove from follower list |
| [NotifyLogin](chat/services/NotifyLogin.md) | Broadcast login to followers |
| [NotifyLogout](chat/services/NotifyLogout.md) | Broadcast logout to followers |

---

## Stream Service — port 5000 (Gin + WebSocket · in-memory)

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
