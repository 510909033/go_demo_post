package go_mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

func demo() {
	mongo.NewDeleteOneModel()
}

func DemoMyMongodb() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://172.20.10.40:27017"))
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://172.20.10.40:49153"))

	log.Println(client)
	log.Println(err)

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	err = client.Ping(ctx, readpref.Primary())
	log.Printf("ping, err=%+v\n", err)

	collection := client.Database("testing").Collection("numbers")

	//Insert(collection)
	Querey(collection)
}

func Insert(collection *mongo.Collection) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, bson.D{{"name", "pi"}, {"value", 3.14159}})
	id := res.InsertedID
	log.Printf("id=%+v\n", id)
	log.Printf("insert err=%+v\n", err)
	log.Printf("res=%+v\n", res)
}

func Querey(collection *mongo.Collection) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.D
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("result=%+v", result)
		log.Printf("result.Map=%+v", result.Map())
		// do something with result....
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
}

func Filter(collection *mongo.Collection) {
	var result struct {
		Value float64
	}
	filter := bson.D{{"name", "pi"}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		log.Println("record does not exist")
	} else if err != nil {
		log.Fatal(err)
	} else {
		//有结果
	}

	//collection.Find()
}
