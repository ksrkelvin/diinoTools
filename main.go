package diinotools

import (
	"github.com/ksrkelvin/diinoTools/pkg/database"
	"github.com/ksrkelvin/diinoTools/pkg/mail"
	"github.com/ksrkelvin/diinoTools/pkg/tools"
)

type Diino struct {
	Db      *database.DB
	Mailler *mail.Mailler
	Tools   *tools.Tools
}

func New() (diino *Diino, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	var newDiino = &Diino{}

	return newDiino, err

}

func (p *Diino) InitDb(host string, port string, dbName string, user string, pass string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	p.Db, err = database.Init(host, port, dbName, user, pass)
	if err != nil {
		return err
	}
	return
}

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
