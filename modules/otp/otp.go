package otp

import (
	"bytes"
	"errors"
	"fmt"
	"net/smtp"
	"text/template"
	"time"

	"lms_backend/helpers"

	"github.com/craftzbay/go_grc/v2/data"
	"github.com/craftzbay/go_grc/v2/gvalidate"

	"lms_backend/database"
	"lms_backend/global"
	"lms_backend/utils"
)

type OtpCode struct {
	Id        uint
	Identity  string
	Code      uint
	CreatedAt data.LocalTime `json:"created_at" gorm:"autoCreateTime"`
}

func RunMigrations() {
	db := database.DBconn
	db.AutoMigrate(OtpCode{})
}

func (*OtpCode) TableName() string {
	return utils.GetTableName("otp_codes")
}

func (p *OtpCode) CheckOtp(identity string, maxAllowedTime time.Time) int64 {
	db := database.DBconn
	return db.Where("identity = ? AND created_at > ?", identity, maxAllowedTime).Find(&p).RowsAffected
}

func (p *OtpCode) Save() error {
	db := database.DBconn
	if err := db.Create(p).Error; err != nil {
		return err
	}
	return nil
}

func CheckOtp(identity string, code uint) error {
	db := database.DBconn
	otps := make([]OtpCode, 0)
	timeLimit := time.Now().Add(time.Minute * -5)
	if err := db.Where("identity = ?", identity).Where("created_at > ?", timeLimit).Find(&otps).Error; err != nil {
		return err
	}
	for _, otp := range otps {
		if otp.Code == code {
			return nil
		}
	}
	return errors.New("Wrong OTP Code")
}

func SendOtp(identity string) (err error) {
	if !gvalidate.IsEmail(identity) {
		return fmt.Errorf("Invalid email address")
	}

	otp := OtpCode{}
	maxAllowedTime := time.Now().Add(-time.Second * time.Duration(global.Conf.OtpWaitSecond))
	cnt := otp.CheckOtp(identity, maxAllowedTime)

	if cnt > 0 {
		return fmt.Errorf("OTP Code already sent")
	}

	otp.Code = uint(helpers.GenerateRandom(6))
	otp.Identity = identity

	if err = otp.Save(); err != nil {
		return err
	}

	senderUsername := global.Conf.GmailUsername
	senderPassword := global.Conf.GmailPassword
	gmailAuth := smtp.PlainAuth("", senderUsername, senderPassword, "smtp.gmail.com")

	t, _ := template.ParseFiles(global.Conf.PathOtpTemplate)
	var body bytes.Buffer
	headers := "MIME-version: 1.0;\nContent-Type: text/html;"
	body.Write([]byte(fmt.Sprintf("Subject: OTP Code\n%s\n\n", headers)))

	t.Execute(&body, struct {
		Code uint
	}{
		Code: otp.Code,
	})
	return smtp.SendMail("smtp.gmail.com:587", gmailAuth, senderUsername, []string{otp.Identity}, body.Bytes())
}
