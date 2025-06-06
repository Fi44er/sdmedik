package mailer

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/jordan-wright/email"
)

// Mailer структура для управления отправкой писем
type Mailer struct {
	SMTPClient *email.Pool
	Template   *template.Template
}

// NewMailer инициализирует новый Mailer с пулом соединений и шаблоном
func NewMailer(smtpHost, smtpPort, username, password, templatePath string, poolSize int) (*Mailer, error) {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return nil, fmt.Errorf("failed to parse email template: %w", err)
	}

	tlsConfig := tls.Config{
		InsecureSkipVerify: true,
	}

	auth := smtp.PlainAuth("", username, password, smtpHost)
	pool, err := email.NewPool(fmt.Sprintf("%s:%s", smtpHost, smtpPort), poolSize, auth, &tlsConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create email pool: %w", err)
	}

	return &Mailer{
		SMTPClient: pool,
		Template:   tmpl,
	}, nil
}

// SendMail отправляет письмо с кодом подтверждения
// func (m *Mailer) SendMail(from, to, subject string, templateData interface{}) error {
// 	e := email.NewEmail()
// 	e.From = from
// 	e.To = []string{to}
// 	e.Subject = subject
//
// 	var body bytes.Buffer
// 	if err := m.Template.Execute(&body, templateData); err != nil {
// 		return fmt.Errorf("failed to execute email template: %w", err)
// 	}
//
// 	e.HTML = body.Bytes()
//
// 	if err := m.SMTPClient.Send(e, 10*time.Second); err != nil {
// 		return fmt.Errorf("failed to send email: %w", err)
// 	}
//
// 	log.Printf("Email sent successfully to %s", to)
// 	return nil
// }

func (m *Mailer) SendMail(from, subject string, templateData interface{}, to []string) error {
	e := email.NewEmail()
	e.From = from
	e.To = to
	e.Subject = subject

	var body bytes.Buffer
	if err := m.Template.Execute(&body, templateData); err != nil {
		return fmt.Errorf("failed to execute email template: %w", err)
	}
	e.HTML = body.Bytes()

	backOff := backoff.NewExponentialBackOff()
	backOff.InitialInterval = 1 * time.Second
	backOff.MaxInterval = 30 * time.Second
	backOff.MaxElapsedTime = 5 * time.Minute

	// Функция для уведомления о неудачных попытках
	notify := func(err error, duration time.Duration) {
		log.Printf("SendMail attempt failed: %v. Next try in %v", err, duration)
	}

	operation := func() error {
		return m.SMTPClient.Send(e, 10*time.Second)
	}

	err := backoff.RetryNotify(operation, backOff, notify)
	if err != nil {
		return fmt.Errorf("failed to send email after multiple attempts: %w", err)
	}

	log.Printf("Email sent successfully to %s", to)
	return nil
}

// SendMailAsync отправляет письмо асинхронно
func (m *Mailer) SendMailAsync(from, subject string, templateData interface{}, to []string) {
	go func() {
		if err := m.SendMail(from, subject, templateData, to); err != nil {
			log.Printf("Error sending email to %s: %v", to, err)
		}
	}()
}
