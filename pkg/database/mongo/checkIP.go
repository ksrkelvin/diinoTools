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

// CheckIP - GetCheckIP
func (p *DB) CheckIP(ip string) (prohibitedPaths models.BlockedIPsStruct, err error) {
	db := p.Database(models.SecurtyDatabase)

	BlocklistCollection := db.Collection(models.BlocklistCollection)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"ip": ip}
	err = BlocklistCollection.FindOne(ctx, filter).Decode(&prohibitedPaths)
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
