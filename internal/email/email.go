package email

import (
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

func Send() {
	m := gomail.NewMessage()
	m.SetHeader("From", "oscarming@126.com")
	m.SetHeader("To", "hy44jt@gmail.com")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>晓明，下午好！</b>!")

	d := gomail.NewDialer(
		viper.GetString("email.smtp.host"),
		viper.GetInt("email.smtp.port"),
		viper.GetString("email.smtp.user"),
		viper.GetString("email.smtp.password"),
	)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
