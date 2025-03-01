package security

import (
	"context"
	"log"
	"time"

	"github.com/ksrkelvin/diinoTools/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// BlockIP - Bloqueia o IP
func (p *DB) BlockIP(blockIP models.BlockedIPsStruct) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	blocklistCollection := p.db.Collection(models.BlocklistCollection)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create a filter to check if the IP already exists
	filter := bson.M{"ip": blockIP.IP} // Assuming blockIP has an IP field to check

	update := bson.M{"$set": blockIP}

	opts := options.Update().SetUpsert(true) // Upsert flag

	// Use UpdateOne with the upsert option set to true
	_, err = blocklistCollection.UpdateOne(ctx, filter, update, opts)

	if err != nil {
		log.Printf("Failed to block IP: %v", err)
	}
	return
}
