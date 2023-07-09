package email

import (
	"log"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

var (
	EMAIL_SMTP_HOST = os.Getenv("EMAIL_SMTP_HOST")
	EMAIL_SMTP_PORT = os.Getenv("EMAIL_SMTP_PORT")
	EMAIL_USER      = os.Getenv("EMAIL_USER")
	EMAIL_PWD       = os.Getenv("EMAIL_PWD")
)

func Send() {
	m := gomail.NewMessage()
	m.SetHeader("From", "oscarming@126.com")
	m.SetHeader("To", "hy44jt@gmail.com")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>晓明，下午好！</b>!")
	var d *gomail.Dialer
	if port, err := strconv.Atoi(EMAIL_SMTP_PORT); err != nil {
		log.Fatalln("EMAIL_SMTP_PORT is not a number")
	} else {
		d = gomail.NewDialer(EMAIL_SMTP_HOST, port, EMAIL_USER, EMAIL_PWD)
	}
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
