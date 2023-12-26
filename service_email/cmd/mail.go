package cmd

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

const (
	smtpAuthAddress   = "smtp.gmail.com"
	smtpServerAddress = "smtp.gmail.com:587"
)

type EmailSender interface {
	SenderEmail(
		subject string,
		content string,
		to []string,
		cc []string,
		bcc []string,
		attachFiles []string,
	) error
}

type GmailSender struct {
	name          string
	fromEmailAddr string
	fromEmailPass string
	EmailAttach   string
}

func NewGmailSender(name, fromEmailAddr, fromEmailPass string) *GmailSender {
	return &GmailSender{
		name:          name,
		fromEmailAddr: fromEmailAddr,
		fromEmailPass: fromEmailPass,
	}
}

func (g *GmailSender) SenderEmail(
	subject string,
	content string,
	to []string,
	cc []string,
	bcc []string,
	attachFiles []string,
) error {

	e := email.NewEmail()

	e.From = fmt.Sprintf("%s <%s>", g.name, g.fromEmailAddr)

	e.Subject = subject

	e.HTML = []byte(content)

	e.To = to

	e.Cc = cc

	e.Bcc = bcc

	for _, attachFile := range attachFiles {
		_, err := e.AttachFile(attachFile)

		if err != nil {
			return fmt.Errorf("attach file error: %w", err)
		}

	}

	smtpAuth := smtp.PlainAuth("", g.fromEmailAddr, g.fromEmailPass, smtpAuthAddress)

	return e.Send(smtpServerAddress, smtpAuth)

}
