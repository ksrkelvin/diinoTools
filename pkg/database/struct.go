package database

import (
	"context"
	"errors"

	"github.com/ksrkelvin/diinoTools/pkg/database/mongo"
)

// DB - Main struct for DB
type DB struct {
	Mongo *mongo.DB
}

// InitMongoDB - Initialize the database connection with MongoDB
func InitMongoDB(uriMongo string) (conn *DB, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	newMongo := &mongo.DB{}

	if uriMongo != "" {
		newMongo, err = mongo.Init(uriMongo)
		if err != nil {
			return conn, err
		}
	}
	if newMongo == nil {
		err = errors.New("No database connection was initialized")
	}

	newConn := &DB{
		Mongo: newMongo,
	}

	return newConn, err
}

// Close - Close the database connection
func (p *DB) Close() {
	if p.Mongo != nil {
		p.Mongo.Disconnect(context.Background())
	}
}
