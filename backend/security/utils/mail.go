package utils

import (
	"fmt"
	"log"
	"math/rand"
	"security/constants"
	"time"

	"gopkg.in/mail.v2"
)

type VerificationCode struct {
	Code       int32
	ExpireTime time.Time
}

func GenerateCode() *VerificationCode {
	source := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(source)
	code := generator.Int31n(899999) + 100000
	expireTime := time.Now().Add(time.Duration(constants.VERI_CODE_TTL) * time.Minute)
	return &VerificationCode{code, expireTime}
}

func SendVerifyCode(userMail string, code int32) bool {
	newMail := mail.NewMessage()
	newMail.SetHeader("From", constants.ADMIN_MAIL_ADDRESS)
	newMail.SetHeader("To", userMail)
	newMail.SetHeader("Subject", "Adam's Go-Project - Your verification code")
	newMail.SetBody("text/html; charset=UTF-8", fmt.Sprintf(
		`<h1 style="font-weight: 600; font-size: 26px;">This is your code:</h1>
		<p style="color: #039e63; font-weight: 800; font-size: 30px;">%d</p>`, code))

	dialer := mail.NewDialer("smtp.gmail.com", 587, constants.ADMIN_MAIL_ADDRESS, constants.ADMIN_MAIL_PASSWORD)
	if err := dialer.DialAndSend(newMail); err != nil {
		log.Println(err)
		return false
	}
	return true
}
