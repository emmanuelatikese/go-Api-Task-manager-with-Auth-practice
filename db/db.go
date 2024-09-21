package apiDB

import (
	"context"
	"log"
	"os"
	"github.com/subosito/gotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Ctx context.Context
var MongoDB *mongo.Database

// the function might not be exported if the function name start with lowercase
func ConnectMongo() error {
	gotenv.Load()
	Ctx = context.TODO()
	mongoUrl := os.Getenv("MONGO_URL")
	if mongoUrl == "" {
		log.Fatal("Url not found")
	}
	clientMongo  := options.Client().ApplyURI(mongoUrl)

	mongdb, err := mongo.Connect(Ctx, clientMongo)
	if err != nil {
		log.Fatal("Not able to connect")
		return err
	}
	MongoDB = mongdb.Database("Api-task")
	return nil
}