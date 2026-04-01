package service

import (
	"general/types"
	"time"
)

type Article interface {
	GetArticle(userId uint64, articleId string) (article *types.Article, code int)
	InsertArticle(userId uint64, article *types.AddArticleData, publishTime *time.Time) uint64
	UpdateArticle(userId uint64, articleId uint64, data *types.UpdateArticleData) (string, int)
	DeleteArticle(userId uint64, articleId uint64) (string, int)
	GetTagsByArticleId(articleId string) (tags []string)
	InsertTags(articleId uint64, tags []string) bool
}

type Articles interface {
	GetNewestList(page, size int64, userId uint64) (list []types.ArticleListData)
	GetViewList(page, size int64, userId uint64) (data []types.ArticleListData)
	GetHotList(page, size int64, userId uint64) (data []types.ArticleListData)
	GetProfileList(page, size int64, userId uint64, selfUserId uint64) (data []types.ArticleListData)
	GetTagList(page, size int64, tag string) (data []types.ArticleListData)
	setTags(data []types.ArticleListData, idList []string)
}

type Collection interface {
	GetCollections(userId uint64, page, size int64) (data []types.CollectionData)
	AddCollection(userId, articleId uint64) bool
	RemoveCollection(userId, articleId uint64) bool
}

type Comment interface {
	AddComment(userId uint64, data *types.AddCommentData) uint64
	UpdateComment(userId uint64, data *types.UpdateCommentData) (string, int)
	DeleteComment(userId uint64, commentIdStr string) (string, int)
}

type Comments interface {
	GetComments(articleId uint64) (data []types.CommentListData)
}

type Follow interface {
	AddFollow(userId uint64, followee uint64) bool
	RemoveFollow(userId uint64, followee uint64) bool
}

type Follower interface {
	RemoveFollower(userId uint64, follower uint64) bool
	GetFollowers(userId uint64) []types.Follower
}

type Follows interface {
	GetFollows(userId uint64) (data []types.Follower)
}

type Profile interface {
	GetProfileWithId(userId uint64) (*types.Profile, int)
	GetProfileWithToken(userId uint64) (*types.SelfProfile, int)
}

type View interface {
	RecordView(articleId string)
}

type Vote interface {
	Vote(userId, sourceId uint64, score int8, voteType *string) (string, uint64)
	UpdateVote(userId, voteId uint64, score int8) bool
}
