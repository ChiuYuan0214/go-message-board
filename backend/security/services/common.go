package services

import (
	"errors"
	"net/http"
	"security/repo"
	"security/store"
	"strings"
)

var _ User = (*UserImpl)(nil)
var _ Profile = (*ProfileImpl)(nil)

type UserImpl struct {
	userRepo   repo.User
	usersStore *store.UsersStore
}

func NewUser(userRepo repo.User, usersStore *store.UsersStore) *UserImpl {
	return &UserImpl{
		userRepo:   userRepo,
		usersStore: usersStore,
	}
}

func (s *ProfileImpl) UpdateColumnsById(data interface{}, id *uint64) (string, int) {
	count, err := s.profileRepo.UpdateColumnsById(data, *id)
	if errors.Is(err, repo.ErrInvalidColumns) {
		return "column name invalid.", http.StatusBadRequest
	}
	if err != nil {
		return "failed to update columns.", http.StatusInternalServerError
	}
	if count < 1 {
		return "entity not exist or nothing to update.", http.StatusBadRequest
	}
	return "", 0
}

func (s *UserImpl) GetUsers(name string, userId uint64) []store.User {
	if len(s.usersStore.GetUsers()) == 0 {
		users, err := s.userRepo.ListUsers()
		if err == nil {
			s.usersStore.SetUsers(users)
		}
	}

	users := s.usersStore.GetUsers()
	var filteredUsers []store.User

	for _, user := range users {
		if strings.Contains(user.UserName, name) && userId != user.UserId {
			filteredUsers = append(filteredUsers, user)
		}
	}

	return filteredUsers
}
