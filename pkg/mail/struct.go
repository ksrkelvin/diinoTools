package mail

import mail "github.com/xhit/go-simple-mail/v2"

// Mailler - Main struct for Mailler
type Mailler struct {
	User   string
	Server *mail.SMTPServer
}

// Init - Initialize the mail connection
func Init(host string, port int, user string, pass string) (mailler *Mailler, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	server := mail.NewSMTPClient()
	server.Host = host
	server.Port = port
	server.Username = user
	server.Password = pass
	server.Encryption = mail.EncryptionSSLTLS

	newMailler := &Mailler{
		User:   user,
		Server: server,
	}

	return newMailler, err

}
