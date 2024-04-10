package rabbitmq

import (
	"errors"
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

func SendEmail(emailDTO EmailDTO) error {
	m := gomail.NewMessage()
	m.SetHeader("From", viper.GetString("mail.username"))
	m.SetHeader("To", emailDTO.To)
	m.SetHeader("Subject", emailDTO.Subject)
	m.SetBody("text/plain", emailDTO.Message)
	d := gomail.NewDialer(viper.GetString("mail.host"),
		viper.GetInt("mail.port"),
		viper.GetString("mail.username"),
		viper.GetString("mail.password"))
	if err := d.DialAndSend(m); err != nil {
		return errors.New("发送邮件失败")
	}

	return nil
}
