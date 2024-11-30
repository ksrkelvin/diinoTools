package gorm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB - Main struct for DB
type DB struct {
	*gorm.DB
}

// InitGorm - Initialize the Gorm connection
func InitGorm(sqlType string, host string, port string, dbName string, user string, pass string) (conn *DB, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	uri := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	switch sqlType {
	case "mysql":
		newConn, err := gorm.Open(mysql.Open(uri), &gorm.Config{})
		if err != nil {
			return conn, err
		}
		conn = &DB{
			newConn,
		}

	}

	return conn, err
}
