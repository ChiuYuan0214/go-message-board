package main

import (
	"security/constants"
	"security/routes"
	"security/services"
	"security/setup"
)

func main() {
	constants.InitEnv()
	db := setup.InitMySQL()
	routes.UsePool(db)
	services.UsePool(db)

	cache := setup.InitCache()
	services.UseCache(cache)

	routes.UseDispatcher()
	setup.InitServer()

	db.Close()
}
