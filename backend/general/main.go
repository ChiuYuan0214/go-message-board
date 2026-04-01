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

	// db := setup.InitGorm()
	// routes.UsePool(db)
	// service.UsePool(db)
	// jobs.UsePool(db)

	// cache := setup.InitCache()
	// defer cache.Client.Close()
	// service.UseCache(cache)
	// jobs.UseCache(cache)

	jobs.UseScheduler()

	router := depin.RunAndSet[routes.Router](new(routes.RouterImpl))

	depin.Set[infra.RDB](new(infra.MySQL))
	depin.Set[infra.Cache](new(infra.Redis))

	depin.Set[repo.Article](new(repo.ArticleImpl))
	depin.Set[repo.Collection](new(repo.CollectionImpl))
	depin.Set[repo.Comment](new(repo.CommentImpl))
	depin.Set[repo.TagMap](new(repo.TagMapImpl))
	depin.Set[repo.Tag](new(repo.TagImpl))
	depin.Set[repo.Vote](new(repo.VoteImpl))

	depin.Set[service.Article](new(service.ArticleImpl))
	depin.Set[service.Vote](new(service.VoteImpl))

	depin.Set[routes.Handler](new(routes.ArticleHandler))
	depin.Set[routes.Handler](new(routes.VoteHandler))

	routes.InitRouter(router)

	router.Serve()
}
