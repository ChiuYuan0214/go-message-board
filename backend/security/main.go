package main

import (
	"security/constants"
	"security/infra"
	"security/jobs"
	"security/repo"
	"security/routes"
	"security/services"
	"security/store"
)

func main() {
	constants.InitEnv()

	db := new(infra.MySQL)
	if err := db.Run(); err != nil {
		panic(err)
	}
	defer db.Stop()

	cache := new(infra.Redis)
	if err := cache.Run(); err != nil {
		panic(err)
	}
	defer cache.Stop()

	usersStore := store.NewUsersStore()

	authRepo := repo.NewAuth(db, cache)
	registerRepo := repo.NewRegister(db)
	profileRepo := repo.NewProfile(db)
	userRepo := repo.NewUser(db)

	authService := services.NewAuth(authRepo)
	registerService := services.NewRegister(registerRepo, usersStore)
	profileService := services.NewProfile(profileRepo)
	userService := services.NewUser(userRepo, usersStore)

	usersSyncJob := jobs.NewUsersSync(userRepo, usersStore)
	usersSyncJob.Run()

	router := &routes.RouterImpl{}

	loginHandler := routes.NewLoginHandler(router, authService)
	loginHandler.Run()

	registerHandler := routes.NewRegisterHandler(router, registerService, authService)
	registerHandler.Run()

	profileHandler := routes.NewProfileHandler(router, authService, profileService)
	profileHandler.Run()

	usersHandler := routes.NewUsersHandler(router, authService, userService)
	usersHandler.Run()

	router.HandleStatic()
	router.Serve()
}
