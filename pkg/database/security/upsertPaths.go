package security

import (
	"context"
	"time"

	"github.com/ksrkelvin/diinoTools/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// UpsertPath - Função para realizar o upsert de um path
func (p *DB) UpsertPath(path string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	pathsCollection := p.db.Collection(models.PathsCollection)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Filtro para encontrar o documento com o path especificado
	filter := bson.M{"path": path}

	// Atualização ou inserção (definição do campo atualizado)
	update := bson.M{
		"$inc": bson.M{"qty": 1}, // Incrementa o campo qty
		"$set": bson.M{
			"timestamp": time.Now(), // Atualiza o timestamp
		},
		"$setOnInsert": bson.M{
			"isProhibited": false, // Define apenas na inserção
		},
	}
	// Opções de upsert
	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)

	// Executa o upsert
	// Resultado será decodificado na struct
	var updatedDoc models.PathsStruct
	err = pathsCollection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&updatedDoc)
	if err != nil {
		return err
	}

	return err
}
