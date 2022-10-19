package pkg

import (
	"crypto/tls"
	"fmt"
	"tpk-backend/app/pkg/config"

	gomail "gopkg.in/gomail.v2"
)

func Smtp2(sub string, to string, body string) error {
	mail := config.LoadMailerStruct()
	host := mail.Host
	from := mail.Username
	password := mail.Password
	fmt.Println(sub)
	fmt.Println(to)
	fmt.Println(body)
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
