package pkg

import (
	"crypto/tls"
	"fmt"
	"tpk-backend/app/config"
	"tpk-backend/app/constants"

	"gopkg.in/gomail.v2"
)

func Smtp2(sub string, to string, body string) error {
	mail := config.LoadMailerStruct()
	host := mail.Host
	from := mail.Username
	password := mail.Password
	m := gomail.NewMessage()
	m.SetHeader("From", "rungmod.kmutt.sit@rungmod.com")

	m.SetHeader("To", to)

	m.SetHeader("Subject", sub)

	m.SetBody("text/plain", body)

	d := gomail.NewDialer(host, 587, from, password)

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func UpdateStatus(to string, reportId int, title string, status string) error {
	detail := fmt.Sprintf("แจ้งเตือนการรายงานปัญหา รหัส: %v หัวข้อ: %v รายละเอียด: ", reportId, title)
	switch status {
	case "waiting":
		detail += constants.BODY_EMAIL_SENDING_REPORT

	case "accept":
		detail += constants.BODY_EMAIL_APPROVE_REPORT

	case "engage":
		detail += constants.BODY_EMAIL_ENAGAGE_REPORT

	case "prepare":
		detail += constants.BODY_EMAIL_PREPARE_REPORT

	case "postpone":
		detail += constants.BODY_EMAIL_POSTPONE_REPORT

	case "cancel":
		detail += constants.BODY_EMAIL_CANCEL_REPORT

	case "success":
		detail += constants.BODY_EMAIL_SUCCESS

	case "defer":
		detail += constants.BODY_EMAIL_DEFER

	case "pending":
		detail += constants.BODY_EMAIL_PENDING
	}
	err := Smtp2(constants.SUBJECT_EMAIL_STATUS_REPORT, to, detail)
	if err != nil {
		return err
	}
	return nil
}
