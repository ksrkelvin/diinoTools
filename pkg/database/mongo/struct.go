package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB - MongoDB
type DB struct {
	*mongo.Client
}

// InitMongoDB - Initialize the MongoDB connection
func InitMongoDB(uriMongo string) (conn *DB, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uriMongo).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	newConn, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	conn = &DB{
		Client: newConn,
	}

	return conn, err
}
