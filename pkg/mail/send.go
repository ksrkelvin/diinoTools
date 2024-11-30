package mail

import (
	"strings"

	mail "github.com/xhit/go-simple-mail/v2"
)

// SendOne - Send an email to one recipient
func (p *Mailler) SendOne(mail string, subject string, body string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	p.send([]string{mail}, []string{}, []string{}, subject, body)
	return
}

// SendMany - Send an email to multiple recipients
func (p *Mailler) SendMany(mail []string, sendType string, subject string, body string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	if strings.EqualFold(sendType, "CC") {
		p.send([]string{}, mail, []string{}, subject, body)
	} else if strings.EqualFold(sendType, "CCO") {
		p.send([]string{}, []string{}, mail, subject, body)
	} else {
		p.send(mail, []string{}, []string{}, subject, body)
	}

	return
}

// Send - Send an email
func (p *Mailler) send(to []string, cc []string, cco []string, subject string, body string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	server := mail.NewSMTPClient()
	server.Host = p.Host
	server.Port = p.Port
	server.Username = p.Username
	server.Password = p.Password
	server.Encryption = mail.EncryptionSSLTLS

	smtpClient, err := server.Connect()
	if err != nil {
		return err
	}

	email := mail.NewMSG()
	email.SetFrom(p.Username)
	email.AddTo(to...)
	email.AddCc(cc...)
	if cco != nil {
		email.AddBccToHeader = true
		email.AddBcc(cco...)
	}

	email.SetSubject(subject)
	email.SetBody(mail.TextHTML, body)

	err = email.Send(smtpClient)
	if err != nil {
		return err
	}

	return err
}
