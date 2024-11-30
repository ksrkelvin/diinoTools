package mongo

import (
	"context"
	"log"
	"time"

	"github.com/ksrkelvin/diinoTools/pkg/models"
)

// BlockIP -  Bloqueia o IP
func (p *DB) BlockIP(blockIP models.BlockedIPsStruct) (err error) {
	db := p.Database(models.SecurtyDatabase)

	blocklistCollection := db.Collection(models.BlocklistCollection)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = blocklistCollection.InsertOne(ctx, blockIP)
	if err != nil {
		log.Printf("Failed to block IP: %v", err)
	}
	return
}
