package jobs

import "time"

func removeLogoutUsersJob() {
	go func() {
		for {
			time.Sleep(30 * time.Minute)
			for userId, client := range *chatStore.Clients {
				if !client.IsOnline && client.LogoutTime.Before(time.Now().Add(-10*time.Minute)) {
					chatStore.DeleteClient(userId)
				}
			}
		}
	}()
}
