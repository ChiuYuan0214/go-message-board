package main

import (
	"general/constants"
	"general/jobs"
	"general/routes"
	"general/services"
	"general/setup"

	"github.com/gin-gonic/gin"
)

func main() {
	constants.InitEnv()

	db := setup.InitGorm()
	routes.UsePool(db)
	services.UsePool(db)
	jobs.UsePool(db)

	cache := setup.InitCache()
	defer cache.Client.Close()
	services.UseCache(cache)
	jobs.UseCache(cache)
	jobs.UseScheduler()

	router := gin.Default()
	routes.InitRouter(router)
	router.Run()
}
