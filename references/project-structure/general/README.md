# General Service Reference

`backend/general` is the article/forum HTTP service on port `8080`.

## Layer Map

- `infra`: MySQL (`infra.RDB`), Redis (`infra.Cache`)
- `repo`: article, vote, comment, tag, tag-map, collection, follower, profile
- `service`: article, articles, collection, comment, comments, follow, follower, follows, profile, view, vote
- `routes`: article, articles, collections, comment, comments, follow, follower, follows, profile, view, vote
- `entities`: article-tag-map, collection, comment, follower, tag, vote
- `types`: article/article-list/profile/comment/follow request-response types

## Infra Responsibilities

- `infra.MySQL`: [Run](infra/mysql/Run.md), [Orm](infra/mysql/Orm.md)
- `infra.Redis`: [Run](infra/redis/Run.md), [GetToken](infra/redis/GetToken.md), [HMSet](infra/redis/HMSet.md), [ZAdd](infra/redis/ZAdd.md)

## Repo Responsibilities

- `repo.ArticleImpl`: [GetArticleDetail](repo/article/GetArticleDetail.md), [GetNewsList](repo/article/GetNewsList.md), [InsertArticle](repo/article/InsertArticle.md)
- `repo.CollectionImpl`: [Create](repo/collection/Create.md), [GetByUserId](repo/collection/GetByUserId.md)
- `repo.CommentImpl`: [Create](repo/comment/Create.md), [GetByArticleId](repo/comment/GetByArticleId.md)
- `repo.FollowerImpl`: [Create](repo/follower/Create.md), [GetFollowersByUserId](repo/follower/GetFollowersByUserId.md), [GetUsersByFollowerId](repo/follower/GetUsersByFollowerId.md)
- `repo.ProfileImpl`: [GetByUserId](repo/profile/GetByUserId.md), [GetSelfById](repo/profile/GetSelfById.md)
- `repo.TagImpl`: [GetByNames](repo/tag/GetByNames.md), [GetTagsByArticleIds](repo/tag/GetTagsByArticleIds.md)
- `repo.TagMapImpl`: [CreateIgnoringConflict](repo/tag_map/CreateIgnoringConflict.md), [DeleteByArticleId](repo/tag_map/DeleteByArticleId.md)
- `repo.VoteImpl`: [GetByUserAndSource](repo/vote/GetByUserAndSource.md), [CreateVote](repo/vote/CreateVote.md), [UpdateScore](repo/vote/UpdateScore.md)

## Services

- `service.ArticleImpl`: [GetArticle](service/article/GetArticle.md), [InsertArticle](service/article/InsertArticle.md), [UpdateArticle](service/article/UpdateArticle.md), [DeleteArticle](service/article/DeleteArticle.md)
- `service.ArticlesImpl`: [GetNewestList](service/articles/GetNewestList.md), [GetViewList](service/articles/GetViewList.md), [GetHotList](service/articles/GetHotList.md), [setTags](service/articles/setTags.md)
- `service.CollectionImpl`: [GetCollections](service/collection/GetCollections.md), [AddCollection](service/collection/AddCollection.md), [RemoveCollection](service/collection/RemoveCollection.md)
- `service.CommentImpl`: [AddComment](service/comment/AddComment.md), [UpdateComment](service/comment/UpdateComment.md), [DeleteComment](service/comment/DeleteComment.md)
- `service.CommentsImpl`: [GetComments](service/comments/GetComments.md)
- `service.FollowImpl`: [AddFollow](service/follow/AddFollow.md), [RemoveFollow](service/follow/RemoveFollow.md)
- `service.FollowerImpl`: [GetFollowers](service/follower/GetFollowers.md), [RemoveFollower](service/follower/RemoveFollower.md)
- `service.FollowsImpl`: [GetFollows](service/follows/GetFollows.md)
- `service.ProfileImpl`: [GetProfileWithId](service/profile/GetProfileWithId.md), [GetProfileWithToken](service/profile/GetProfileWithToken.md)
- `service.ViewImpl`: [RecordView](service/view/RecordView.md)
- `service.VoteImpl`: [Vote](service/vote/Vote.md), [UpdateVote](service/vote/UpdateVote.md)

## Routes

- `routes.RouterImpl`: [Run](routes/router/Run.md), [Get](routes/router/Get.md), [Post](routes/router/Post.md), [Serve](routes/router/Serve.md)
- `routes.ArticleHandler`: [Run](routes/article_handler/Run.md), [get](routes/article_handler/get.md), [add](routes/article_handler/add.md), [update](routes/article_handler/update.md), [delete](routes/article_handler/delete.md)
- `routes.ArticlesHandler`: [Run](routes/articles_handler/Run.md), [get](routes/articles_handler/get.md)
- `routes.CollectionsHandler`: [Run](routes/collections_handler/Run.md), [get](routes/collections_handler/get.md), [add](routes/collections_handler/add.md), [delete](routes/collections_handler/delete.md)
- `routes.CommentHandler`: [Run](routes/comment_handler/Run.md), [add](routes/comment_handler/add.md), [update](routes/comment_handler/update.md), [delete](routes/comment_handler/delete.md)
- `routes.CommentsHandler`: [Run](routes/comments_handler/Run.md), [get](routes/comments_handler/get.md)
- `routes.FollowHandler`: [Run](routes/follow_handler/Run.md), [add](routes/follow_handler/add.md), [remove](routes/follow_handler/remove.md)
- `routes.FollowerHandler`: [Run](routes/follower_handler/Run.md), [get](routes/follower_handler/get.md), [remove](routes/follower_handler/remove.md)
- `routes.FollowsHandler`: [Run](routes/follows_handler/Run.md), [get](routes/follows_handler/get.md)
- `routes.ProfileHandler`: [Run](routes/profile_handler/Run.md), [get](routes/profile_handler/get.md)
- `routes.ViewHandler`: [Run](routes/view_handler/Run.md), [record](routes/view_handler/record.md)
- `routes.VoteHandler`: [Run](routes/vote_handler/Run.md), [add](routes/vote_handler/add.md), [update](routes/vote_handler/update.md)

## Wiring Notes

- `main.go` registers `infra -> repo -> service -> routes` through depin, then calls `depin.Run()`.
- General uses `service/` (singular) in source; docs keep the real source layout.
- Redis infra is shared by auth-token handling and article ranking/view caches.
