package main

import (
	"crypto/tls"
	"github.com/xhit/go-simple-mail/v2"
	stdmail "net/mail"
	"os"
	"strconv"
)

func smtpSendMail(toAddress string, toName string, contentType string, subject string, body string) error {
	server := mail.NewSMTPClient()

	// These are validated in `smtpConfigure`
	server.Host = os.Getenv("SMTP_HOST")
	server.Port, _ = strconv.Atoi(os.Getenv("SMTP_PORT"))
	server.Username = os.Getenv("SMTP_USERNAME")
	server.Password = os.Getenv("SMTP_PASSWORD")

	server.TLSConfig = &tls.Config{InsecureSkipVerify: os.Getenv("SMTP_SKIP_HOST_VERIFY") == "true"}

	smtpClient, err := server.Connect()

	if err != nil {
		return err
	}

	fromAddress := stdmail.Address{"Commento", os.Getenv("SMTP_FROM_ADDRESS")}
	to := stdmail.Address{toName, toAddress}
	email := mail.NewMSG()
	email.SetFrom(fromAddress.String())
	email.AddTo(to.String())
	email.SetSubject(subject)
	email.SetBody(mail.TextPlain, body)

	if email.Error != nil {
		return email.Error
	}

	return email.Send(smtpClient)
}
