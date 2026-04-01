package repo

import (
	"general/dto"
	"general/entities"
	"general/enum"
	"general/types"
	"time"
)

type Article interface {
	InsertArticle(userId uint64, article *types.AddArticleData, publishTime *time.Time) (uint64, error)
	UpdateArticle(articleId uint64, data map[string]interface{}) (err error)
	DeleteArticle(articleId uint64) error
	GetArticleById(articleId uint64) (types.Article, error)
	GetArticleDetail(userId uint64, articleId string) (types.Article, error)
	GetNewsList(userId uint64, start, size int64) (list []types.ArticleListData, err error)
	GetViewList(userId uint64, articleIds []string) (list []types.ArticleListData, err error)
	GetHotList(userId uint64, articleIds []string) (list []types.ArticleListData, err error)
	GetProfileList(userId, selfUserId uint64, start, size int64) (list []types.ArticleListData, err error)
	GetTagList(tag string) (list []types.ArticleListData, err error)
}

type Vote interface {
	GetByUserAndSource(userId, sourceId uint64) (entities.Vote, error)
	CreateVote(userId, sourceId uint64, score int8, voteType string) (uint64, error)
	GetById(voteId uint64) (entities.Vote, error)
	UpdateScore(voteId uint64, score int8) error
	DeleteVoteBySourceId(source enum.VoteSource, sourceId uint64) (err error)
}

type Comment interface {
	Create(userId, articleId uint64, title, content string) (newComment *entities.Comment, err error)
	Updates(commentId uint64, fields map[string]any) (err error)
	DeleteById(commentId uint64) (err error)
	DeleteByArticleId(articleId uint64) (err error)
	GetById(commentId uint64) (comment entities.Comment, err error)
	GetByArticleId(articleId uint64) (data []types.CommentListData, err error)
}

type Tag interface {
	Create(name string) (err error)
	GetTagsByArticleId(articleId string) (tags []string, err error)
	GetByNames(names []string) (tags []entities.Tag, err error)
	GetTagsByArticleIds(articleIds []string) (data []dto.ArticleTag, err error)
}

type TagMap interface {
	DeleteByArticleId(articleId uint64) (err error)
	CreateIgnoringConflict(articleId uint64, tagId uint64) (err error)
}

type Collection interface {
	Create(userId, articleId uint64) (err error)
	DeleteByUserIdAndArticleId(userId, articleId uint64) (err error)
	DeleteByArticleId(articleId uint64) (err error)
	GetByUserId(userId uint64, start, size int64) (data []types.CollectionData, err error)
}

type Follower interface {
	Create(userId, followerId uint64) (err error)
	GetByUserIdAndFollowerId(userId, followerId uint64) (follower entities.Follower, err error)
	GetFollowersByUserId(userId uint64) (followers []types.Follower, err error)
	GetUsersByFollowerId(followerId uint64) (users []types.Follower, err error)
	DeleteByUserIdAndFollowerId(userId, followerId uint64) (err error)
}

type Profile interface {
	GetByUserId(userId uint64) (profile types.Profile, err error)
	GetSelfById(userId uint64) (profile types.SelfProfile, err error)
}
