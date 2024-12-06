package diinotools

import (
	"github.com/ksrkelvin/diinoTools/pkg/auth"
	"github.com/ksrkelvin/diinoTools/pkg/database"
	"github.com/ksrkelvin/diinoTools/pkg/mail"
	"github.com/ksrkelvin/diinoTools/pkg/security"
	"github.com/ksrkelvin/diinoTools/pkg/tools"
)

// Diino - Main struct for DiinoTools
type Diino struct {
	Db       *database.DB
	Mailler  *mail.Mailler
	Tools    *tools.Tools
	Auth     *auth.Auth
	Security *security.Security
}

// New - Create a new instance of DiinoTools
func New() (diino *Diino, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	var newDiino = &Diino{
		Tools: &tools.Tools{},
	}
	return newDiino, err

}

/*
InitDb - Initialize the database connection

if uriMongo is diferent of "" the connection will be tried with MongoDB
if any of others parameters is equal of "" the connection will not be tried with MySQL
MongoDB Connection:
  - uriMongo: MongoDB URI

Mysql Connection:
  - host:  Host
  - port:  Port
  - dbName:  Database Name
  - user:  User
  - pass:  Password
  - sqlType:  Type (mysql, oracle, postgres, etc)
*/
func (p *Diino) InitDb(uriMongo string, host string, port string, dbName string, user string, pass string, sqlType string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	p.Db, err = database.InitDB(uriMongo, host, port, dbName, user, pass, sqlType)

	return
}

// InitMail - Initialize the mail connection
func (p *Diino) InitMail(host string, port int, user string, pass string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	p.Mailler, err = mail.Init(host, port, user, pass)
	if err != nil {
		return err
	}
	return
}

// InitSecurity - Initialize the security connection
func (p *Diino) InitSecurity() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	p.Security = security.InitSecurity(p.Db)
	return
}

// InitAuth - Initialize the auth connection
func (p *Diino) InitAuth(secret string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	p.Auth, err = auth.Init(secret)
	return
}

// Close - Close the database connection
func (p *Diino) Close() (err error) {
	p.Db.Close()
	return
}
