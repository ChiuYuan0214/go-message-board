package store

import "sync"

type User struct {
	UserId    uint64 `json:"userId"`
	UserName  string `json:"username"`
	UserImage string `json:"userImage"`
}

type UsersStore struct {
	usersLock sync.RWMutex
	users     []User
}

func NewUsersStore() *UsersStore {
	return &UsersStore{
		users: make([]User, 0),
	}
}

func (s *UsersStore) GetUsers() []User {
	s.usersLock.RLock()
	defer s.usersLock.RUnlock()

	return append([]User(nil), s.users...)
}

func (s *UsersStore) AddUser(user User) {
	s.usersLock.Lock()
	defer s.usersLock.Unlock()

	s.users = append(s.users, user)
}

func (s *UsersStore) SetUsers(newUsers []User) {
	s.usersLock.Lock()
	defer s.usersLock.Unlock()

	s.users = append([]User(nil), newUsers...)
}
