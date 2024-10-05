package apiDB

import (
	"context"
	"log"
	"os"
	"github.com/subosito/gotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Ctx context.Context
	UserCollection *mongo.Collection
	TaskCollection *mongo.Collection
)
// the function might not be exported if the function name start with lowercase
func ConnectMongo() error {
	gotenv.Load()
	Ctx = context.TODO()
	mongoUrl := os.Getenv("MONGO_URL")
	if mongoUrl == "" {
		log.Fatal("Url not found")
	}
	clientMongo  := options.Client().ApplyURI(mongoUrl)

	mongoClient, err := mongo.Connect(Ctx, clientMongo)
	if err != nil {
		log.Fatal("Not able to connect")
		return err
	}
	MongoDB := mongoClient.Database("Api-task")
	UserCollection = MongoDB.Collection("userCollection")
	TaskCollection = MongoDB.Collection("taskCollection")
	return nil
}