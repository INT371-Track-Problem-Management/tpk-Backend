package pkg

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/smtp"
	"tpk-backend/app/pkg/config"
)

func SSLemail(to string, subj string, body string) {

	// err := godotenv.Load("../.env")
	// if err != nil {
	// 	log.Fatalf("Some error occured. Err: %s", err)
	// }

	// from := os.Getenv("MAILER_USERNAME")
	// password := os.Getenv("MAILER_PASSWORD")

	// smtpHost := os.Getenv("MAILER_HOST")
	// smtpPort := os.Getenv("MAILER_PORT")
	mail := config.LoadGmail()
	from := mail.Username
	password := mail.Password

	smtpHost := mail.Host
	smtpPort := mail.Port
	servername := smtpHost + ":" + smtpPort
	hostdns := "smtp.example.tld:465"
	host, _, _ := net.SplitHostPort(hostdns)

	headers := make(map[string]string)
	headers["From"] = from
	headers["To"] = to
	headers["Subject"] = subj

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body
	auth := smtp.PlainAuth("", from, password, host)

	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	conn, err := tls.Dial("tcp", servername, tlsconfig)
	if err != nil {
		log.Panic(err)
	}
	c, err := smtp.NewClient(conn, host)
	if err != nil {
		log.Panic(err)
	}
	if err = c.Auth(auth); err != nil {
		log.Panic(err)
	}
	if err = c.Mail(from); err != nil {
		log.Panic(err)
	}
	if err = c.Rcpt(to); err != nil {
		log.Panic(err)
	}
	w, err := c.Data()
	if err != nil {
		log.Panic(err)
	}
	_, err = w.Write([]byte(message))
	if err != nil {
		log.Panic(err)
	}
	err = w.Close()
	if err != nil {
		log.Panic(err)
	}
	c.Quit()
	fmt.Println("Email Sent Successfully!")
}
