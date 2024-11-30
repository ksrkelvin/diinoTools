package diinotools

import (
	"errors"

	"github.com/ksrkelvin/diinoTools/pkg/database"
	"github.com/ksrkelvin/diinoTools/pkg/mail"
	"github.com/ksrkelvin/diinoTools/pkg/tools"
)

// Diino - Main struct for DiinoTools
type Diino struct {
	Db      *database.DB
	Mailler *mail.Mailler
	Tools   *tools.Tools
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
  - host: MySQL Host
  - port: MySQL Port
  - dbName: MySQL Database Name
  - user: MySQL User
  - pass: MySQL Password
*/
func (p *Diino) InitDb(uriMongo string, host string, port string, dbName string, user string, pass string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	if uriMongo != "" {
		p.Db.Mongo, err = database.InitMongoDB(uriMongo)
		if err != nil {
			return err
		}
	}
	if host != "" && port != "" && dbName != "" && user != "" && pass != "" {
		p.Db.Mysql, err = database.InitMySQL(host, port, dbName, user, pass)
		if err != nil {
			return err
		}
	}
	if p.Db.Mongo == nil && p.Db.Mysql == nil {
		err = errors.New("No database connection was initialized")
	}

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
