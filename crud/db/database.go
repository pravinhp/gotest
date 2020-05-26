package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	Url = "mongodb://0.0.0.0:27017"
)

type Connection struct {
	Url string
}

func Connect() *mongo.Collection {
	url := Url

	clientOptions := options.Client().ApplyURI(url)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	fmt.Println("ping")
	if err != nil {
		log.Fatal(err)
	}

	return client.Database("mongocrud").Collection("person")
}

// func Find() string {
// 	collection := Connect(url)
// 	data, _ := collection.Find(context, bson.M{})
// 	return "test"
// }9

func Insert() {
	// c := Connection{Url}
	// collection := c.Connect()
	// res, err := collection.InsertOne(context.TODO(), person)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
