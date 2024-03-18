package services

import (
	"log"
	"stream/constants"
	"stream/types"
)

type OwnerService struct {
	owner *types.Owner
	LiveService
}

func NewOwnerService(owner *types.Owner) *OwnerService {
	ownerService := &OwnerService{owner: owner}
	ownerService.SetLive(owner.GetLive())
	return ownerService
}

func (ownerService *OwnerService) Handle() {
	owner := ownerService.owner
	for {
		messageType, message, err := owner.GetLiveConn().ReadMessage()
		if err != nil {
			log.Println(err)
			owner.GetLive().OwnerExit()
			return
		}

		switch messageType {
		case constants.LIVE:
			ownerService.PushStream(owner.GetUserId(), owner.GetLive().GetWatchers(), message)
		}
	}
}
