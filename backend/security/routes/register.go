package routes

import (
	"log"
	"net/http"
	"security/services"
	"security/types"
	"security/utils"
	"sync"
	"time"
)

type RegisterHandler struct {
	router          Router
	registerService services.Register
	authService     services.Auth
	registerMap     MethodMapType
	verifyMap       MethodMapType
	resendCodeMap   MethodMapType
}

func NewRegisterHandler(router Router, registerService services.Register, authService services.Auth) *RegisterHandler {
	return &RegisterHandler{
		router:          router,
		registerService: registerService,
		authService:     authService,
	}
}

func (h *RegisterHandler) Run() {
	h.registerMap = make(MethodMapType)
	h.verifyMap = make(MethodMapType)
	h.resendCodeMap = make(MethodMapType)

	h.registerMap.post(h.newRegister)
	h.verifyMap.post(h.doVerify)
	h.resendCodeMap.post(h.resendCode)

	h.router.Handle("/register", h.handleRegister)
	h.router.Handle("/verifyCode", h.handleVerifyCode)
	h.router.Handle("/resendVerificationCode", h.handleResendCode)
}

func (h *RegisterHandler) handleRegister(writer http.ResponseWriter, req *http.Request) {
	setHeader(writer, "json")
	res, status := h.registerMap.useHandler(writer, req)
	DoResponse(res, status, writer)
}

func (h *RegisterHandler) newRegister(req *http.Request) (res interface{}, statusCode int) {
	data := &types.RegisterData{}
	message, status := utils.ParseBody(req.Body, data)
	if message != "" {
		return newRes("fail").message(message), status
	}

	username := data.Username
	email := data.Email
	if username == "" || email == "" || data.Password == "" {
		return newRes("fail").message("username, email and password cannot be empty."), http.StatusBadRequest
	}

	isExist := h.registerService.CheckEmailExist(email)
	if isExist {
		return newRes("fail").message("user already exist"), http.StatusBadRequest
	}

	var wg sync.WaitGroup
	var userId int64
	var codeId int64
	var isSent bool
	var expireTime time.Time
	var userIdChan = make(chan int64)

	wg.Add(1)
	go func() {
		defer wg.Done()
		// encrypt password
		hashedPassword, err := utils.HashPassword(data.Password)
		isErr(err)
		// insert user info into database
		userId = h.registerService.AddNewUser(username, email, hashedPassword, data.Phone, data.Job, data.Address)
		userIdChan <- userId
		close(userIdChan)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		// generate verification code
		veriCode := utils.GenerateCode()
		isSent = utils.SendVerifyCode(email, veriCode.Code)
		userId := <-userIdChan
		if userId == 0 {
			return
		}
		// insert verification code into database
		codeId = h.registerService.InsertVerificationCode(userId, veriCode.Code, veriCode.ExpireTime)
		expireTime = veriCode.ExpireTime
		h.registerService.ScheduleCodeInvalidation(codeId, veriCode)
	}()

	wg.Wait()

	if userId == 0 {
		return newRes("fail").message("cannot create this user"), http.StatusInternalServerError
	}
	if !isSent {
		return newRes("fail").message("failed to send code."), http.StatusInternalServerError
	}
	if codeId == 0 {
		log.Printf("failed to record verification code for userId %d.", userId)
	}

	return newRes("success").setItem("expireTime", expireTime).setItem("userId", userId), http.StatusOK
}
