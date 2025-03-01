package security

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ksrkelvin/diinoTools/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetProhibitedPaths - GetProhibitedPaths
func (p *DB) GetProhibitedPaths(path string) (prohibitedPaths models.PathsStruct, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	pathsCollection := p.db.Collection(models.PathsCollection)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"path": path}
	err = pathsCollection.FindOne(ctx, filter).Decode(&prohibitedPaths)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("path not found!")
			return prohibitedPaths, nil
		}
		log.Fatalf("Error to find path on mongo: %v", err)
		return
	}

	return

}
