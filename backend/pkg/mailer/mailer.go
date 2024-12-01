package mailer

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
)

func SendMail(from, password, host, port, to string) error {
	var body bytes.Buffer

	tmpl, err := template.ParseFiles("pkg/mailer/template/index.html")
	if err != nil {
		return err
	}

	body.Write([]byte("MIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\r\n"))
	body.Write([]byte(fmt.Sprintf("Subject: Регистрация на Хакатон\r\n\r\n")))

	if err := tmpl.Execute(&body, nil); err != nil {
		return err
	}

	auth := smtp.PlainAuth("", from, password, host)
	if err := smtp.SendMail(host+":"+port, auth, from, []string{to}, body.Bytes()); err != nil {
		return err
	}

	return nil
}
