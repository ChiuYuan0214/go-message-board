package jobs

import (
	"chat/store"
	"chat/types"
)

var chatStore = store.GetChatStore()

var mongo *types.MongoClient

func UseMongo(mc *types.MongoClient) {
	mongo = mc
}

func UseScheduler() {
	validateHistoryJob()
	incrementHistoryRefJob()
	syncHistoryJob()
	removeLogoutUsersJob()
}
