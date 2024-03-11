package store

import "sync"

type User struct {
	UserId    uint64 `json:"userId"`
	UserName  string `json:"username"`
	UserImage string `json:"userImage"`
}

var usersLock = sync.Mutex{}
var users = []User{}

func GetUsers() []User {
	return users
}

func AddUser(user User) {
	usersLock.Lock()
	defer usersLock.Unlock()
	users = append(users, user)
}

func SetUsers(newUsers []User) {
	usersLock.Lock()
	defer usersLock.Unlock()
	users = newUsers
}
