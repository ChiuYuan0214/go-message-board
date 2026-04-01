package main

import (
	"chat/constants"
	"chat/infra"
	"chat/jobs"
	"chat/repo"
	"chat/routes"
	"chat/services"

	"github.com/ChiuYuan0214/depin"
)

func main() {
	constants.InitEnv()

	router := depin.RunAndSet[routes.Router](new(routes.RouterImpl))

	depin.Set[infra.RDB](new(infra.MySQL))
	depin.Set[infra.Cache](new(infra.Redis))
	depin.Set[infra.Dynamo](new(infra.DynamoDB))

	depin.Set[repo.Token](new(repo.TokenImpl))
	depin.Set[repo.Follow](new(repo.FollowImpl))
	depin.Set[repo.History](new(repo.HistoryImpl))

	depin.Set[services.Chat](new(services.ChatImpl))
	depin.Set[services.Token](new(services.TokenImpl))
	depin.Set[services.Event](new(services.EventImpl))
	depin.Set[services.History](new(services.HistoryImpl))
	depin.Set[services.Message](new(services.MessageImpl))
	depin.Set[services.Follow](new(services.FollowImpl))
	depin.Set[services.FollowList](new(services.FollowListImpl))
	depin.Set[services.Notify](new(services.NotifyImpl))
	depin.Set[jobs.Scheduler](new(jobs.SchedulerImpl))

	depin.Set[routes.Handler](new(routes.ChatHandler))
	depin.Run()

	router.Serve()
}
