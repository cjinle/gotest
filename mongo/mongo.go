package mongo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
)

var collection *mongo.Collection

func init() {
	collection = GetCollection()
}

func Close() {}

func GetCollection() *mongo.Collection {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return client.Database("test").Collection("testcol")
}

func Find() {
	// var result bson.M
	var result struct {
		Id   primitive.ObjectID `bson:"_id"`
		Name string             `bson:"name"`
	}
	err := collection.FindOne(context.Background(), bson.D{}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result.Id.Hex())
}
