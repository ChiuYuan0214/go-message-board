package main

import (
	"general/constants"
	"general/infra"
	"general/jobs"
	"general/repo"
	"general/routes"
	"general/service"

	"github.com/ChiuYuan0214/depin"
)

func main() {
	constants.InitEnv()

	jobs.UseScheduler()

	router := depin.RunAndSet[routes.Router](new(routes.RouterImpl))

	depin.Set[infra.RDB](new(infra.MySQL))
	depin.Set[infra.Cache](new(infra.Redis))

	depin.Set[repo.Article](new(repo.ArticleImpl))
	depin.Set[repo.Collection](new(repo.CollectionImpl))
	depin.Set[repo.Comment](new(repo.CommentImpl))
	depin.Set[repo.Follower](new(repo.FollowerImpl))
	depin.Set[repo.Profile](new(repo.ProfileImpl))
	depin.Set[repo.Tag](new(repo.TagImpl))
	depin.Set[repo.TagMap](new(repo.TagMapImpl))
	depin.Set[repo.Vote](new(repo.VoteImpl))

	depin.Set[service.Article](new(service.ArticleImpl))
	depin.Set[service.Articles](new(service.ArticlesImpl))
	depin.Set[service.Collection](new(service.CollectionImpl))
	depin.Set[service.Comment](new(service.CommentImpl))
	depin.Set[service.Comments](new(service.CommentsImpl))
	depin.Set[service.Follow](new(service.FollowImpl))
	depin.Set[service.Follower](new(service.FollowerImpl))
	depin.Set[service.Follows](new(service.FollowsImpl))
	depin.Set[service.Profile](new(service.ProfileImpl))
	depin.Set[service.View](new(service.ViewImpl))
	depin.Set[service.Vote](new(service.VoteImpl))

	depin.Set[routes.Handler](new(routes.ArticleHandler))
	depin.Set[routes.Handler](new(routes.ArticlesHandler))
	depin.Set[routes.Handler](new(routes.CollectionsHandler))
	depin.Set[routes.Handler](new(routes.CommentHandler))
	depin.Set[routes.Handler](new(routes.CommentsHandler))
	depin.Set[routes.Handler](new(routes.FollowHandler))
	depin.Set[routes.Handler](new(routes.FollowerHandler))
	depin.Set[routes.Handler](new(routes.FollowsHandler))
	depin.Set[routes.Handler](new(routes.ProfileHandler))
	depin.Set[routes.Handler](new(routes.ViewHandler))
	depin.Set[routes.Handler](new(routes.VoteHandler))

	router.Serve()
}
