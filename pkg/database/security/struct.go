package security

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// DB - MongoDB
type DB struct {
	db *mongo.Database
}

// Init - Initialize the MongoDB connection
func Init(mongoCli *mongo.Client) (conn *DB, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	conn = &DB{
		db: mongoCli.Database("security"),
	}

	return conn, err
}
