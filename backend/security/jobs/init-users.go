package jobs

import (
	"log"
	"security/store"
)

func initUsers() {
	rows, err := connPool.Query(`select u.user_id, u.username, ifnull(i.file_name, '') from users u left join images i on i.user_id = u.user_id`)
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()
	users := []store.User{}
	for rows.Next() {
		var user store.User
		err = rows.Scan(&user.UserId, &user.UserName, &user.UserImage)
		if err != nil {
			log.Println(err)
			continue
		}
		users = append(users, user)
	}
	store.SetUsers(users)
}
