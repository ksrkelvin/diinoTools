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

// InitDB - Initialize the database connection
func InitDB(uriMongo string, host string, port string, dbName string, user string, pass string, sqlType string) (conn *DB, err error) {
	newConn := &DB{}

	if uriMongo != "" {
		newConn.Mongo, err = mongo.InitMongoDB(uriMongo)
		if err != nil {
			return conn, err
		}
	}
	if host != "" && port != "" && dbName != "" && user != "" && pass != "" {
		newConn.Gorm, err = gorm.InitGorm(sqlType, host, port, dbName, user, pass)
		if err != nil {
			return conn, err
		}
	}
	if newConn.Mongo == nil && newConn.Gorm == nil {
		err = errors.New("No database connection was initialized")
	}
	return
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
