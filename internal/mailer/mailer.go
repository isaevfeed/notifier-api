package mailer

import (
	"fmt"
	"os"
	"strings"

	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
)

type Mailer struct {
	Email    string
	Password string
}

func New() *Mailer {
	Email, _ := os.LookupEnv("SMTP_EMAIL")
	Password, _ := os.LookupEnv("SMTP_PASS")

	return &Mailer{Email, Password}
}

func (m *Mailer) Send(email, subject, message string) error {
	auth := sasl.NewPlainClient("", m.Email, m.Password)

	to := []string{email}
	msg := strings.NewReader(fmt.Sprintf("To: %s\r\n", email) +
		fmt.Sprintf("Subject: %s\r\n", subject) + "\r\n" + message + "\r\n")

	err := smtp.SendMail("smtp.gmail.com:587", auth, m.Email, to, msg)
	if err != nil {
		return err
	}

	return nil
}
