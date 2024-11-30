package mongo

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
func (p *DB) GetProhibitedPaths(path string) (prohibitedPaths models.ProhibitedPathsStruct, err error) {
	db := p.Database(models.SecurtyDatabase)

	prohibitedPathsCollection := db.Collection(models.ProhibitedPathsCollection)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"path": path}
	err = prohibitedPathsCollection.FindOne(ctx, filter).Decode(&prohibitedPaths)
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
