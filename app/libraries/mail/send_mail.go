package mail

import (
	"net/smtp"

	"github.com/spf13/viper"
)

func SendEmail(to []string, cc []string, subject string, message string) error {

	from := viper.GetString("email.from")
	password := viper.GetString("email.pass")
	smtpHost := viper.GetString("email.host")
	smtpPort := viper.GetString("email.port")

	msg := []byte(message)
	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, msg)
	return err
}
