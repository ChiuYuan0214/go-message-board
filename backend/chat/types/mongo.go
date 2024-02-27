package types

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	Client      *mongo.Client
	ChatHistory *mongo.Collection
}

func (mc *MongoClient) Close() {
	if err := mc.Client.Disconnect(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func (mc *MongoClient) BatchInsert(list []interface{}) {
	result, err := mc.ChatHistory.InsertMany(context.Background(), list)
	if err != nil {
		log.Fatal("error when InsertMany:", err)
	}

	log.Println("Inserted IDs:", result.InsertedIDs)
}

func (mc *MongoClient) FindAll(condition primitive.D, options *options.FindOptions) *[]Chat {
	cursor, err := mc.ChatHistory.Find(context.Background(), condition, options)
	if err != nil {
		log.Fatal("error when Find:", err)
	}
	defer cursor.Close(context.Background())

	var chats []Chat
	if err := cursor.All(context.Background(), &chats); err != nil {
		log.Fatal("error when cursor.All:", err)
	}
	return &chats
}
