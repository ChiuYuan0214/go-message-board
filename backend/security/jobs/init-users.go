package jobs

import (
	"log"
	"security/repo"
	"security/store"
	"time"
)

type UsersSyncJob struct {
	userRepo   repo.User
	usersStore *store.UsersStore
}

func NewUsersSync(userRepo repo.User, usersStore *store.UsersStore) *UsersSyncJob {
	return &UsersSyncJob{
		userRepo:   userRepo,
		usersStore: usersStore,
	}
}

func (j *UsersSyncJob) Run() {
	go j.initUsers()
}

func (j *UsersSyncJob) initUsers() {
	for {
		users, err := j.userRepo.ListUsers()
		if err != nil {
			log.Println(err)
			time.Sleep(1 * time.Hour)
			continue
		}

		j.usersStore.SetUsers(users)
		time.Sleep(1 * time.Hour)
	}
}
