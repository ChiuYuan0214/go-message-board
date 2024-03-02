package jobs

import (
	"chat/store"
	"chat/types"
)

var chatStore = store.GetChatStore()

var dynamo *types.DynamoClient

func UseDynamo(dc *types.DynamoClient) {
	dynamo = dc
}

func UseScheduler() {
	validateHistoryJob()
	incrementHistoryRefJob()
	syncHistoryJob()
	removeLogoutUsersJob()
}
