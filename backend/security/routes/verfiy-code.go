package routes

import (
	"fmt"
	"net/http"
	"security/services"
	"security/utils"
)

type VerifyData struct {
	UserId int64 `json:"userId"`
	Code   int32 `json:"code"`
}

var verifyMap MethodMapType = map[string]HandlerType{}

func init() {
	verifyMap.post(doVerify)
}

func handleVerifyCode(writer http.ResponseWriter, req *http.Request) {
	setHeader(writer, "json")
	res, status := verifyMap.useHandler(writer, req)
	DoResponse(res, status, writer)
}

func doVerify(req *http.Request) (res interface{}, statusCode int) {
	data := &VerifyData{}
	message, status := utils.ParseBody(req.Body, data)
	if message != "" {
		return newRes("fail").message(message), status
	}
	if data.Code == 0 || data.UserId == 0 {
		return newRes("fail").message("userId and code cannot be empty."), http.StatusBadRequest
	}

	// find respective verification code by userId
	rows, err := connPool.Query(`select code from verification_codes where user_id = ? and valid = true`, data.UserId)
	if isErr(err) {
		return newRes("fail").message("failed to query."), http.StatusInternalServerError
	}
	exist := rows.Next()
	if !exist {
		return newRes("fail").message("code does not exist."), http.StatusOK
	}
	var code string
	err = rows.Scan(&code)
	if err != nil {
		return newRes("fail").message("something went wrong."), http.StatusInternalServerError
	}
	if code != fmt.Sprintf("%06d", data.Code) {
		return newRes("fail").message("code does not match."), http.StatusOK
	}

	services.ActivateUser(data.UserId)

	// generate token
	token := services.GenerateToken(data.UserId)
	if token == nil {
		return newRes("fail").message("failed to generate token."), http.StatusInternalServerError
	}

	return newRes("success").setItem("token", token.Token).setItem("expireTime", token.ExpireTime), http.StatusOK
}
