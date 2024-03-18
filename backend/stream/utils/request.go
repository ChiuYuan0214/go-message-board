package utils

import (
	"net/http"
	"strconv"
)

func GetTokenFromQuery(r *http.Request) string {
	return r.URL.Query().Get("token")
}

func GetLiveIdFromQuery(r *http.Request) uint64 {
	id, err := strconv.ParseUint(r.URL.Query().Get("liveId"), 10, 64)
	if err != nil {
		return 0
	}
	return id
}
