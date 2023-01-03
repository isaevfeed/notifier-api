package mailer

import (
	"fmt"
	"os"
	"strings"

	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
)

type Mailer struct {
	Addr     string
	Email    string
	Password string
}

func New() *Mailer {
	Host, _ := os.LookupEnv("SMTP_HOST")
	Port, _ := os.LookupEnv("SMTP_PORT")
	Email, _ := os.LookupEnv("SMTP_EMAIL")
	Password, _ := os.LookupEnv("SMTP_PASS")

	return &Mailer{Addr: fmt.Sprintf("%s:%s", Host, Port), Email: Email, Password: Password}
}

func (m *Mailer) Send(email, subject, message string) error {
	auth := sasl.NewPlainClient("", m.Email, m.Password)

	to := []string{email}
	msg := strings.NewReader(fmt.Sprintf("To: %s\r\n", email) +
		fmt.Sprintf("Subject: %s\r\n", subject) + "\r\n" + message + "\r\n")

	err := smtp.SendMail(m.Addr, auth, m.Email, to, msg)
	if err != nil {
		return err
	}

	return nil
}
