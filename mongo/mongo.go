package mongo

import (
	"context"
	"log"

	// "time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
)

type Result struct {
	Id   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
	Sex  int                `bson:"sex"`
}

var collection *mongo.Collection

func init() {
	collection = GetCollection()
}

func Close() {}

func GetCollection() *mongo.Collection {
	// client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// err = client.Connect(ctx)
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatal(err)
	}
	return client.Database("test").Collection("testcol")
}

func Find() {
	// var result bson.M
	var result Result
	log.Println("============ FindOne ===========")
	err := collection.FindOne(context.Background(), bson.D{}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result)

	log.Println("============ Find ===========")

	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	// log.Println(cursor)

	log.Println("============ Cursor Decode All ===========")
	var resultArr []Result
	err = cursor.All(context.Background(), &resultArr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resultArr)

	for cursor.Next(context.Background()) {
		log.Println("========== Current =============")
		val, _ := cursor.Current.Elements()
		log.Println(val[1])
		val2 := cursor.Current.Lookup("_id")
		log.Println(val2.ObjectID().Hex())

		val3, err := cursor.Current.LookupErr("sex")
		log.Println(val3, err)

		log.Println("========== Decode To Insterface =============")
		var v interface{}
		cursor.Decode(&v)
		log.Println(v.(primitive.D)[0].Value.(primitive.ObjectID).Hex())

	}

}

func Insert() {
	result, err := collection.InsertOne(
		context.Background(),
		bson.D{
			{"name", "mongoinsert"},
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result)

}
