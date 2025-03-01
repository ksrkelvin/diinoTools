package database

import (
	"context"
	"errors"

	"github.com/ksrkelvin/diinoTools/pkg/database/gorm"
	"github.com/ksrkelvin/diinoTools/pkg/database/mongo"
)

// DB - Main struct for DB
type DB struct {
	Gorm  *gorm.DB
	Mongo *mongo.DB
}

// InitMongoDB - Initialize the database connection with MongoDB
func InitMongoDB(uriMongo string) (conn *mongo.DB, err error) {
	newConn := &mongo.DB{}

	if uriMongo != "" {
		newConn, err = mongo.InitMongoDB(uriMongo)
		if err != nil {
			return conn, err
		}
	}
	if newConn == nil {
		err = errors.New("No database connection was initialized")
	}
	return newConn, err
}

// InitSQL - Initialize the database connection with SQL
func InitSQL(host string, port string, dbName string, user string, pass string, sqlType string) (conn *gorm.DB, err error) {
	newConn := &gorm.DB{}

	if host != "" && port != "" && dbName != "" && user != "" && pass != "" {
		newConn, err = gorm.InitGorm(sqlType, host, port, dbName, user, pass)
		if err != nil {
			return conn, err
		}
	}
	if newConn == nil {
		err = errors.New("No database connection was initialized")
	}
	return newConn, err
}

// Close - Close the database connection
func (p *DB) Close() {
	if p.Gorm != nil {
		sqlDB, err := p.Gorm.DB.DB()
		if err != nil {
			return
		}
		sqlDB.Close()
	}
	if p.Mongo != nil {
		p.Mongo.Disconnect(context.Background())
	}
}
