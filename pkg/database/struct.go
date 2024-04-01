package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	Db *gorm.DB
}

func Init(host string, port string, dbName string, user string, pass string) (conn *DB, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	dsn := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return conn, err
	}

	newConn := &DB{
		Db: db,
	}

	return newConn, err
}
