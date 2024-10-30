package services

import (
	"electomock/config"
	"fmt"
	"net/smtp"
)

type EmailService struct {
	config config.SMTPConfig
}

func NewEmailService(config config.SMTPConfig) *EmailService {
	return &EmailService{config: config}
}

func (s EmailService) SendmagicLink(email, token string) error {
	auth := smtp.PlainAuth("", s.config.Username, s.config.Password, s.config.Host)

	msg := fmt.Sprintf(`From: %s
		To: %s
		Subject: Your Login Link
		Content-Type: text/html
		
		<h1>Your Magic Login Link</h1>
		<p>Click the link below to log in:</p>
		<a href="http://localhost:8080/auth/verify?token=%s">Log in</a>
		`, s.config.From, email, token)
	return smtp.SendMail(
		s.config.Host+":"+s.config.Port, auth,
		s.config.From,
		[]string{email},
		[]byte(msg),
	)
}
