package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB - Main struct for DB
type DB struct {
	Mysql *gorm.DB
	Mongo *mongo.Client
}

// InitMySQL - Initialize the MySQL connection
func InitMySQL(host string, port string, dbName string, user string, pass string) (conn *gorm.DB, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	uri := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	newConn, err := gorm.Open(mysql.Open(uri), &gorm.Config{})
	if err != nil {
		return conn, err
	}

	return newConn, err
}

// InitMongoDB - Initialize the MongoDB connection
func InitMongoDB(uriMongo string) (conn *mongo.Client, err error) {
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

	return newConn, err
}
