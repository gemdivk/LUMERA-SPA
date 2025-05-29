package mail

import (
	"net/smtp"
	"os"
)

type SmtpSender struct{}

func NewSmtpSender() *SmtpSender {
	return &SmtpSender{}
}

func (s *SmtpSender) Send(to, subject, body string) error {
	from := os.Getenv("SMTP_EMAIL")
	pass := os.Getenv("SMTP_PASSWORD")
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" + body

	return smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))
}
