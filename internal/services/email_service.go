package services

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"strings"
)

type EmailService struct {
	host     string
	port     string
	username string
	password string
}

func NewEmailService(host, port, username, password string) *EmailService {
	return &EmailService{
		host:     host,
		port:     port,
		username: username,
		password: password,
	}
}

func (s *EmailService) SendContactNotification(contactName, contactEmail, subject, message string) error {
	if s.username == "" || s.password == "" {
		// Skip email sending if credentials are not configured
		return nil
	}

	to := []string{s.username}
	msg := fmt.Sprintf(`From: %s
To: %s
Subject: New Contact Form Submission: %s

Name: %s
Email: %s
Subject: %s

Message:
%s

---
This message was sent from your portfolio contact form.`,
		s.username,
		strings.Join(to, ","),
		subject,
		contactName,
		contactEmail,
		subject,
		message)

	auth := smtp.PlainAuth("", s.username, s.password, s.host)

	addr := fmt.Sprintf("%s:%s", s.host, s.port)

	// Create TLS config
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         s.host,
	}

	// Connect to SMTP server
	conn, err := tls.Dial("tcp", addr, tlsConfig)
	if err != nil {
		return fmt.Errorf("failed to connect to SMTP server: %v", err)
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, s.host)
	if err != nil {
		return fmt.Errorf("failed to create SMTP client: %v", err)
	}
	defer client.Close()

	// Authenticate
	if err = client.Auth(auth); err != nil {
		return fmt.Errorf("failed to authenticate: %v", err)
	}

	// Send email
	if err = client.Mail(s.username); err != nil {
		return fmt.Errorf("failed to set sender: %v", err)
	}

	if err = client.Rcpt(s.username); err != nil {
		return fmt.Errorf("failed to set recipient: %v", err)
	}

	w, err := client.Data()
	if err != nil {
		return fmt.Errorf("failed to create message writer: %v", err)
	}

	_, err = w.Write([]byte(msg))
	if err != nil {
		return fmt.Errorf("failed to write message: %v", err)
	}

	err = w.Close()
	if err != nil {
		return fmt.Errorf("failed to close message writer: %v", err)
	}

	return nil
}
