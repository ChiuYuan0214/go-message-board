package jobs

import (
	"log"
	"security/store"
	"time"
)

func initUsers() {
	for {
		rows, err := connPool.Query(`select u.user_id, u.username, ifnull(i.file_name, '') from users u left join images i on i.user_id = u.user_id`)
		if err != nil {
			log.Println(err)
			time.Sleep(1 * time.Hour)
			continue
		}
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
		rows.Close()
		store.SetUsers(users)
		time.Sleep(1 * time.Hour)
	}
}
