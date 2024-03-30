package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Conn struct {
	Db *gorm.DB
}

func Init() (conn *Conn, err error) {
	var newConn = &Conn{}

	newConn.Db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return conn, err
	}

	return newConn, err
}
