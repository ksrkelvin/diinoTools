package mail

import mail "github.com/xhit/go-simple-mail/v2"

func (p *Mailler) Send(to []string, cc []string, subject string, body string) (err error) {
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
	server.Encryption = mail.EncryptionTLS

	smtpClient, err := server.Connect()
	if err != nil {
		return err
	}

	email := mail.NewMSG()
	email.SetFrom(p.Username)
	for _, v := range to {
		email.AddTo(v)
	}
	for _, v := range cc {
		email.AddTo(v)
	}
	email.SetSubject(subject)

	email.SetBody(mail.TextHTML, body)

	err = email.Send(smtpClient)
	if err != nil {
		return err
	}

	return err
}
