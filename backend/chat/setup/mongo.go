package setup

import (
	"chat/constants"
	"chat/types"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BsonMessage struct{}

func InitMongo() *types.MongoClient {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(
		"mongodb://"+constants.MONGO_INITDB_ROOT_USERNAME+":"+constants.MONGO_INITDB_ROOT_PASSWORD+
			"@"+constants.MONGO_IP))
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database(constants.MONGO_INITDB_DATABASE).Collection("chats")
	log.Println("connected to mongodb")
	return &types.MongoClient{Client: client, ChatHistory: collection}
}
