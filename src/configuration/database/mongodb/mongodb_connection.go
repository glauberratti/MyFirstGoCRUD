package mongodb

import (
	"context"
	"os"

	"github.com/glauberratti/MyFirstGoCRUD/src/configuration/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const MONGODB_URL = "MONGODB_URL"
const MONGODB_DB = "MONGODB_DB"

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {
	mongodb_uri := os.Getenv(MONGODB_URL)
	dbName := os.Getenv(MONGODB_DB)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodb_uri))

	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		logger.Error("Erro ao tentar se conectar com o banco de dados", err)
		panic("")
	}

	client.Database(dbName)
	logger.Info("Conectou com sucesso")

	return client.Database(dbName), nil
}
