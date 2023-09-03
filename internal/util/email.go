package util

import (
	"net/smtp"
	// ... potentially other imports for more advanced email functionality ...
)

const (
	SMTPServer   = "smtp.example.com"
	SMTPPort     = "587"
	SMTPUser     = "your_email@example.com"
	SMTPPassword = "your_email_password"
)

// SendEmail sends an email to the specified recipient.
func SendEmail(recipient string, subject string, body string) error {
	from := SMTPUser
	msg := "From: " + from + "\n" +
		"To: " + recipient + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	err := smtp.SendMail(SMTPServer+":"+SMTPPort,
		smtp.PlainAuth("", SMTPUser, SMTPPassword, SMTPServer),
		from, []string{recipient}, []byte(msg))

	return err
}

// ... Potentially add other email utility functions, like generating password reset links, etc. ...
