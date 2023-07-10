package email

import (
	"fmt"

	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

func newDialer() *gomail.Dialer {
	return gomail.NewDialer(
		viper.GetString("email.smtp.host"),
		viper.GetInt("email.smtp.port"),
		viper.GetString("email.smtp.user"),
		viper.GetString("email.smtp.password"),
	)
}

func newMessage(to string, subject string, body string) *gomail.Message {
	m := gomail.NewMessage()
	m.SetHeader("From", "oscarming@126.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	return m
}

func Send() {
	m := newMessage("hy44jt@gmail.com", "Hello", "Hello, <h1>Simon</h1>")
	d := newDialer()

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

func SendValidationCode(email, code string) error {
	m := newMessage(
		email,
		fmt.Sprintf("[%s] 验证码", code),
		fmt.Sprintf(`你正在登录或注册账号，你的验证码是：%s`, code),
	)
	d := newDialer()
	return d.DialAndSend(m)
}
